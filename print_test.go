package main

import (
	"testing"
)

func TestSetMax(t *testing.T) {
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
		actual := SetMax(c.current, c.testLen)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func TestSetSpc(t *testing.T) {
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
		actual := SetSpc(c.max, c.original)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func TestSetSpcStrm(t *testing.T) {
	tc := []struct {
		max      int
		original string
		expected string
	}{
		{10, "four", "           four"},
		{5, "eleven", "    eleven"},
	}

	for _, c := range tc {
		actual := SetSpcStrm(c.max, c.original)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}

func TestSetHdr(t *testing.T) {
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
		actual := SetHdr(c.header, c.total)
		if actual != c.expected {
			t.Errorf("expected %v; got %v", c.expected, actual)
		}
	}
}
