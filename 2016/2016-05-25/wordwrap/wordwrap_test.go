package wordwrap

import "testing"

func TestWordwrap(t *testing.T) {
	result := Wrap("", 0)
	expected := ""
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

	result = Wrap("foo", 10)
	expected = "foo"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

	result = Wrap("foo bar", 4)
	expected = "foo\nbar"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

	result = Wrap("foo bar", 20)
	expected = "foo bar"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

	result = Wrap("foo bar", 7)
	expected = "foo bar"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

	result = Wrap("foo bar", 6)
	expected = "foo\nbar"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

	result = Wrap("bar bar", 6)
	expected = "bar\nbar"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

}
