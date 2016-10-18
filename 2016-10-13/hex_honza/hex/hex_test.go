package hex

import "testing"

var testCases = []struct {
	input  string
	output string
}{
	{"H", "48"},
	{"Hello, world!", "48656c6c6f2c20776f726c6421"},
}

func TestByteToHex(t *testing.T) {
	expectedResult := "7b"
	result := ByteToHexByte(byte('{'))
	if result != expectedResult {
		t.Errorf("Expected: %v, got: %v", expectedResult, result)
	}
}

func TestBytesToText(t *testing.T) {
	for _, tc := range testCases {
		if result := BytesToText(tc.input); result != tc.output {
			t.Errorf("Expected: %v, got: %v", tc.output, result)
		}
	}
}
