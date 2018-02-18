// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"math"
	"strings"
)

// MaxLengths is used to format spaces for printing
type MaxLengths struct {
	Price  int
	Change int
	Size   int
	Bid    int
	Ask    int
	High   int
	Low    int
	Open   int
	Volume int
}

// setMax checks if max length for field needs to be reset
func setMax(currMax int, testLen int) int {
	if testLen > currMax {
		return testLen
	}
	return currMax
}

// FmtPrint contains structure for printing so it only has to be calculated once
type FmtPrint struct {
	Title   string
	Headers string
	Footer  string
}

// fmtColHdr formats column headers and returns a string
func fmtColHdr(max *MaxLengths) string {
	b := strings.Builder{}
	b.WriteString(" Product")
	b.WriteString(setSpc(max.Price, "Price"))
	b.WriteString(setSpc(max.Size, "Last Size"))
	b.WriteString(setSpc(max.Change, "Change"))
	b.WriteString(setSpc(max.Bid, "Bid"))
	b.WriteString(setSpc(max.Ask, "Ask"))
	b.WriteString(setSpc(max.High, "High"))
	b.WriteString(setSpc(max.Low, "Low"))
	b.WriteString(setSpc(max.Volume, "Volume"))
	b.WriteString(" ")
	return b.String()
}

// fmtTitle returns header(or footer with empty string parameter) centered based on total max length, left margin(1), product column(7), and spaces between columns(48)
func fmtTitle(title string, total int) string {
	b := strings.Builder{}
	var lMargin, rMargin int
	whitespace := total - len(title)

	if whitespace%2 == 0 {
		lMargin = whitespace / 2
		rMargin = lMargin
	} else {
		lMargin = int(math.Floor(float64(whitespace) / 2))
		rMargin = lMargin + 1
	}

	for i := 0; i < lMargin; i++ {
		b.WriteString(" ")
	}
	b.WriteString(title)
	for i := 0; i < rMargin; i++ {
		b.WriteString(" ")
	}
	return b.String()
}
