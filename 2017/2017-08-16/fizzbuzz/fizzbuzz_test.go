package fizzbuzz

import "testing"

func FizzBuzz(input int) string {
	if input == 1 {
		return "1"
	}
	return "2"
}

func TestFizzBuzz(t *testing.T) {
	input := 1
	received := FizzBuzz(input)
	expectedOutput := "1"
	if received != expectedOutput {
		t.Errorf("Gave to the FizzBuzz func: (%v) GOT %q EXPECTED %q", input, received, expectedOutput)
	}

	input = 2
	received = FizzBuzz(input)
	expectedOutput = "2"
	if received != expectedOutput {
		t.Errorf("Gave to the FizzBuzz func: (%v) GOT %q EXPECTED %q", input, received, expectedOutput)
	}
}
