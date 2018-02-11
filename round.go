package main

import (
	"fmt"
	"strconv"
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
