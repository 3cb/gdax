package main

import (
	"testing"
)

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
