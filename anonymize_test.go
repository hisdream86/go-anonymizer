package goanonymizer

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

type Example struct {
	// Anonymize with asterisk (*)
	MaskedField string `anonymize:"asterisk"`

	// Anonymize to empty string ("")
	EmptyField string `anonymize:"empty"`

	// Anonymize with custom handler
	CustomField string `anonymize:"mysha256"`

	// Don't anonymize
	PlainField string

	NestedField

	ArrayField []string `anonymize:"asterisk"`
}

type NestedField struct {
	MaskedField string `anonymize:"asterisk"`
}

var example = Example{
	MaskedField: "MaskedField",
	EmptyField:  "EmptystringField",
	CustomField: "CustomField",
	PlainField:  "PlainField",
	NestedField: NestedField{
		MaskedField: "NestedMaskedField",
	},
	ArrayField: []string{"Field 1", "Field 2"},
}

func TestAnonymization(t *testing.T) {
	target := example
	oMaskedField := target.MaskedField
	oEmptyField := target.EmptyField
	oCustomField := target.CustomField
	oPlainField := target.PlainField
	oNestedMaskedField := target.NestedField.MaskedField

	err := Anonymize(&target)
	if err != nil {
		t.Error(err)
	}

	if target.MaskedField != asteriskReplacer(oMaskedField) {
		t.Error("target.CustomField != sha256Replacer(oCustomField)")
	}

	if target.EmptyField != emptystringReplacer(oEmptyField) {
		t.Error("target.EmptyField != emptystringReplacer(oEmptyField)")
	}

	if target.CustomField != oCustomField {
		t.Error("target.CustomField != oCustomField")
	}

	if target.PlainField != oPlainField {
		t.Error("target.PlainField != oPlainField")
	}

	if target.NestedField.MaskedField != asteriskReplacer(oNestedMaskedField) {
		t.Error("target.NestedField.MaskedField != oNestedMaskedField")
	}
}

func TestTargetIsNotPointer(t *testing.T) {
	target := example
	err := Anonymize(target)
	if err.Error() != "target is not pointer" {
		t.Error(err)
	}
}

func TestCustomReplacer(t *testing.T) {
	target := example
	oCustomField := target.CustomField

	sha256Replacer := func(source string) string {
		h := sha256.New()
		h.Write([]byte(source))
		return hex.EncodeToString(h.Sum(nil))
	}

	AddCustomReplacer("mysha256", sha256Replacer)

	err := Anonymize(&target)
	if err != nil {
		t.Error(err)
	}

	RemoveCustomReplacer("mysha256")

	if target.CustomField != sha256Replacer(oCustomField) {
		t.Error("target.CustomField != sha256Replacer(oCustomField)")
	}
}
