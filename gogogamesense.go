package gogogamesense

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// GoGoGameSense implements the JSON API for the SteelSeries Engine 3
// GameSense service
type GoGoGameSense struct {
	// GameName holds the name of the game for API calls
	GameName string
	// APIEndpoint specifies the SteelSeries JSON API Endpoint
	APIEndpoint string
	// client is the HTTP client used to communicate with the API
	client *http.Client
}

const (
	// requestTimeoutSeconds is the waiting time for an API response
	requestTimeoutSeconds int = 2
)

// New creates a new instance of the GoGoGameSense client.
//
// It will attempt to find the SteelSeries JSON API endpoint and return an
// error if it was not possible
func New() (*GoGoGameSense, error) {
	corePropsPath := filepath.Join(os.Getenv("PROGRAMDATA"), "SteelSeries", "SteelSeries Engine 3", "coreProps.json")
	corePropsContent, err := ioutil.ReadFile(corePropsPath)
	if err != nil {
		return nil, fmt.Errorf("SteelSeries Engine 3 not running or not installed: %s", err)
	}
	var coreProps CoreProps
	err = json.Unmarshal(corePropsContent, &coreProps)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse SteelSeries Engine 3 Core Props: %s", err)
	}

	return &GoGoGameSense{
		APIEndpoint: coreProps.Address,
		client:      createHTTPClient(),
	}, nil
}

// NewWithEndpoint creates a new instance of the GoGoGameSense client.
//
// It requires the SteelSeries JSON API endpoint. You can find information on
// where to find this endpoint at
// https://github.com/SteelSeries/gamesense-sdk/blob/master/doc/api/sending-game-events.md
func NewWithEndpoint(apiEndpoint string) (*GoGoGameSense, error) {
	// Remove trailing slash
	apiEndpoint = strings.TrimSuffix(apiEndpoint, "/")
	return &GoGoGameSense{
		APIEndpoint: apiEndpoint,
		client:      createHTTPClient(),
	}, nil
}

// createHTTPClient creates a new HTTP client
func createHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: time.Duration(requestTimeoutSeconds) * time.Second,
	}
	return client
}

// RegisterGame registers your game with the SteelSeries engine
func (gs *GoGoGameSense) RegisterGame(metadata GameMetadata) error {
	if metadata.DeInitializeTimerMS == 0 {
		metadata.DeInitializeTimerMS = 15000
	}
	if metadata.DeInitializeTimerMS < 1000 {
		return fmt.Errorf("DeInitializeTimerMS must be between 1000 and 60000")
	}
	if metadata.DeInitializeTimerMS > 60000 {
		return fmt.Errorf("DeInitializeTimerMS must be between 1000 and 60000")
	}

	gs.GameName = strings.TrimSpace(metadata.Name)
	if gs.IsAllowedGameOrEventName(gs.GameName) == false {
		return fmt.Errorf(
			"The game name may only contain uppercase letters, 0-9, hyphens or underscores",
		)
	}

	gs.GameName = metadata.Name
	err := gs.call("game_metadata", metadata)

	return err
}

// DeregisterGame removes your game from the SteelSeries engine
func (gs *GoGoGameSense) DeregisterGame() error {

	metadata := GameMetadata{
		Name: gs.GameName,
	}
	err := gs.call("remove_game", metadata)

	return err
}

// BindGameEvent creates a new game event the engine will react to
func (gs *GoGoGameSense) BindGameEvent(binding GameEventBinding) error {
	binding.GameName = gs.GameName
	if gs.IsAllowedGameOrEventName(binding.Name) == false {
		return fmt.Errorf(
			"The event name may only contain uppercase letters, 0-9, hyphens or underscores",
		)
	}

	err := gs.call("bind_game_event", binding)

	return err
}

// UnbindGameEvent removes an existing game event
func (gs *GoGoGameSense) UnbindGameEvent(binding GameEventBinding) error {
	binding.GameName = gs.GameName
	if gs.IsAllowedGameOrEventName(binding.Name) == false {
		return fmt.Errorf(
			"The event name may only contain uppercase letters, 0-9, hyphens or underscores",
		)
	}

	err := gs.call("remove_game_event", binding)

	return err
}

// SendGameEvent sends the event to the SteelSeries engine
func (gs *GoGoGameSense) SendGameEvent(event GameEvent) error {

	event.GameName = gs.GameName
	if gs.IsAllowedGameOrEventName(event.Name) == false {
		return fmt.Errorf(
			"The event name may only contain uppercase letters, 0-9, hyphens or underscores",
		)
	}

	err := gs.call("game_event", event)

	return err
}

// SendHeartbeat sends the keep-alive signal to the engine
func (gs *GoGoGameSense) SendHeartbeat() error {
	heartbeat := GameHeartbeat{
		GameName: gs.GameName,
	}
	err := gs.call("game_heartbeat", heartbeat)
	return err
}

// IsAllowedGameOrEventName checks if the name complies with the SteelSeries
// rules of game and event names are limited to the following characters:
// Uppercase A-Z, the digits 0-9, hyphen, and underscore.
func (gs *GoGoGameSense) IsAllowedGameOrEventName(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}
	var isAllowed = regexp.MustCompile(`^[A-Z0-9_-]*$`).MatchString
	return isAllowed(name)
}

// call the API endpoint with the given payload object
func (gs *GoGoGameSense) call(endpoint string, payloadObject interface{}) error {
	payload, err := json.Marshal(payloadObject)
	if err != nil {
		return err
	}

	callEndpoint := fmt.Sprintf("http://%s/%s", gs.APIEndpoint, endpoint)
	req, err := http.NewRequest("POST", callEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("Unable to create API request '%s': %s", endpoint, err)
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := gs.client.Do(req)
	if err != nil && response == nil {
		return fmt.Errorf("Unable to call API '%s': %s", endpoint, err)
	}
	responseContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Unable to read API response: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return nil
	}

	var apiError APIError
	err = json.Unmarshal(responseContent, &apiError)
	if err != nil {
		return fmt.Errorf("Unable to read API error information: %s. Content %s.", err, string(responseContent))
	}
	return fmt.Errorf("Error %s: %s", response.Status, apiError.Error)
}
