package comparator

import (
	"github.com/maurofran/hamcrest4go/constraints"
	"github.com/maurofran/hamcrest4go/matcher"
)

// GreaterThan creates a new matcher for grater than comparison.
func GreaterThan[T constraints.Ordered](value T) matcher.Matcher[T] {
	return greaterThan[T]{ref: value}
}

type greaterThan[T constraints.Ordered] struct {
	ref T
}

func (e greaterThan[T]) Matches(value T) bool {
	return value > e.ref
}

func (e greaterThan[T]) DescribeTo(description matcher.Description) {
	description.AppendText("a value greater than ")
	description.AppendValue(e.ref)
}

func (e greaterThan[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendValue(actual)
	description.AppendText(" was greater than ")
	description.AppendValue(e.ref)
}
