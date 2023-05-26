package gotool

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Decimal struct {
	Value   interface{}
	Decimal decimal.Decimal
	Error   error
}

// InterfaceToDecimal - Interface轉Decimal
func (d *Decimal) InterfaceToDecimal() *Decimal {
	if d.Value == nil {
		d.Decimal = decimal.Zero
		d.Error = errors.New("Value can't nil")
		return d
	}
	switch v := d.Value.(type) {
	case float64:
		d.Decimal = decimal.NewFromFloat(d.Value.(float64))
	case float32:
		d.Decimal = decimal.NewFromFloat32(d.Value.(float32))
	case int:
		d.Decimal = decimal.NewFromInt(int64(d.Value.(int)))
	case int32:
		d.Decimal = decimal.NewFromInt32(d.Value.(int32))
	case int64:
		d.Decimal = decimal.NewFromInt(d.Value.(int64))
	case string:
		d.Decimal, d.Error = decimal.NewFromString(d.Value.(string))
	case []byte:
		d.Decimal, d.Error = decimal.NewFromString(string(v))
	default:
		d.Decimal = decimal.Zero
		d.Error = errors.New("not support class")
	}
	return d
}

// DecimalMulToInt64
func (d *Decimal) DecimalMulToInt64(value int64) (int64, error) {
	if d.Decimal.IsZero() || value == 0 {
		d.Error = errors.New("Value can't Zero")
		return 0, d.Error
	}

	return d.Decimal.Mul(decimal.NewFromInt(value)).IntPart(), nil
}

// DecimalInt64DivToString 回傳含指定小數位數的浮點數字串
// after the decimal point.
//
// Example:
//
//	fixed(2) // output: "0.00"
//	fixed(0) // output: "0"
//	fixed(0) // output: "5"
//	fixed(1) // output: "5.4"
//	fixed(2) // output: "5.45"
//	fixed(3) // output: "5.450"
//	fixed(-1) // output: "540"
func (d *Decimal) DecimalInt64DivToString(value int64, fixed int32) (string, error) {
	if d.Decimal.IsZero() || value == 0 {
		d.Error = errors.New("Value can't Zero")
		return "0", d.Error
	}

	return d.Decimal.Div(decimal.NewFromInt(value)).StringFixed(fixed), nil
}
