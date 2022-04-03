package comparator

import (
	"github.com/maurofran/hamcrest4go/constraints"
	"github.com/maurofran/hamcrest4go/matcher"
)

// LessThanOrEqualTo creates a new matcher for grater than comparison.
func LessThanOrEqualTo[T constraints.Ordered](value T) matcher.Matcher[T] {
	return lessThanOrEqualTo[T]{ref: value}
}

type lessThanOrEqualTo[T constraints.Ordered] struct {
	ref T
}

func (e lessThanOrEqualTo[T]) Matches(value T) bool {
	return value <= e.ref
}

func (e lessThanOrEqualTo[T]) DescribeTo(description matcher.Description) {
	description.AppendText("a value lesser than to ")
	description.AppendValue(e.ref)
}

func (e lessThanOrEqualTo[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendValue(actual)
	description.AppendText(" was lesser than ")
	description.AppendValue(e.ref)
}
