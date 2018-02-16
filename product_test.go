package main

import (
	"testing"
)

func Test_setSpc(t *testing.T) {
	tc := []struct {
		max      int
		original string
		expected string
	}{
		{10, "four", "           four"},
		{15, "four", "                four"},
		{4, "four", "     four"},
	}

	for _, c := range tc {
		actual := setSpc(c.max, c.original)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func Test_setSpcStrm(t *testing.T) {
	tc := []struct {
		max      int
		original string
		expected string
	}{
		{10, "four", "           four"},
		{5, "eleven", "    eleven"},
	}

	for _, c := range tc {
		actual := setSpcStrm(c.max, c.original)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func Test_fmtRow(t *testing.T) {
	product := Product{
		ID:     "BTC-USD",
		Price:  "       8888.13",
		Size:   "     0.12345678",
		Change: "         1.02%",
		Bid:    "         8888.13",
		Ask:    "         8888.16",
		High:   "         9000.20",
		Low:    "         8888.13",
		Volume: "         5000",
	}
	tc := []struct {
		pair     Product
		expected string
	}{
		{
			pair:     product,
			expected: "\n BTC/USD       8888.13     0.12345678         1.02%         8888.13         8888.16         9000.20         8888.13         5000",
		},
	}

	for _, c := range tc {
		got := (c.pair).fmtRow()
		if got != c.expected {
			t.Errorf("expected %v; got %v", c.expected, got)
		}
	}
}
