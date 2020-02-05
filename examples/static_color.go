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
	_ = gs
	_ = metadata
	// err = gs.RegisterGame(metadata)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }

	gs.GameName = "TEST001"
	// event := gogogamesense.GameEvent{
	// 	Name: "TestEvent",
	// 	Handlers: []gogogamesense.GameHandler{
	// 		gogogamesense.GameHandler{
	// 			DeviceType: gogogamesense.DeviceTypeRGB2Zone,
	// 			Zone:       gogogamesense.DeviceZoneMouseLogo,
	// 			Mode:       gogogamesense.ModeColor,
	// 			Color: gogogamesense.SingleColor{
	// 				Red:   255,
	// 				Green: 0,
	// 				Blue:  128,
	// 			},
	// 		},
	// 	},
	// }

	event := gogogamesense.GameEventBinding{
		Name:   "HEALTH",
		IconID: gogogamesense.IconHealth,
		Handlers: []gogogamesense.GameHandler{
			gogogamesense.GameHandler{
				DeviceType: gogogamesense.DeviceTypeRGB2Zone,
				Zone:       gogogamesense.DeviceZoneMouseLogo,
				Mode:       gogogamesense.ModeColor,
				Rate: gogogamesense.Frequency{
					Values: []gogogamesense.RangeFrequency{
						gogogamesense.RangeFrequency{
							Low:            0,
							High:           10,
							BlinkPerSecond: 5,
						},
					},
				},
				Color: []gogogamesense.RangedColor{
					gogogamesense.RangedColor{
						Low:  0,
						High: 10,
						Color: gogogamesense.SingleColor{
							Red:   255,
							Green: 0,
							Blue:  0,
						},
					},
					gogogamesense.RangedColor{
						Low:  11,
						High: 50,
						Color: gogogamesense.SingleColor{
							Red:   255,
							Green: 180,
							Blue:  0,
						},
					},
					gogogamesense.RangedColor{
						Low:  51,
						High: 100,
						Color: gogogamesense.SingleColor{
							Red:   0,
							Green: 255,
							Blue:  0,
						},
					},
				},
			},
		},
	}

	err = gs.BindGameEvent(event)
	if err != nil {
		panic(err)
	}

}
