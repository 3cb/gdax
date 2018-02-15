package main

import (
	"testing"
)

func Test_rndPrice(t *testing.T) {
	tc := []struct {
		input    string
		expected string
	}{
		{"0.00040356", "0.00040"},
		{"0.00658766", "0.00659"},
		{"8623.12345678", "8623.12"},
		{"8623.12545678", "8623.13"},
		{"8623.12845678", "8623.13"},
	}

	for _, c := range tc {
		actual := rndPrice(c.input)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func Test_rndSize(t *testing.T) {
	tc := []struct {
		input    string
		expected string
	}{
		{"0.123456789", "0.12345679"},
		{"0.123456783", "0.12345678"},
		{"10.123456783", "10.12345678"},
		{"10.123456787", "10.12345679"},
	}

	for _, c := range tc {
		actual := rndSize(c.input)
		if actual != c.expected {
			t.Errorf("expected %v;got %v", c.expected, actual)
		}
	}
}

func Test_rndVol(t *testing.T) {
	tc := []struct {
		input    string
		expected string
	}{
		{"333.12345678", "333"},
		{"333.56789238", "334"},
		{"333.89745623", "334"},
	}

	for _, c := range tc {
		actual := rndVol(c.input)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func Test_setDelta(t *testing.T) {
	tc := []struct {
		price    string
		open     string
		expected string
	}{
		{"8888.12345678", "8700.12345678", "2.16%"},
		{"10000.12345678", "7500.12345678", "33.33%"},
		{"6000.00000000", "7500.12345678", "-20.00%"},
	}

	for _, c := range tc {
		actual := setDelta(c.price, c.open)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}
