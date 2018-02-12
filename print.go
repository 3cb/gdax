// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"bytes"
	"fmt"
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
func print(state map[string]Product, m *MaxLengths) {
	c := color.New(color.FgBlack, color.BgWhite)
	headers := fmt.Sprintf("\n Product%v%v%v%v%v%v%v%v ", setSpc(m.Price, "Price"), setSpc(m.Size, "Last Size"), setSpc(m.Delta, "Change"), setSpc(m.Bid, "Bid"), setSpc(m.Ask, "Ask"), setSpc(m.High, "High"), setSpc(m.Low, "Low"), setSpc(m.Volume, "Volume"))

	c.Printf("%v", setHdr("GDAX Cryptocurrency Exchange", len(headers)))
	c.Print(headers)

	state["BTC-USD"].Color.Printf("\n BTC/USD%v%v%v%v%v%v%v%v", state["BTC-USD"].Price, state["BTC-USD"].Size, state["BTC-USD"].Delta, state["BTC-USD"].Bid, state["BTC-USD"].Ask, state["BTC-USD"].High, state["BTC-USD"].Low, state["BTC-USD"].Volume)
	state["BTC-EUR"].Color.Printf("\n BTC/EUR%v%v%v%v%v%v%v%v", state["BTC-EUR"].Price, state["BTC-EUR"].Size, state["BTC-EUR"].Delta, state["BTC-EUR"].Bid, state["BTC-EUR"].Ask, state["BTC-EUR"].High, state["BTC-EUR"].Low, state["BTC-EUR"].Volume)
	state["BTC-GBP"].Color.Printf("\n BTC/GBP%v%v%v%v%v%v%v%v\n", state["BTC-GBP"].Price, state["BTC-GBP"].Size, state["BTC-GBP"].Delta, state["BTC-GBP"].Bid, state["BTC-GBP"].Ask, state["BTC-GBP"].High, state["BTC-GBP"].Low, state["BTC-GBP"].Volume)

	state["BCH-USD"].Color.Printf("\n BCH/USD%v%v%v%v%v%v%v%v", state["BCH-USD"].Price, state["BCH-USD"].Size, state["BCH-USD"].Delta, state["BCH-USD"].Bid, state["BCH-USD"].Ask, state["BCH-USD"].High, state["BCH-USD"].Low, state["BCH-USD"].Volume)
	state["BCH-BTC"].Color.Printf("\n BCH/BTC%v%v%v%v%v%v%v%v", state["BCH-BTC"].Price, state["BCH-BTC"].Size, state["BCH-BTC"].Delta, state["BCH-BTC"].Bid, state["BCH-BTC"].Ask, state["BCH-BTC"].High, state["BCH-BTC"].Low, state["BCH-BTC"].Volume)
	state["BCH-EUR"].Color.Printf("\n BCH/EUR%v%v%v%v%v%v%v%v\n", state["BCH-EUR"].Price, state["BCH-EUR"].Size, state["BCH-EUR"].Delta, state["BCH-EUR"].Bid, state["BCH-EUR"].Ask, state["BCH-EUR"].High, state["BCH-EUR"].Low, state["BCH-EUR"].Volume)

	state["ETH-USD"].Color.Printf("\n ETH/USD%v%v%v%v%v%v%v%v", state["ETH-USD"].Price, state["ETH-USD"].Size, state["ETH-USD"].Delta, state["ETH-USD"].Bid, state["ETH-USD"].Ask, state["ETH-USD"].High, state["ETH-USD"].Low, state["ETH-USD"].Volume)
	state["ETH-BTC"].Color.Printf("\n ETH/BTC%v%v%v%v%v%v%v%v", state["ETH-BTC"].Price, state["ETH-BTC"].Size, state["ETH-BTC"].Delta, state["ETH-BTC"].Bid, state["ETH-BTC"].Ask, state["ETH-BTC"].High, state["ETH-BTC"].Low, state["ETH-BTC"].Volume)
	state["ETH-EUR"].Color.Printf("\n ETH/EUR%v%v%v%v%v%v%v%v\n", state["ETH-EUR"].Price, state["ETH-EUR"].Size, state["ETH-EUR"].Delta, state["ETH-EUR"].Bid, state["ETH-EUR"].Ask, state["ETH-EUR"].High, state["ETH-EUR"].Low, state["ETH-EUR"].Volume)

	state["LTC-USD"].Color.Printf("\n LTC/USD%v%v%v%v%v%v%v%v", state["LTC-USD"].Price, state["LTC-USD"].Size, state["LTC-USD"].Delta, state["LTC-USD"].Bid, state["LTC-USD"].Ask, state["LTC-USD"].High, state["LTC-USD"].Low, state["LTC-USD"].Volume)
	state["LTC-BTC"].Color.Printf("\n LTC/BTC%v%v%v%v%v%v%v%v", state["LTC-BTC"].Price, state["LTC-BTC"].Size, state["LTC-BTC"].Delta, state["LTC-BTC"].Bid, state["LTC-BTC"].Ask, state["LTC-BTC"].High, state["LTC-BTC"].Low, state["LTC-BTC"].Volume)
	state["LTC-EUR"].Color.Printf("\n LTC/EUR%v%v%v%v%v%v%v%v\n", state["LTC-EUR"].Price, state["LTC-EUR"].Size, state["LTC-EUR"].Delta, state["LTC-EUR"].Bid, state["LTC-EUR"].Ask, state["LTC-EUR"].High, state["LTC-EUR"].Low, state["LTC-EUR"].Volume)

	c.Printf("%v\n", setHdr("", len(headers)))

}

// checks if max length for field needs to be reset
func getMax(currMax int, testLen int) int {
	if testLen > currMax {
		return testLen
	}
	return currMax
}

// adds space in front of string based on max length for single quote
func setSpc(max int, orig string) string {
	buf := bytes.Buffer{}
	buf.WriteString("     ")
	diff := max - len(orig)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			buf.WriteString(" ")
		}
	}
	buf.WriteString(orig)
	return buf.String()
}

func setSpcStrm(max int, orig string) string {
	buf := bytes.Buffer{}
	diff := 5 + max - len(orig)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			buf.WriteString(" ")
		}
	}
	buf.WriteString(orig)
	return buf.String()
}

// returns header(or footer with empty string parameter) centered based on total max length, left margin(1), product column(7), and spaces between columns(48)
func setHdr(header string, total int) string {
	buf := bytes.Buffer{}
	var lMargin, rMargin int

	line := total - len(header)
	lMargin = line / 2
	rMargin = lMargin
	for i := 0; i < lMargin; i++ {
		buf.WriteString(" ")
	}
	buf.WriteString(header)
	for i := 0; i < rMargin; i++ {
		buf.WriteString(" ")
	}
	return buf.String()
}
