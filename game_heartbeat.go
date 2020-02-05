package gogogamesense

// GameHeartbeat defines a keep-alive instruction
type GameHeartbeat struct {
	GameName string `json:"game"`
}
