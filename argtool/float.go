package argtool

func RangeBetweenFloat(min, max float64) FloatValidator {
	return func(v float64, name string) error {
		if min > v || v > max {
			return &InvalidArgumentError{
				Name:   name,
				Reason: ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}
