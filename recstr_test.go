package recstr

import (
	"fmt"
	"github.com/alo1719/mygo"
	"testing"
	"time"
)

type Struct struct {
	Int             int
	Str             string
	Float           float64
	IntSlice        []int
	FstStruct       FstStruct
	FstStructPtr    *FstStruct
	FstStructSlc    []FstStruct
	SecStructPtrSlc []*SecStruct
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
	secStruct = SecStruct{
		Any:    "Any",
		hidden: "hidden",
	}
	fstStruct = FstStruct{
		IntPtr:    mygo.P(4),
		StrPtr:    mygo.P("FstStrPtr"),
		FloatStr:  mygo.P(3.14),
		Int2Str:   map[int]string{1: "a", 2: "b", 3: "c"},
		SecStruct: secStruct,
	}
	s = Struct{
		Int:             1,
		Str:             "Str",
		IntSlice:        []int{1, 2, 3},
		FstStruct:       fstStruct,
		FstStructPtr:    &fstStruct,
		FstStructSlc:    []FstStruct{fstStruct, fstStruct},
		SecStructPtrSlc: []*SecStruct{&secStruct, &secStruct},
	}
)

func Test(t *testing.T) {
	fmt.Println(Of(s))
}

func TestPerformance(t *testing.T) {
	tt := time.Now()
	for i := 0; i < 1000; i++ {
		Of(s)
	}
	fmt.Println(time.Since(tt))
}
