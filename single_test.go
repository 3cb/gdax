package main

import (
	"reflect"
	"testing"
)

func Test_quoteSingle(t *testing.T) {
	type args struct {
		state map[string]Product
		pairs []string
		max   *MaxLengths
	}
	tests := []struct {
		name string
		args args
		want *FmtPrint
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quoteSingle(tt.args.state, tt.args.pairs, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quoteSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processTradeCh(t *testing.T) {
	max := &MaxLengths{
		Price:  5,
		Delta:  6,
		Size:   9,
		Bid:    3,
		Ask:    3,
		High:   4,
		Low:    3,
		Open:   4,
		Volume: 6,
	}
	initState := make(map[string]Product)
	initState["BTC-USD"] = Product{ID: "BTC-USD"}
	nonEmptyState := make(map[string]Product)
	nonEmptyState["BTC-USD"] = Product{
		ID:     "BTC-USD",
		Bid:    "8888.12",
		Ask:    "8888.15",
		High:   "9000.56",
		Low:    "8500.56",
		Open:   "8700.98",
		Volume: "22356",
	}
	type args struct {
		state map[string]Product
		max   *MaxLengths
		trade *Product
	}
	tc := []struct {
		name     string
		args     args
		expected Product
	}{
		{
			"test product: init state, default maxLengths",
			args{
				state: initState,
				max:   max,
				trade: &Product{
					ID:    "BTC-USD",
					Price: "8456.12345678",
					Size:  "0.12345678",
				},
			},
			Product{
				ID:    "BTC-USD",
				Price: "8456.12",
				Size:  "0.12345678",
			},
		},
		{
			"test product: non-empty state, default maxLengths",
			args{
				state: nonEmptyState,
				max:   max,
				trade: &Product{
					ID:    "BTC-USD",
					Price: "8456.12345678",
					Size:  "0.12345678",
				},
			},
			Product{
				ID:     "BTC-USD",
				Price:  "8456.12",
				Size:   "0.12345678",
				Bid:    "8888.12",
				Ask:    "8888.15",
				High:   "9000.56",
				Low:    "8500.56",
				Open:   "8700.98",
				Volume: "22356",
			},
		},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := processTradeCh(c.args.state, c.args.max, c.args.trade)
			if got != c.expected {
				t.Errorf("expected %v; got %v", c.expected, got)
			}
		})
	}
}

// func Test_processStatsCh(t *testing.T) {
// 	max := &MaxLengths{
// 		Price:  5,
// 		Delta:  6,
// 		Size:   9,
// 		Bid:    3,
// 		Ask:    3,
// 		High:   4,
// 		Low:    3,
// 		Open:   4,
// 		Volume: 6,
// 	}
// 	initState := make(map[string]Product)
// 	initState["BTC-USD"] = Product{ID: "BTC-USD"}
// 	nonEmptyState := make(map[string]Product)
// 	nonEmptyState["BTC-USD"] = Product{
// 		ID:    "BTC-USD",
// 		Price: "8888.12",
// 		Size:  "0.78945612",
// 		Bid:   "8888.12",
// 		Ask:   "8888.15",
// 	}
// 	type args struct {
// 		state map[string]Product
// 		max   *MaxLengths
// 		stats *Stats
// 	}
// 	tc := []struct {
// 		name string
// 		args args
// 		expected Product
// 	}{
// 		{}
// 	}
// }
