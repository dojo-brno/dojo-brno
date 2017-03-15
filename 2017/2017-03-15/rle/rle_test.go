package rle

import (
	"reflect"
	"runtime"
	"testing"
)

var implementations = []func(string) string{
	Regexp,
}

func Test(t *testing.T) {
	in := "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW"
	want := "12W1B12W3B24W1B14W"
	for _, impl := range implementations {
		if got := impl(in); got != want {
			t.Errorf("%v(%q) = %q, want %q", nameOf(impl), in, got, want)
		}
	}
}

// nameOf returns the name of the function f. From
// https://www.reddit.com/r/golang/comments/11z36h/snippet_get_a_funcs_name_at_runtime/
// https://play.golang.org/p/Dyj99EjRVm
func nameOf(f interface{}) string {
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return v.String()
}
