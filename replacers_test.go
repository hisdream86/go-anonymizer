package goanonymizer

import (
	"testing"
)

func TestEmptyStringReplacer(t *testing.T) {
	source := "My Test String"
	if emptystringReplacer(source) != "" {
		t.Error("asteriskReplacer(source) != \"\"")
	}
}

func TestAsteriskReplacer(t *testing.T) {
	source := "My Test String"
	if asteriskReplacer(source) != "**************" {
		t.Error("asteriskReplacer(source) != \"**************\"")
	}
}
