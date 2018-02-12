package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

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

func rndSize(size string) string {
	num, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return "-----"
	}
	num = float64(int64(num*100000000+0.5)) / 100000000
	return fmt.Sprintf("%.8f", num)
}

func rndVol(vol string) string {
	num, err := strconv.ParseFloat(vol, 64)
	if err != nil {
		return "-----"
	}
	return fmt.Sprint(int64(num + 0.5))
}

// returns price delta rounded to two decimal places as a string
// returns the print color based on the delta
func setDeltaColor(price string, open string) (string, *color.Color) {
	p, _ := strconv.ParseFloat(price, 64)
	o, _ := strconv.ParseFloat(open, 64)
	delta := ((p - o) / o) * 100
	c := &color.Color{}
	if delta > 0 {
		c.Add(color.FgGreen, color.Bold)
	} else {
		c.Add(color.FgRed, color.Bold)
	}
	buf := bytes.Buffer{}
	buf.WriteString(strconv.FormatFloat(delta, 'f', 2, 64))
	buf.WriteString("%")
	return buf.String(), c
}
