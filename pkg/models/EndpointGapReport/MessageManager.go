package EndpointGapReport

import "reflect"

type GapType string

const (
	InputMessageAccountabilityData  GapType = "IMAD"
	OutputMessageAccountabilityData GapType = "OMAD"
)

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
