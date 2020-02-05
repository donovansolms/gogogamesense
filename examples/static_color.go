package main

import (
	"fmt"
	"os"

	"github.com/donovansolms/gogogamesense"
)

func main() {
	fmt.Printf("\n=================\n= GoGoGameSense =\n=================\n\n")
	fmt.Printf("This example will set the colour of a Rival mouse's logo to red\n\n")

	gs, err := gogogamesense.NewWithEndpoint("http://192.168.88.161:49738/")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	metadata := gogogamesense.GameMetadata{
		Name:        "MINE001",
		DisplayName: "Minesweeper",
		Developer:   "CodeBoy",
	}
	err = gs.RegisterGame(metadata)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

}
