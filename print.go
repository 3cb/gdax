package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
	fmt.Printf("\n%v\n", setHdr("Live GDAX Quotes:", m.getTotal()))
	fmt.Printf("\n    Price%v%v%v%v%v%v%v", setSpc(m.Price, "Price"), setSpc(m.Size, "Last Size"), setSpc(m.Bid, "Bid"), setSpc(m.Ask, "Ask"), setSpc(m.High, "High"), setSpc(m.Low, "Low"), setSpc(m.Volume, "Volume"))

	fmt.Printf("\n  BTC/USD%v%v%v%v%v%v%v\n", state["BTC-USD"].Price, state["BTC-USD"].Size, state["BTC-USD"].Bid, state["BTC-USD"].Ask, state["BTC-USD"].High, state["BTC-USD"].Low, state["BTC-USD"].Volume)
	fmt.Printf("\n  BTC/EUR%v%v%v%v%v%v%v\n", state["BTC-EUR"].Price, state["BTC-EUR"].Size, state["BTC-EUR"].Bid, state["BTC-EUR"].Ask, state["BTC-EUR"].High, state["BTC-EUR"].Low, state["BTC-EUR"].Volume)
	fmt.Printf("\n  BTC/GBP%v%v%v%v%v%v%v\n", state["BTC-GBP"].Price, state["BTC-GBP"].Size, state["BTC-GBP"].Bid, state["BTC-GBP"].Ask, state["BTC-GBP"].High, state["BTC-GBP"].Low, state["BTC-GBP"].Volume)
	fmt.Printf("\n  ETH/USD%v%v%v%v%v%v%v\n", state["ETH-USD"].Price, state["ETH-USD"].Size, state["ETH-USD"].Bid, state["ETH-USD"].Ask, state["ETH-USD"].High, state["ETH-USD"].Low, state["ETH-USD"].Volume)
	fmt.Printf("\n  ETH/BTC%v%v%v%v%v%v%v\n", state["ETH-BTC"].Price, state["ETH-BTC"].Size, state["ETH-BTC"].Bid, state["ETH-BTC"].Ask, state["ETH-BTC"].High, state["ETH-BTC"].Low, state["ETH-BTC"].Volume)
	fmt.Printf("\n  ETH/EUR%v%v%v%v%v%v%v\n", state["ETH-EUR"].Price, state["ETH-EUR"].Size, state["ETH-EUR"].Bid, state["ETH-EUR"].Ask, state["ETH-EUR"].High, state["ETH-EUR"].Low, state["ETH-EUR"].Volume)
	fmt.Printf("\n  LTC/USD%v%v%v%v%v%v%v\n", state["LTC-USD"].Price, state["LTC-USD"].Size, state["LTC-USD"].Bid, state["LTC-USD"].Ask, state["LTC-USD"].High, state["LTC-USD"].Low, state["LTC-USD"].Volume)
	fmt.Printf("\n  LTC/BTC%v%v%v%v%v%v%v\n", state["LTC-BTC"].Price, state["LTC-BTC"].Size, state["LTC-BTC"].Bid, state["LTC-BTC"].Ask, state["LTC-BTC"].High, state["LTC-BTC"].Low, state["LTC-BTC"].Volume)
	fmt.Printf("\n  LTC/EUR%v%v%v%v%v%v%v\n", state["LTC-EUR"].Price, state["LTC-EUR"].Size, state["LTC-EUR"].Bid, state["LTC-EUR"].Ask, state["LTC-EUR"].High, state["LTC-EUR"].Low, state["LTC-EUR"].Volume)
}

// checks if max length for field needs to be reset
func getMax(currMax int, testLen int) int {
	if testLen > currMax {
		return testLen
	}
	return currMax
}

// adds space in front of string based on max length
func setSpc(max int, orig string) string {
	buf := bytes.Buffer{}
	buf.WriteString("      ")
	diff := max - len(orig)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			buf.WriteString(" ")
		}
	}
	buf.WriteString(orig)
	return buf.String()
}

// returns header centered based on total max length, left margin, product column, and spaces between columns(28)
func setHdr(header string, total int) string {
	buf := bytes.Buffer{}
	margin := int(9+28+total-len(header)) / 2
	for i := 0; i < margin; i++ {
		buf.WriteString(" ")
	}
	buf.WriteString(header)
	return buf.String()
}
