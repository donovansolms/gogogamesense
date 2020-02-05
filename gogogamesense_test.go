package gogogamesense_test

import (
	"os"
	"testing"

	"github.com/donovansolms/gogogamesense"
)

var gs *gogogamesense.GoGoGameSense

func TestMain(m *testing.M) {
	endpoint := "http://127.0.0.1:80"
	var err error
	gs, err = gogogamesense.NewWithEndpoint(endpoint)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	_, err := gogogamesense.New()
	if err != nil {
		t.Error(err)
	}
}

func TestNewWithEndpoint(t *testing.T) {
	endpoint := "http://127.0.0.1:80"
	gs, err := gogogamesense.NewWithEndpoint(endpoint)
	if err != nil {
		t.Error(err)
	}
	if gs.APIEndpoint != endpoint {
		t.Errorf(
			"Incorrect endpoint set, expected %s, got %s", endpoint, gs.APIEndpoint,
		)
	}
}

func TestIsAllowedGameOrEventName(t *testing.T) {
	mustPassTestcases := []string{
		"TEST",
		"TEST_",
		"_TEST",
		"TE_ST",
		"TE_S_T",
		"TEST-",
		"TE-ST",
		"-TEST",
		"001",
		"TEST001",
		"TEST_001",
		"TEST_00-1",
	}
	for _, testcase := range mustPassTestcases {
		isValid := gs.IsAllowedGameOrEventName(testcase)
		if !isValid {
			t.Errorf("Testcase must pass for %s", testcase)
		}
	}

	mustFailTestcases := []string{
		"",
		"TE ST",
		"TE=ST",
		"TE$ST",
		"test",
		"te-ST_CASE",
	}
	for _, testcase := range mustFailTestcases {
		isValid := gs.IsAllowedGameOrEventName(testcase)
		if isValid {
			t.Errorf("Testcase must fail for %s", testcase)
		}
	}

}
