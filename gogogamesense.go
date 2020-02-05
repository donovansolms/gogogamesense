package gogogamesense

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// GoGoGameSense implements the JSON API for the SteelSeries GameSense Engine
type GoGoGameSense struct {
	// gameName holds the name of the game for API calls
	gameName string
	// apiURL specifies the SteelSeries JSON API Endpoint
	apiEndpoint string
}

// New creates a new instance of the GoGoGameSense client.
//
// It will attempt to find the SteelSeries JSON API endpoint and return an
// error if it was not possible
func New() (*GoGoGameSense, error) {

	return &GoGoGameSense{}, nil
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
		apiEndpoint: apiEndpoint,
	}, nil
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

	gs.gameName = strings.TrimSpace(gs.gameName)
	if gs.IsAllowedGameOrEventName(gs.gameName) == false {
		return fmt.Errorf(
			"The game name may only contain uppercase letters, 0-9, hyphens or underscores",
		)
	}

	gs.gameName = metadata.Name
	err := gs.call("game_metadata", metadata)

	return err
}

// IsAllowedGameOrEventName checks if the name complies with the SteelSeries
// rules of game and event names are limited to the following characters:
// Uppercase A-Z, the digits 0-9, hyphen, and underscore.
func (gs *GoGoGameSense) IsAllowedGameOrEventName(name string) bool {
	var isAllowed = regexp.MustCompile(`^[A-Z0-9_-]*$`).MatchString
	return isAllowed(name)
}

// call the API endpoint with the given payload object
func (gs *GoGoGameSense) call(endpoint string, payloadObject interface{}) error {
	payload, err := json.Marshal(payloadObject)
	if err != nil {
		return err
	}
	fmt.Println(string(payload))

	return nil
}
