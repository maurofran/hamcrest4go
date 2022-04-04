package compare

import (
	"github.com/maurofran/hamcrest4go/constraints"
	"github.com/maurofran/hamcrest4go/matcher"
)

// LessThanOrEqualsTo creates a new matcher for grater than comparison.
func LessThanOrEqualsTo[T constraints.Ordered](value T) matcher.Matcher[T] {
	return lessThanOrEqualsTo[T]{ref: value}
}

type lessThanOrEqualsTo[T constraints.Ordered] struct {
	ref T
}

func (e lessThanOrEqualsTo[T]) Matches(value T) bool {
	return value <= e.ref
}

func (e lessThanOrEqualsTo[T]) DescribeTo(description matcher.Description) {
	description.AppendText("a value less than or equals to ")
	description.AppendValue(e.ref)
}

func (e lessThanOrEqualsTo[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendValue(actual)
	description.AppendText(" was less than or equals to ")
	description.AppendValue(e.ref)
}
