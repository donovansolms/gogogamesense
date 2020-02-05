package gogogamesense

// GameMetadata defines the JSON structure for registering a game API calls
type GameMetadata struct {
	Name                string `json:"game"`
	DisplayName         string `json:"game_display_name"`
	Developer           string `json:"developer"`
	DeInitializeTimerMS uint   `json:"deinitialize_timer_length_ms"`
}
