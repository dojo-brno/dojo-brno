package wallet

import "testing"

func rateProvider(fromCurrency string, toCurrency string) float64 {
	return 25
}

func InternetRateProvider(fromCurrency string, toCurrency string) string {
	// token := ""
	// url :=
	//=r.get("http://data.fixer.io/api/latest?access_key=<access_token>")

	return "..."
}

/////////////////////////////////

func TestInternetRateProvider(t *testing.T) {
	fromCurrency := "EUR"
	toCurrency := "CZK"
	var notWantedCoefficient string
	notWantedCoefficient = ""

	if coefficient := InternetRateProvider(fromCurrency, toCurrency); coefficient == notWantedCoefficient {
		t.Fail()
	}
}

func TestRateProvider(t *testing.T) {
	fromCurrency := "EUR"
	toCurrency := "CZK"
	wantedCoefficient := float64(25)

	if coefficient := rateProvider(fromCurrency, toCurrency); wantedCoefficient != coefficient {
		t.Errorf("rateProvider(%q, %q):%v, but we WANT %v", fromCurrency, toCurrency, coefficient, wantedCoefficient)
	}
}

///

type Wallet struct {
	stock Stock
}
type Stock struct {
	name   string
	amount int
}

func (w Wallet) Value() float64 {
	return float64(w.stock.amount)
}

func TestValue(t *testing.T) {
	tests := []struct {
		wallet      Wallet
		wantedValue float64
	}{
		{Wallet{}, 0},
		{Wallet{Stock{"peroleum", 100}}, 100},
	}

	for _, tt := range tests {
		if got := tt.wallet.Value(); got != tt.wantedValue {
			t.Errorf("`%v`.Value():%v, WANT %v", tt.wallet, got, tt.wantedValue)
		}

	}
}
