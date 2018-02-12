# gdax
gdax is a command line tool that allows the user to get snapshot or streaming quotes from Coinbase's GDAX exchange

**Note:** This package needs to be updated with new cryptocurrency products recently added to the exchange(bitcoin cash)

## Installation
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