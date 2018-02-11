package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	stream := flag.Bool("s", false, "stream cryptocurrency prices")
	flag.Parse()

	pairs := []string{"BTC-USD", "BTC-EUR", "BTC-GBP", "ETH-USD", "ETH-BTC", "ETH-EUR", "LTC-USD", "LTC-BTC", "LTC-EUR"}

	if !*stream {
		quotes := make(map[string]GDAXTrade)
		quoteSingle(pairs, quotes)
	} else {
		clear()
		quoteStream(pairs)
	}
}

// clears terminal for linux, mac, and windows
func clear() {
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

func round(price string) string {
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

func print(spaces map[string]string, quotes map[string]GDAXTrade) {
	println("\nLive GDAX Quotes:")
	fmt.Printf("\nBTC/USD:%v%v\n", spaces["BTC-USD"], quotes["BTC-USD"].Price)
	fmt.Printf("\nBTC/EUR:%v%v\n", spaces["BTC-EUR"], quotes["BTC-EUR"].Price)
	fmt.Printf("\nBTC/GBP:%v%v\n", spaces["BTC-GBP"], quotes["BTC-GBP"].Price)
	fmt.Printf("\nETH/USD:%v%v\n", spaces["ETH-USD"], quotes["ETH-USD"].Price)
	fmt.Printf("\nETH/BTC:%v%v\n", spaces["ETH-BTC"], quotes["ETH-BTC"].Price)
	fmt.Printf("\nETH/EUR:%v%v\n", spaces["ETH-EUR"], quotes["ETH-EUR"].Price)
	fmt.Printf("\nLTC/USD:%v%v\n", spaces["LTC-USD"], quotes["LTC-USD"].Price)
	fmt.Printf("\nLTC/BTC:%v%v\n", spaces["LTC-BTC"], quotes["LTC-BTC"].Price)
	fmt.Printf("\nLTC/EUR:%v%v\n", spaces["LTC-EUR"], quotes["LTC-EUR"].Price)
}
