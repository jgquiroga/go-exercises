package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Julian"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Julian")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Julian")`)
	}
}

func TestEmptyName(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q`, msg)
	}
}
