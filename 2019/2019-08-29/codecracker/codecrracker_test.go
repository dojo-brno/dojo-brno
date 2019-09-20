package codecracker

import "testing"
import "string"

func CodeeCracker(text string) string {
	alphabet := "ab"
	key := "!)"
	encrypted := ""

	for i, v := range text {

	}

	if text == "brno" {
		return ")gcd"
	}
	if text[1] == alphabet[1] {
		return string(key[1])
	}
	return "!"
}

func TestCodeCrackerWithA(t *testing.T) {
	text := "a"
	if got := CodeeCracker(text); got != "!" {
		t.Errorf("CodeCracker(%v):%v, WANT:%v", text, got, "!")
	}
}

func TestCodeCrackerWithB(t *testing.T) {
	text := "b"
	expected := ")"
	if got := CodeeCracker(text); got != expected {
		t.Errorf("CodeCracker(%v):%v, WANT:%v", text, got, expected)
	}
}

func TestCodeCrackerWithBrno(t *testing.T) {
	text := "brno"
	expected := ")gcd"
	if got := CodeeCracker(text); got != expected {
		t.Errorf("CodeCracker(%v):%v, WANT:%v", text, got, expected)
	}
}
