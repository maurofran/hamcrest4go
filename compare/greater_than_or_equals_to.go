package compare

import (
	"github.com/maurofran/hamcrest4go/constraints"
	"github.com/maurofran/hamcrest4go/matcher"
)

// GreaterThanOrEqualsTo creates a new matcher for grater than comparison.
func GreaterThanOrEqualsTo[T constraints.Ordered](value T) matcher.Matcher[T] {
	return greaterThanOrEqualsTo[T]{ref: value}
}

type greaterThanOrEqualsTo[T constraints.Ordered] struct {
	ref T
}

func (e greaterThanOrEqualsTo[T]) Matches(value T) bool {
	return value >= e.ref
}

func (e greaterThanOrEqualsTo[T]) DescribeTo(description matcher.Description) {
	description.AppendText("a value greater than or equals to ")
	description.AppendValue(e.ref)
}

func (e greaterThanOrEqualsTo[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendValue(actual)
	description.AppendText(" was greater than or equals to ")
	description.AppendValue(e.ref)
}
