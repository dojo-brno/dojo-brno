package birthnumber

import "testing"

func TestIsBirthNumberValid(t *testing.T) {
	tests := []struct {
		testReason  string
		birthNumber string
		want        bool
	}{
		{"Initial valid 10 digit", "950407/7715", true},
		{"Initial invalid 10 digit", "921013/0213", false},
		{"Force implementation %11", "991211/0449", false},
		{"Force implementation The Rest of Division == 10", "780123/3540", true},
		{"Pre 1954 9 digit => no validation", "530407/771", true},
		{"Pre 1954 9 digit with invalid month", "531407/771", false},
		/*{"905704/1411", true},
		{"921015/0213", true},
		{"950607/5623", true},
		{"990610/8905", true},
		{"951215/6687", true},
		{"991212/0449", true},
		{"991118/9629", true},
		{"950607/5625", false},
		{"951607/5625", false},
		{"990611/8905", false},
		{"990510/8905", false},
		{"950406/7715", false},*/
	}

	for _, tt := range tests {
		if got := IsBirthNumberValid(tt.birthNumber); got != tt.want {
			t.Errorf("IsBirthNumberValid(%q) = %v WANT %v", tt.birthNumber, got, tt.want)
		}
	}
}
