package argtool

// Reason
const (
	ERR_OUT_OF_RANGE                = "out of range"
	ERR_NON_NEGATIVE_INTENGER       = "should be a non-negative integer"
	ERR_NON_NEGATIVE_INTENGER_SLICE = "should be a non-negative integer slice"
	ERR_NEGATIVE_INTENGER           = "should be a negative integer"
	ERR_NON_NEGATIVE_NUMBER         = "should be a non-negative number"
	ERR_NON_NEGATIVE_NUMBER_SLICE   = "should be a non-negative number slice"
	ERR_NAN_OR_INFINITY             = "cannot be -inf, +inf or NaN"
	ERR_EMPTY_STRING                = "cannot be an empty string"
	ERR_INVALID_INTEGER_ASSERTION   = "specified integer %d is invalid"
)

type (
	IntegerValidator func(v int64, name string) error
	FloatValidator   func(v float64, name string) error
)

type NonNegative_Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Negative_Integer interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type NonNegative_Number interface {
	~float32 | ~float64 | ~complex64 | ~complex128
}

type Slice_NonNegative_Integer interface {
	~[]int64 | ~[]int32 | ~[]int16 | ~[]int8 | ~[]int
}
type Slice_Negative_Integer interface {
	~[]uint64 | ~[]uint32 | ~[]uint16 | ~[]uint8 | ~[]uint
}
type Slice_NonNegative_Number interface {
	~[]float64 | ~[]float32 | ~[]complex64 | ~[]complex128
}
