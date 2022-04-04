package compare

import (
	"github.com/maurofran/hamcrest4go/matcher"
)

// EqualsTo creates a new matcher for equality comparison.
func EqualsTo[T comparable](value T) matcher.Matcher[T] {
	return equalsTo[T]{ref: value}
}

type equalsTo[T comparable] struct {
	ref T
}

func (e equalsTo[T]) Matches(value T) bool {
	return e.ref == value
}

func (e equalsTo[T]) DescribeTo(description matcher.Description) {
	description.AppendText("a value equal to ")
	description.AppendValue(e.ref)
}

func (e equalsTo[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendValue(actual)
	description.AppendText(" was equal to ")
	description.AppendValue(e.ref)
}
