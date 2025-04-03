package ConnectionCheck_004_001_02

import "reflect"

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
