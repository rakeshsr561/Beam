package managers

import (
	"testing"
)

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func TestRemoveDuplicateUrls(t *testing.T) {
	input := []string{"abc.com", "amp.com", "abc.com"}
	output := RemoveDuplicateUrls(input)
	actual := len(output) == 2
	actual = actual && contains(output, "abc.com")
	actual = actual && contains(output, "amp.com")

	expected := true
	if actual != expected {
		t.Errorf("Expected (%t) is not same as"+
			" actual (%t)", expected, actual)
	}
}
