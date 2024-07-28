package recstr

import (
	"bytes"
	"fmt"
	"reflect"
)

type option func(c *config)

type config struct {
	recursionLimit int
}

var (
	c = config{}
)

func Of(v any, options ...option) string {
	c.recursionLimit = 2
	for _, option := range options {
		option(&c)
	}
	w := bytes.NewBuffer(nil)
	parse(reflect.ValueOf(v), w, 0)
	return w.String()
}

func RecursionLimit(n int) option {
	return func(c *config) {
		c.recursionLimit = n
	}
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
		if depth >= c.recursionLimit {
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
	default:
		// TODO: parse ptr in Slice/Array/Map
		w.WriteString(fmt.Sprintf("%#v", v))
	}
}
