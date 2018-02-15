// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

// clears terminal for linux, mac, and windows
func clearScr() {
	switch sys := runtime.GOOS; sys {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		return
	}
}

// prints price quotes to terminal
func print(state map[string]Product, format *FmtPrint) {
	c := color.New(color.FgBlack, color.BgWhite)

	c.Println(format.Title)
	c.Print(format.Headers)
	state["BTC-USD"].Color.Println(state["BTC-USD"].Row)
	state["BTC-EUR"].Color.Println(state["BTC-EUR"].Row)
	state["BTC-GBP"].Color.Println(state["BTC-GBP"].Row)
	println()
	state["BCH-USD"].Color.Println(state["BCH-USD"].Row)
	state["BCH-BTC"].Color.Println(state["BCH-BTC"].Row)
	state["BCH-EUR"].Color.Println(state["BCH-EUR"].Row)
	println()
	state["ETH-USD"].Color.Println(state["ETH-USD"].Row)
	state["ETH-BTC"].Color.Println(state["ETH-BTC"].Row)
	state["ETH-EUR"].Color.Println(state["ETH-EUR"].Row)
	println()
	state["LTC-USD"].Color.Println(state["LTC-USD"].Row)
	state["LTC-BTC"].Color.Println(state["LTC-BTC"].Row)
	state["LTC-EUR"].Color.Println(state["LTC-EUR"].Row)
	c.Println(format.Footer)
}
