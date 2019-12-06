# Go Anonymizer
Simple anonymization tool for de-identifying your structurized data. You can easily anonymize your data with *Struct Tag* `anonymize:"{replacer}"`.

*Note: only string values are supported*

# Installation
    $ go get github.com/hisdream86/go-anonymizer

# Example
```go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	anonymizer "github.com/hisdream86/go-anonymizer"
)

type Example struct {
	// Anonymize with asterisk (*)
	MaskedField string `anonymize:"asterisk"`
	// Anonymize to empty string ("")
	EmptyField string `anonymize:"empty"`
	// Anonymize with custom handler
	CustomField string `anonymize:"mysha256"`
	// Anoymize nested Field
	NestedField struct {
		MaskedField string `anonymize:"asterisk"`
	}
	// Anonymize array Field
	ArrayField []string `anonymize:"asterisk"`
	// Don't anonymize
	PlainField string
}

type NestedField struct {
	MaskedField string `anonymize:"asterisk"`
}

func main() {
	var target = Example{
		MaskedField: "MaskedField",
		EmptyField:  "EmptystringField",
		CustomField: "CustomField",
		PlainField:  "PlainField",
		NestedField: NestedField{
			MaskedField: "NestedMaskedField",
		},
		ArrayField: []string{"Field 1", "Field 2"},
	}

	// Register custom replacer for anonymizing data with SHA256 hashing
	anonymizer.AddCustomReplacer("mysha256", func(source string) string {
		h := sha256.New()
		h.Write([]byte(source))
		return hex.EncodeToString(h.Sum(nil))
	})

	if err := anonymizer.Anonymize(&target); err != nil {
		fmt.Println("Fail to anonymize target data.")
	}

	fmt.Println(target)
}
```

# Replacers

## Default Replacers
Go Anonymizer provide following default replacers.

### asterisk
Replace your data with multiple asterisks `*`. The number of asterisk character is same with your string's length.

> "Hello World" -> "***********"

### empty
Replace your data with empty string.

> "Hello World" -> ""

## Custom Replacers
If you want to use your custom replacer, you can use `AddCustomReplacer()` and `RemoveCustomReplacer()`.