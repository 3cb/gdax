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
