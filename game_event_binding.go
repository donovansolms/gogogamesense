package gogogamesense

// GameEventBinding defines the JSON structure for creating event bindings
type GameEventBinding struct {
	GameName string        `json:"game"`
	Name     string        `json:"event"`
	MinValue uint          `json:"min_value,omitempty"`
	MaxValue uint          `json:"max_value,omitempty"`
	IconID   uint8         `json:"icon_id,omitempty"`
	Handlers []GameHandler `json:"handlers"`
}
