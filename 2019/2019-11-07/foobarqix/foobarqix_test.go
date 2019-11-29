package foobarqix

import (
	"strconv"
	"testing"
)

func addFBQ(n int, prime int, addFBQ string) string {
	if n%prime == 0 {
		return addFBQ
	}
	return ""
}

func GetFBX(n int) string {
	nstring := strconv.Itoa(n)
	snumber := ""

	snumber += addFBQ(n, 3, "Foo")
	snumber += addFBQ(n, 5, "Bar")
	snumber += addFBQ(n, 7, "Qix")

	for _, nchar := range nstring {
		if nchar == '3' {
			snumber += "Foo"
		}
		if nchar == '5' {
			snumber += "Bar"
		}
		if nchar == '7' {
			snumber += "Qix"
		}
	}

	if snumber == "" {
		snumber = nstring
	}

	return snumber
}

func TestGetFBX(t *testing.T) {
	type testCase = struct {
		n   int
		fbx string
	}

	testCases := []testCase{
		{1, "1"},
		{2, "2"},
		{3, "FooFoo"},
		{4, "4"},
		{5, "BarBar"},
		{6, "Foo"},
		{7, "QixQix"},
		{8, "8"},
		{9, "Foo"},
		{10, "Bar"},
		{11, "11"},
		{12, "Foo"},
		{13, "Foo"},
		{14, "Qix"},
		{15, "FooBarBar"},
		{21, "FooQix"},
		{33, "FooFooFoo"},
		{30, "FooBarFoo"},
		{51, "FooBar"},
		{53, "BarFoo"},
	}

	for _, tt := range testCases {
		if got := GetFBX(tt.n); got != tt.fbx {
			t.Errorf("GetFBX(%v):%v, WANT:%v", tt.n, got, tt.fbx)
		}
	}

}
