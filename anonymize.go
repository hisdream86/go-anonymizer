package goanonymizer

import (
	"errors"
	"reflect"
)

func replace(tag string, v reflect.Value) {
	var replacer Replacer
	if br, ok := builtin[tag]; ok {
		replacer = br
	} else if cr, ok := custom[tag]; ok {
		replacer = cr
	}
	if replacer != nil {
		replaced := replacer(v.String())
		v.SetString(replaced)
	}
}

func anonymize(tag string, v reflect.Value) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		replace(tag, v)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			anonymize(tag, v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.Type().NumField(); i++ {
			tag := v.Type().Field(i).Tag.Get("anonymize")
			anonymize(tag, v.Field(i))
		}
	}
}

// Anonymize anonymize all tagged field of the target data.
func Anonymize(target interface{}) error {
	v := reflect.ValueOf(target)

	for v.Kind() != reflect.Ptr {
		return errors.New("target is not pointer")
	}

	anonymize("", reflect.ValueOf(target))

	return nil
}
