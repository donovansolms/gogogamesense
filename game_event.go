package gogogamesense

// GameEvent defines the JSON structure for triggering events
type GameEvent struct {
	GameName string        `json:"game"`
	Name     string        `json:"event"`
	Data     GameEventData `json:"data"`
}

// GameEventData holds the event information to send
type GameEventData struct {
	Value uint              `json:"value"`
	Frame map[string]string `json:"frame"`
}
