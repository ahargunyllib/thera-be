package decimal

import "github.com/shopspring/decimal"

type CustomDecimalInterface interface {
	NewFromString(string) (decimal.Decimal, error)
	NewFromFloat64(float64) decimal.Decimal
}

type CustomDecimalStruct struct{}

var Decimal = getDecimal()

func getDecimal() CustomDecimalInterface {
	return &CustomDecimalStruct{}
}

func (d *CustomDecimalStruct) NewFromString(str string) (decimal.Decimal, error) {
	dec, err := decimal.NewFromString(str)

	if err != nil {
		return dec, err
	}

	return dec, nil
}

func (d *CustomDecimalStruct) NewFromFloat64(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f)
}
