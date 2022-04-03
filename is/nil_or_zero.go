package is

import "github.com/maurofran/hamcrest4go/matcher"

// NotNilOrZero gets a matcher for not nilOrZero.
func NotNilOrZero[T comparable]() matcher.Matcher[T] {
	return Not[T](NilOrZero[T]())
}

// NilOrZero gets a matcher for nil value.
func NilOrZero[T comparable]() matcher.Matcher[T] {
	return nilOrZero[T]{}
}

type nilOrZero[T comparable] struct {
}

func (nilOrZero[T]) Matches(value T) bool {
	var z T
	return value == z
}

func (nilOrZero[T]) DescribeTo(description matcher.Description) {
	description.AppendText("nilOrZero")
}

func (nilOrZero[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
