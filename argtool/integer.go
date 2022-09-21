package argtool

import (
	"fmt"
	"sort"
)

func IntegerNotIn(values ...int64) IntegerValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v int64, name string) error {
		i := sort.Search(len(values), func(i int) bool { return values[i] >= v })
		if i < len(values) && values[i] == v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(ERR_INVALID_INTEGER_ASSERTION, v),
			}
		}
		return nil
	}
}

func RangeBetween(min, max int64) IntegerValidator {
	return func(v int64, name string) error {
		if min > v || v > max {
			return &InvalidArgumentError{
				Name:   name,
				Reason: ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}
