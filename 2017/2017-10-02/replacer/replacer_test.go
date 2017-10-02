package replacer

import "testing"

func TestDictionaryReplacer(t *testing.T) {
	tests := []struct {
		input      string
		dictionary map[string]string
		want       string
	}{
		{"", map[string]string{}, ""},
		{"$temp$", map[string]string{"temp": "temporary"}, "temporary"},
		{"$temp$", map[string]string{"temp": "abc"}, "abc"},
		{"$temp$ $temp$", map[string]string{"temp": "temporary"}, "temporary temporary"},
		{"$temp$ $temp$", map[string]string{"temp": "$temp2$", "temp2": "temporary2"}, "$temp2$ $temp2$"},
	}
	for _, tt := range tests {

		input := tt.input
		output := DictionaryReplacer(input, tt.dictionary)
		if output != tt.want {
			t.Errorf("replacer.DictionaryReplacer failed for the input of %q, the expected output was %q, but function returned %q", tt.input, tt.want, output)

		}
	}
}
