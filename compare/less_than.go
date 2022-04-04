package compare

import (
	"github.com/maurofran/hamcrest4go/constraints"
	"github.com/maurofran/hamcrest4go/matcher"
)

// LessThan creates a new matcher for grater than comparison.
func LessThan[T constraints.Ordered](value T) matcher.Matcher[T] {
	return lessThan[T]{ref: value}
}

type lessThan[T constraints.Ordered] struct {
	ref T
}

func (e lessThan[T]) Matches(value T) bool {
	return value < e.ref
}

func (e lessThan[T]) DescribeTo(description matcher.Description) {
	description.AppendText("a value lesser than ")
	description.AppendValue(e.ref)
}

func (e lessThan[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendValue(actual)
	description.AppendText(" was lesser than ")
	description.AppendValue(e.ref)
}
