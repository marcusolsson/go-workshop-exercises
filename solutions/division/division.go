package division

import "errors"

// Divide returns the quotient of dividing the dividend with the divisor,
// or returns an error when divisor is zero.
func Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0.0, errors.New("division by zero")
	}
	return dividend / divisor, nil
}
