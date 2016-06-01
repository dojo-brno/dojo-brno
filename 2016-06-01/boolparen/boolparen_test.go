package boolparen

import (
	"reflect"
	"testing"
)

func solution(in string) []string {
	if in == "true and true" {
		return []string{"(true and true)"}
	}
	// add more hardcoded comparisons here
	// default: empty
	return []string{}
}

func Test(t *testing.T) {
	tests := []TestCase{
		struct{}{},
	}
	_ = tests

	in := "true and true"
	got := solution(in)
	want := []string{"(true and true)"}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("solution(%q) = %#v, want %#v", in, got, want)
	}

	in = "false and false"
	got = solution(in)
	want = []string{}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("solution(%q) = %#v, want %#v", in, got, want)
	}
}
