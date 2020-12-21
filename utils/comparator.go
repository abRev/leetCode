package utils

// Comparator will make type assertion (see IntComparator for example),
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type Comparator func(a, b interface{}) int

// IntComparator provides a basic comparison on int
func IntComparator(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
