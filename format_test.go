package main

import "testing"

func Test_setMax(t *testing.T) {
	tc := []struct {
		current  int
		testLen  int
		expected int
	}{
		{10, 8, 10},
		{8, 10, 10},
		{10, 20, 20},
	}

	for _, c := range tc {
		actual := setMax(c.current, c.testLen)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func Test_fmtTitle(t *testing.T) {
	tc := []struct {
		header   string
		total    int
		expected string
	}{
		{"four", 24, "          four          "},
		{"four", 23, "         four          "},
		{"one", 24, "          one           "},
		{"one", 23, "          one          "},
	}

	for _, c := range tc {
		actual := fmtTitle(c.header, c.total)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func Test_fmtColHdr(t *testing.T) {
	max := &MaxLengths{
		Price:  8,
		Change: 7,
		Size:   10,
		Bid:    8,
		Ask:    8,
		High:   8,
		Low:    8,
		Open:   8,
		Volume: 7,
	}
	tc := []struct {
		max      *MaxLengths
		expected string
	}{
		{
			max:      max,
			expected: " Product        Price      Last Size      Change          Bid          Ask         High          Low      Volume ",
		},
	}

	for _, c := range tc {
		got := fmtColHdr(c.max)
		if got != c.expected {
			t.Errorf("\nexpected %v;got\n%v", c.expected, got)
		}
	}
}
