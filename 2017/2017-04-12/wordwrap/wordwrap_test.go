package wordwrap

import "testing"

func TestWordWrap(t *testing.T) {
	tests := []struct {
		text       string
		columnSize int
		want       string
	}{
		{
			"", 5, "",
		},
		{
			"word", 5, "word",
		},
		{
			"word word", 5, "word\nword",
		},
		{
			"word word word", 5, "word\nword\nword",
		},
		{
			"word word word", 100, "word word word",
		},
		{
			"word word word", 10, "word word\nword",
		},
		{
			"word word word", 9, "word word\nword",
		},
		{
			"Love Czec Republic", 10, "Love Czec\nRepublic",
		},
		{
			"Love Czech Republic", 10, "Love Czech\nRepublic",
		},
		{
			"Love Czech      Republic", 10, "Love Czech\nRepublic",
		},
		{
			"LoveCzechRepublic", 10, "LoveCzechR\nepublic",
		},
		{
			"LoveCzechRepublic", 3, "Lov\neCz\nech\nRep\nubl\nic",
		},
		{
			"Love          Republic", 10, "Love\nRepublic",
		},
		{
			"Love\nCzech", 7, "Love\nCzech",
		},
		{
			"Love\nCzech", 4, "Love\nCzec\nh",
		},
	}
	for _, tt := range tests {
		result := Wrap(tt.text, tt.columnSize)
		if result != tt.want {
			t.Errorf("Wrap(%q, %v) = %q, but wanted %q",
				tt.text, tt.columnSize, result, tt.want)
		}
	}

}
