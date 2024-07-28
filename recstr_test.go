package recstr

import (
	"fmt"
	"github.com/alecthomas/repr"
	"testing"
)

type Struct struct {
	Int          int
	Str          string
	Float        float64
	IntSlice     []int
	FstStruct    FstStruct
	FstStructPtr *FstStruct
}

type FstStruct struct {
	IntPtr    *int
	StrPtr    *string
	FloatStr  *float64
	Int2Str   map[int]string
	SecStruct SecStruct
}

type SecStruct struct {
	Any    any
	hidden any
}

var (
	fstStruct = FstStruct{
		IntPtr:   P(4),
		StrPtr:   P("FstStrPtr"),
		FloatStr: P(3.14),
		Int2Str:  map[int]string{1: "a", 2: "b", 3: "c"},
		SecStruct: SecStruct{
			Any:    "Any",
			hidden: "hidden",
		},
	}
	s = Struct{
		Int:          1,
		Str:          "Str",
		IntSlice:     []int{1, 2, 3},
		FstStruct:    fstStruct,
		FstStructPtr: &fstStruct,
	}
)

func TestCompare2Repr(t *testing.T) {
	fmt.Println(repr.String(s))
	fmt.Println(Of(s, RecursionLimit(3)))
}

func P[T any](v T) *T {
	return &v
}
