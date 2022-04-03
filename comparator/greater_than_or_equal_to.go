package comparator

import (
	"github.com/maurofran/hamcrest4go/constraints"
	"github.com/maurofran/hamcrest4go/matcher"
)

// GreaterThanOrEqualTo creates a new matcher for grater than comparison.
func GreaterThanOrEqualTo[T constraints.Ordered](value T) matcher.Matcher[T] {
	return greaterThanOrEqualTo[T]{ref: value}
}

type greaterThanOrEqualTo[T constraints.Ordered] struct {
	ref T
}

func (e greaterThanOrEqualTo[T]) Matches(value T) bool {
	return value >= e.ref
}

func (e greaterThanOrEqualTo[T]) DescribeTo(description matcher.Description) {
	description.AppendText("a value greater than or equal to ")
	description.AppendValue(e.ref)
}

func (e greaterThanOrEqualTo[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendValue(actual)
	description.AppendText(" was greater than or equal to ")
	description.AppendValue(e.ref)
}
