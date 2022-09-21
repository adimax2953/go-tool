package argtool

import (
	"encoding/json"
	"math"
)

func Assert(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func NonEmptyString(v string, name string) error {
	if len(v) == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_EMPTY_STRING,
		}
	}
	return nil
}

func NonNanNorInf(v float64, name string) error {
	if isInfinity(v) || isNan(v) {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NAN_OR_INFINITY,
		}
	}
	return nil
}

func NonNegativeInteger(v int64, name string) error {
	if v < 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_NEGATIVE_INTENGER,
		}
	}
	return nil
}

func NonNegativeNumber(v float64, name string) error {
	if math.Signbit(v) {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_NEGATIVE_NUMBER,
		}
	}
	return nil
}

func JsonInteger(v json.Number, name string, validators ...IntegerValidator) error {
	integer, err := v.Int64()
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: err.Error(),
		}
	}
	for _, validate := range validators {
		if err := validate(integer, name); err != nil {
			return err
		}
	}
	return nil
}

func JsonNumber(v json.Number, name string, validators ...FloatValidator) error {
	float, err := v.Float64()
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: err.Error(),
		}
	}
	/* NOTE: normalize the float64. avoid the "-0" be treated as Signbit() carried value
	 * e.g:
	 *   var s json.Number = "-0"
	 *   f, _ := s.Float64()
	 *   math.Signbit(f)   // return true
	 * but:
	 *   var f float64 = -0
	 *   math.Signbit(f)   // return false
	 */
	if float == 0 {
		float = 0
	}
	for _, validate := range validators {
		if err := validate(float, name); err != nil {
			return err
		}
	}
	return nil
}

func ThrowError(name, reason string) error {
	return &InvalidArgumentError{
		Name:   name,
		Reason: reason,
	}
}
