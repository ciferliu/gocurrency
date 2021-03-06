package currency

import "testing"

func init() {
	Factory.NewCurrency("USD", 2)
	Factory.NewCurrency("CNY", 2)
}

func TestNewCurrency(t *testing.T) {
	//case 1:
	currencyCode := ""
	_, err := Factory.NewCurrency(currencyCode, 2)
	if err == nil {
		t.Errorf("Factory.NewCurrency(%s, 2), shoule be return an error, but no error return", currencyCode)
	}

	//case 2:
	currencyCode = "123"
	_, err = Factory.NewCurrency(currencyCode, 2)
	if err == nil {
		t.Errorf("Factory.NewCurrency(%s, 2), shoule be return an error, but no error return", currencyCode)
	}

	//case 3:
	currencyCode = "ab"
	_, err = Factory.NewCurrency(currencyCode, 2)
	if err == nil {
		t.Errorf("Factory.NewCurrency(%s, 2), shoule be return an error, but no error return", currencyCode)
	}

	//case 3:
	currencyCode = "usd"
	got, err := Factory.NewCurrency(currencyCode, 2)
	want := Currency{"USD", 2}
	if want != got {
		t.Errorf("Factory.NewCurrency(%s, 2) == %v, want %v", currencyCode, got, want)
	}
}

func TestNewAmountInBasicUnit(t *testing.T) {
	//case 1:
	currencyCode := ""
	basicUnitValue := "1"
	_, err := Factory.NewAmountInBasicUnit(currencyCode, basicUnitValue)
	if err == nil {
		t.Errorf("Factory.NewAmountInBasicUnit(%s, %s), shoule be return an error, but no error return", currencyCode, basicUnitValue)
	}

	//case 2:
	currencyCode = "ABC"
	basicUnitValue = "1"
	_, err = Factory.NewAmountInBasicUnit(currencyCode, basicUnitValue)
	if err == nil {
		t.Errorf("Factory.NewAmountInBasicUnit(%s, %s), shoule be return an error, but no error return", currencyCode, basicUnitValue)
	}

	//case 3:
	currencyCode = "usd"
	basicUnitValue = "-1"
	amount, _ := Factory.NewAmountInBasicUnit(currencyCode, basicUnitValue)
	got := amount.String()
	want := "USD -1.00"
	if got != want {
		t.Errorf("Factory.NewAmountInBasicUnit(%s, %s) == %s, want %s", currencyCode, basicUnitValue, got, want)
	}
}

func TestNewAmountInMinorUnit(t *testing.T) {
	//case 1:
	currencyCode := ""
	minorUnitValue := int64(100)
	_, err := Factory.NewAmountInMinorUnit(currencyCode, minorUnitValue)
	if err == nil {
		t.Errorf("Factory.NewAmountInMinorUnit(%s, %d), shoule be return an error, but no error return", currencyCode, minorUnitValue)
	}

	//case 2:
	currencyCode = "ABC"
	_, err = Factory.NewAmountInMinorUnit(currencyCode, minorUnitValue)
	if err == nil {
		t.Errorf("Factory.NewAmountInMinorUnit(%s, %d), shoule be return an error, but no error return", currencyCode, minorUnitValue)
	}

	//case 3:
	currencyCode = "usd"
	minorUnitValue = int64(-1)
	amount, _ := Factory.NewAmountInMinorUnit(currencyCode, minorUnitValue)
	got := amount.String()
	want := "USD -0.01"
	if got != want {
		t.Errorf("Factory.NewAmountInMinorUnit(%s, %d) == %s, want %s", currencyCode, minorUnitValue, got, want)
	}
}
