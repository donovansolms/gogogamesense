package gogogamesense

// GameHandler defines the JSON structure for event handlers to handle color
// changes
type GameHandler struct {
	DeviceType string      `json:"device-type"`
	Zone       string      `json:"zone"`
	Rate       Frequency   `json:"rate,omitempty"`
	Color      interface{} `json:"color"`
	Mode       string      `json:"mode"`
}

// Frequency defines blinking effects
type Frequency struct {
	Values []RangeFrequency `json:"frequency,omitempty"`
}

// RangeFrequency defines blinking effects for a range
type RangeFrequency struct {
	Low            uint `json:"low"`
	High           uint `json:"high"`
	BlinkPerSecond uint `json:"frequency"`
}

// SingleColor defines a color in RGB
type SingleColor struct {
	Red   uint8 `json:"red"`
	Green uint8 `json:"green"`
	Blue  uint8 `json:"blue"`
}

// RangedColor defines a value range for which a color will be applied
type RangedColor struct {
	Low   uint        `json:"low"`
	High  uint        `json:"high"`
	Color SingleColor `json:"color"`
}
