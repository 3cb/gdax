# gdax
gdax is a command line tool that allows the user to get snapshot or streaming quotes from Coinbase's GDAX exchange: https://www.gdax.com/trade/BTC-USD

![Diagram](https://images2.imgbox.com/f0/a6/H0bIfdx5_o.png?download=true)

API Docs: https://docs.gdax.com/

## Installation
Make sure you have Go installed: https://golang.org/doc/install

To install gdax run:
```
go get github.com/3cb/gdax
```
Make sure your `PATH` includes the `$GOPATH/bin` directory so the commands can be used:
```
export PATH=$PATH:$GOPATH/bin
```

## Usage
For a snapshot quote simply run the binary from the command line:
```
gdax
```
For streaming quotes through a websocket connection run with the `-s` flag:
```
gdax -s
```

**That's it!**