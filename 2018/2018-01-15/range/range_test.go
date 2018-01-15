package intrvl

import (
	"reflect"
	"testing"
)

// [2,6) allPoints = {2,3,4,5}

// assert(zacatek < konec)
type Interval struct {
	zacatek         int
	konec           int
	zacatekOtevreny bool
	konecOtevreny   bool
}

func AllPoints(i Interval) []int {
	vysledek := []int{}
	if !i.zacatekOtevreny {
		vysledek = append(vysledek, i.zacatek)
	}
	for j := i.zacatek + 1; j < i.konec; j++ {
		vysledek = append(vysledek, j)
	}
	if !i.konecOtevreny && i.zacatek != i.konec {
		vysledek = append(vysledek, i.konec)
	}
	return vysledek
}

func TestAllPoints(t *testing.T) {
	testy := []struct {
		inter  Interval
		chceme []int
	}{
		{
			inter: Interval{
				zacatekOtevreny: false,
				zacatek:         2,
				konec:           6,
				konecOtevreny:   true,
			},
			chceme: []int{2, 3, 4, 5},
		},
		{
			inter: Interval{
				zacatekOtevreny: true,
				zacatek:         2,
				konec:           6,
				konecOtevreny:   true,
			},
			chceme: []int{3, 4, 5},
		},
		{
			inter: Interval{
				zacatekOtevreny: true,
				zacatek:         2,
				konec:           6,
				konecOtevreny:   false,
			},
			chceme: []int{3, 4, 5, 6},
		},

		{
			inter: Interval{
				zacatekOtevreny: true,
				zacatek:         2,
				konec:           10,
				konecOtevreny:   true,
			},
			chceme: []int{3, 4, 5, 6, 7, 8, 9},
		},
		{
			inter: Interval{
				zacatekOtevreny: false,
				zacatek:         2,
				konec:           10,
				konecOtevreny:   false,
			},
			chceme: []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			inter: Interval{
				zacatekOtevreny: false,
				zacatek:         2,
				konec:           10,
				konecOtevreny:   true,
			},
			chceme: []int{2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			inter: Interval{
				zacatekOtevreny: false,
				zacatek:         -2,
				konec:           10,
				konecOtevreny:   true,
			},
			chceme: []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			inter: Interval{
				zacatekOtevreny: true,
				zacatek:         2,
				konec:           3,
				konecOtevreny:   true,
			},
			chceme: []int{},
		},
		{
			inter: Interval{
				zacatekOtevreny: false,
				zacatek:         2,
				konec:           2,
				konecOtevreny:   false,
			},
			chceme: []int{2},
		},
	}
	for _, tt := range testy {
		if vysledek := AllPoints(tt.inter); !reflect.DeepEqual(vysledek, tt.chceme) {
			t.Errorf("AllPoints(%v) = %v chceme %v", tt.inter, vysledek, tt.chceme)
		}
	}
}

func CreateInterval(int, int, bool, bool) (Interval, error) {
	return Interval{/*zacatek: 0, konec: 0,*/ zacatekOtevreny: false, konecOtevreny: true}, nil
}

func TestCreateInterval(t *testing.T) {
	testy := []struct {
		zacatek, konec       int
		zOtevreny, kOtevreny bool
		chceme               Interval
		chyba                error
	}{
		{
			0, 0, true, true, Interval{zacatekOtevreny: true, konecOtevreny: true}, nil,
		},
	}
	for _, tt := range testy {
		if vysledek, chyba := CreateInterval(tt.zacatek, tt.konec, tt.zOtevreny, tt.kOtevreny); tt.chyba != chyba {
			t.Errorf("CreateInterval(%v, %v, %v, %v) = (%v, %v) chceme (%v, %v)", tt.zacatek, tt.konec, tt.zOtevreny, tt.kOtevreny, vysledek, chyba, tt.chceme, tt.chyba)
		} else {
			if chyba == tt.chyba {

			}
		}
	}
}
