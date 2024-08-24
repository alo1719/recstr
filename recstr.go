package recstr

import (
	"bytes"
	"fmt"
	"reflect"
)

type option func(c *config)

type config struct {
	recursionLimit int
	lengthLimit    int
}

// Enough for most cases
var (
	c = config{
		recursionLimit: 3,
		lengthLimit:    25,
	}
)

func Of(v any, options ...option) string {
	oc := c
	for _, option := range options {
		option(&c)
	}
	w := bytes.NewBuffer(nil)
	parse(reflect.ValueOf(v), w, 0)
	c = oc
	return w.String()
}

func RecursionLimit(n int) option {
	return func(c *config) {
		c.recursionLimit = n
	}
}

func LengthLimit(n int) option {
	return func(c *config) {
		c.lengthLimit = n
	}
}

func SetGlobalRecursionLimit(n int) {
	c.recursionLimit = n
}

func SetGlobalLengthLimit(n int) {
	c.lengthLimit = n
}

func parse(v reflect.Value, w *bytes.Buffer, depth int) {
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			w.WriteString("nil")
			return
		}
		w.WriteString("&")
		parse(v.Elem(), w, depth)
	case reflect.Struct:
		w.WriteString(v.Type().Name())
		w.WriteString("{")
		if depth > c.recursionLimit {
			w.WriteString("...") // deep dark recursion
		} else {
			for i := 0; i < v.NumField(); i++ {
				if i > 0 {
					w.WriteString(", ")
				}
				w.WriteString(v.Type().Field(i).Name)
				w.WriteString(": ")
				parse(v.Field(i), w, depth+1)
			}
		}
		w.WriteString("}")
	case reflect.Slice, reflect.Array:
		w.WriteString("[]" + v.Type().Elem().String() + "{")
		if depth > c.recursionLimit {
			w.WriteString("...")
		} else {
			for i := 0; i < v.Len(); i++ {
				if i > 0 {
					w.WriteString(", ")
				}
				if i == c.lengthLimit {
					w.WriteString("...")
					break
				}
				parse(v.Index(i), w, depth)
			}
		}
		w.WriteString("}")
	default:
		// TODO: parse ptr in Map
		w.WriteString(fmt.Sprintf("%#v", v))
	}
}
