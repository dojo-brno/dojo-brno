package numbers_test

import (
	"testing"

	"github.com/dojo-brno/dojo-brno/2016/2016-06-08/numbers"
)



const lcd7 = `
 _
  |
  |
`
const lcd8 = `
 _
  |
  |
`
func TestLCD(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 7,
			want: lcd7,
		},
		{
			input: 8,
			want: `
 _
|_|
|_|
`,
		},
	}
	for _, test := range tests {
		if got := numbers.ToLCD(test.input); got != test.want {
			t.Errorf("ToLCD(%d) = %s, want %s", test.input, got, test.want)
		}
	}
}
