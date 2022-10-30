package utils

import (
	"fmt"
	"testing"
)

func TestValidUrl(t *testing.T) {
	actual := ValidateUrl("http://www.google.com")
	fmt.Print(actual)
	expected := true
	if actual != expected {
		t.Errorf("Expected (%t) is not same as"+
			" actual (%t)", expected, actual)
	}
}

func TestInValidUrl(t *testing.T) {
	actual := ValidateUrl("http/www.google.com")
	fmt.Print(actual)
	expected := false
	if actual != expected {
		t.Errorf("Expected (%t) is not same as"+
			" actual (%t)", expected, actual)
	}
}
