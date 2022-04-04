package matcher

// Fn is the type of the matcher function.
type Fn[T any] func(value T) bool

// Custom creates a custom Matcher with supplied description and matcher function.
func Custom[T any](fixedDescription string, matchesFn Fn[T]) Matcher[T] {
	return custom[T]{
		fn:               matchesFn,
		fixedDescription: fixedDescription,
	}
}

type custom[T any] struct {
	fn               Fn[T]
	fixedDescription string
}

func (c custom[T]) Matches(value T) bool {
	return c.fn(value)
}

func (custom[T]) DescribeMismatch(actual T, description Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}

func (c custom[T]) DescribeTo(description Description) {
	description.AppendText(c.fixedDescription)
}
