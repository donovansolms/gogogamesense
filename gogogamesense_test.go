package gogogamesense_test

import (
	"fmt"
	"testing"
)

func TestIsAllowedGameOrEventName(t *testing.T) {
	fmt.Println(gs.IsAllowedGameOrEventName("TEST"))
	fmt.Println(gs.IsAllowedGameOrEventName("TES-T"))
	fmt.Println(gs.IsAllowedGameOrEventName("TEST$"))
	fmt.Println(gs.IsAllowedGameOrEventName("TEST"))
	fmt.Println(gs.IsAllowedGameOrEventName("TEST"))
	fmt.Println(gs.IsAllowedGameOrEventName("TEST"))
}
