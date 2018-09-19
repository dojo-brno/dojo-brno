package stringcalculator

import (
	"strconv"
	"strings"
	"testing"
)

func StringCalculator(number string) string {
	if number == "" {
		return "0"
	}
	numbers := strings.Split(number, ",")
	s := 0
	for _, n := range numbers {
		v, _ := strconv.Atoi(n)
		s = s + v
	}
	return strconv.Itoa(s)
}

func TestStringCalculator(t *testing.T) {
	// first test
	input := ""
	want := "0"
	if want != StringCalculator(input) {
		t.Errorf("Some text")
	}

	// second test
	input = "1"
	want = "1"
	if got := StringCalculator(input); want != got {
		t.Errorf("Input as %v  was misscounted as %v but we are expecting %v", input, got, want)
	}

	// third  test
	input = "1,1"
	want = "2"
	if got := StringCalculator(input); want != got {
		t.Errorf("Input as %v  was misscounted as %v but we are expecting %v", input, got, want)
	}

	// fourth  test
	input = "2,2"
	want = "4"
	if got := StringCalculator(input); want != got {
		t.Errorf("Input as %v  was misscounted as %v but we are expecting %v", input, got, want)
	}

  /* For Jan Houska as QA but for Radomir Ludva is this test case green and do not force as to do anythink with production code
	// random numbersw
	input = "1,2,3,4,5"
	want = "15"
	if got := StringCalculator(input); want != got {
		t.Errorf("Input as %v  was misscounted as %v but we are expecting %v", input, got, want)
	}*/
}
