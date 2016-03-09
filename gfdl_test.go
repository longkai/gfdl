package main

import (
	"testing"
)

func TestShortName(t *testing.T) {
	cases := []string{
		"http://a/b/c.go",
		"a/b/c.go",
		"c.go",
		"",
	}
	answers := []string{
		"c.go",
		"c.go",
		"c.go",
		"",
	}

	for i, s := range cases {
		if v := shortName(s); v != answers[i] {
			t.Errorf("test %s, want %s, got %s\n", s, answers[i], v)
		}
	}
}

func TestRegex(t *testing.T) {
	cases := []string{
		"http://a/b/c.go",
		"https://a.com",
		"ftp://a/b/c.go",
		"",
	}
	answers := []bool{
		true,
		true,
		false,
		false,
	}
	for i, s := range cases {
		if v := regex.FindString(s) != ""; v != answers[i] {
			t.Errorf("test %s, want %b, got %v\n", s, answers[i], v)
		}
	}
}
