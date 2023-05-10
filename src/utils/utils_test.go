package utils

import (
	"strings"
	"testing"
)

func TestSplitAndRemoveSpaces(t *testing.T) {
	for _, word := range []string{"a", "a,b", "a, b", "a, b,", ",a, b,", ",a, b, "} {
		actual := SplitAndRemoveSpaces(word)
		for _, s := range actual {
			if strings.Contains(s, " ") {
				t.Errorf("Expected no spaces, got %s", s)
			}

			if s == "" {
				t.Errorf("Expected no empty strings, got %s", s)
			}
		}
	}
}
