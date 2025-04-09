package FedwireFundsPaymentStatus_002_001_10

import "reflect"

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
