package goanonymizer

import (
	"errors"
	"strings"
)

type Replacer func(string) string

var builtin map[string]Replacer
var custom map[string]Replacer

func init() {
	builtin = make(map[string]Replacer)
	custom = make(map[string]Replacer)

	builtin["empty"] = emptystringReplacer
	builtin["asterisk"] = asteriskReplacer
}

// AddCustomReplacer add a custom replacer with received name.
// It returns error when the name is empty or replacer is null.
func AddCustomReplacer(name string, replacer Replacer) error {
	if len(name) == 0 {
		return errors.New("replacer name is null")
	}

	if replacer == nil {
		return errors.New("replacer is nil")
	}

	custom[name] = replacer

	return nil
}

// RemoveCustomReplacer removes a custom replacer which has received name.
// It returns error when the name is empty or no replacer is exists.
func RemoveCustomReplacer(name string) error {
	if len(name) == 0 {
		return errors.New("replacer name is null")
	}

	if _, ok := custom[name]; !ok {
		return errors.New("replacer is not exists")
	}

	return nil
}

func emptystringReplacer(source string) string {
	return ""
}

func asteriskReplacer(source string) string {
	v := []rune(source)
	masked := make([]string, len(v))
	for idx := range masked {
		masked[idx] = "*"
	}
	return strings.Join(masked, "")
}
