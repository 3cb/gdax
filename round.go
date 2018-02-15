// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// rndPrice rounds prices to 2 or 5 decimal places
func rndPrice(price string) string {
	num, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return "-----"
	}
	if num >= 10 {
		num = float64(int64(num*100+0.5)) / 100
		return fmt.Sprintf("%.2f", num)

	}
	num = float64(int64(num*100000+0.5)) / 100000
	return fmt.Sprintf("%.5f", num)
}

// rndSize rounds last size data to 8 decimal places
func rndSize(size string) string {
	num, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return "-----"
	}
	num = float64(int64(num*100000000+0.5)) / 100000000
	return fmt.Sprintf("%.8f", num)
}

// rndVol rounds volume data to the nearest whole number
func rndVol(vol string) string {
	num, err := strconv.ParseFloat(vol, 64)
	if err != nil {
		return "-----"
	}
	return fmt.Sprint(int64(num + 0.5))
}

// setDelta returns price delta rounded to two decimal places as a string
// returns the print color based on the delta
func setDelta(price string, open string) string {
	p, _ := strconv.ParseFloat(price, 64)
	o, _ := strconv.ParseFloat(open, 64)
	delta := ((p - o) / o) * 100
	buf := bytes.Buffer{}
	buf.WriteString(strconv.FormatFloat(delta, 'f', 2, 64))
	buf.WriteString("%")
	return buf.String()
}

// SetColor uses the delta filed of Product type to set color either red or green
func SetColor(delta string) *color.Color {
	c := &color.Color{}
	slice := strings.Split(delta, "%")
	d, _ := strconv.ParseFloat(slice[0], 64)
	if d > 0 {
		c.Add(color.FgGreen, color.Bold)
	} else {
		c.Add(color.FgRed, color.Bold)
	}
	return c
}
