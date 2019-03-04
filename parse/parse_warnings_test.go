package parse

import "testing"

func TestParseSingleWarning(t *testing.T) {
	warnings := parseWarnings("Warning: resource_name: <warning detail>")

	expectInt(t, 1, len(warnings))
	expectString(t, "resource_name", warnings[0].Id.Name)
	expectString(t, " <warning detail>", warnings[0].Detail)
}

func TestParseMultipleWarnings(t *testing.T) {
	warnings := parseWarnings("Warning: r1: w1\nWarning: r2: w2\nWarning: r3: w3")

	expectInt(t, 3, len(warnings))

	expectString(t, "r1", warnings[0].Id.Name)
	expectString(t, " w1", warnings[0].Detail)

	expectString(t, "r2", warnings[1].Id.Name)
	expectString(t, " w2", warnings[1].Detail)

	expectString(t, "r3", warnings[2].Id.Name)
	expectString(t, " w3", warnings[2].Detail)
}

func TestParseNoWarnings(t *testing.T) {
	warnings := parseWarnings("Here are some things that are NOT warnings")

	expectInt(t, 0, len(warnings))
}
