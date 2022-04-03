package is

import "github.com/maurofran/hamcrest4go/matcher"

// Not implements the negation of a given mathcer.
func Not[T any](original matcher.Matcher[T]) matcher.Matcher[T] {
	return not[T]{original: original}
}

type not[T any] struct {
	original matcher.Matcher[T]
}

func (n not[T]) Matches(value T) bool {
	return !n.original.Matches(value)
}

func (n not[T]) DescribeTo(description matcher.Description) {
	description.AppendText("not ")
	description.AppendDescriptionOf(n.original)
}

func (n not[T]) DescribeMismatch(actual T, description matcher.Description) {
	n.original.DescribeMismatch(actual, description)
}
