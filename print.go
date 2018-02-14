// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"

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

	c.Print(format.Title)
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

// FmtColHdr formats column headers and returns a string
func FmtColHdr(max *MaxLengths) string {
	buf := bytes.Buffer{}
	buf.WriteString("\n Product")
	buf.WriteString(SetSpc(max.Price, "Price"))
	buf.WriteString(SetSpc(max.Size, "Last Size"))
	buf.WriteString(SetSpc(max.Delta, "Change"))
	buf.WriteString(SetSpc(max.Bid, "Bid"))
	buf.WriteString(SetSpc(max.Ask, "Ask"))
	buf.WriteString(SetSpc(max.High, "High"))
	buf.WriteString(SetSpc(max.Low, "Low"))
	buf.WriteString(SetSpc(max.Volume, "Volume"))
	buf.WriteString(" ")
	return buf.String()
}

// FmtRow formats all Product data into a single row so it can be printed
func FmtRow(pair Product) string {
	buf := bytes.Buffer{}
	buf.WriteString("\n ")
	buf.WriteString(strings.Join(strings.Split(pair.ID, "-"), "/"))
	buf.WriteString(pair.Price)
	buf.WriteString(pair.Size)
	buf.WriteString(pair.Delta)
	buf.WriteString(pair.Bid)
	buf.WriteString(pair.Ask)
	buf.WriteString(pair.High)
	buf.WriteString(pair.Low)
	buf.WriteString(pair.Volume)
	return buf.String()
}

// SetMax checks if max length for field needs to be reset
func SetMax(currMax int, testLen int) int {
	if testLen > currMax {
		return testLen
	}
	return currMax
}

// SetSpc adds space in front of string based on max length for single quote
func SetSpc(max int, orig string) string {
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

// SetSpcStrm adds spaces to string according to max lengths calculated in single quote
func SetSpcStrm(max int, orig string) string {
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

// FmtTitle returns header(or footer with empty string parameter) centered based on total max length, left margin(1), product column(7), and spaces between columns(48)
func FmtTitle(title string, total int) string {
	buf := bytes.Buffer{}
	var lMargin, rMargin int
	whitespace := total - len(title)

	if whitespace%2 == 0 {
		lMargin = whitespace / 2
		rMargin = lMargin
	} else {
		lMargin = (whitespace - 1) / 2
		rMargin = lMargin + 1
	}

	for i := 0; i < lMargin; i++ {
		buf.WriteString(" ")
	}
	buf.WriteString(title)
	for i := 0; i < rMargin; i++ {
		buf.WriteString(" ")
	}
	return buf.String()
}
