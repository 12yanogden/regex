package regex

import (
	"testing"
)

func TestPretty(t *testing.T) {
	regex := "(this(that))"
	expected := `(
	this(
		that
	)
)`
	actual := Pretty(regex)

	if expected != actual {
		t.Fatal("\nExpected: " + expected + "\nActual: " + actual)
	}
}

func TestReversibility(t *testing.T) {
	expected := "^(([[:alnum:]]| )*)?=?$"
	actual := Ugly(Pretty(expected))

	if expected != actual {
		t.Fatal("\nExpected: " + expected + "\nActual: " + actual)
	}
}

func TestPrettyWithCommentsToRegex(t *testing.T) {
	pretty := `(
	// comment
	this
	(
			that	// comment
	)
)`
	expected := "(this(that))"
	actual := Ugly(pretty)

	if expected != actual {
		t.Fatal("\nExpected: " + expected + "\nActual: " + actual)
	}
}
