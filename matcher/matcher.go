package matcher

// Fn is the type of the matcher function.
type Fn[T any] func(value T) bool

type Matcher[T any] interface {
	SelfDescribing
	Matches(value T) bool
	DescribeMismatch(actual T, description Description)
}

type matcher[T any] struct {
	fn Fn[T]
}

func (m matcher[T]) Matches(value T) bool {
	return m.fn(value)
}

func (matcher[T]) DescribeMismatch(actual T, description Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}

// Custom creates a custom Matcher with supplied description and matcher function.
func Custom[T any](fixedDescription string, matchesFn Fn[T]) Matcher[T] {
	return custom[T]{
		matcher:          matcher[T]{fn: matchesFn},
		fixedDescription: fixedDescription,
	}
}

type custom[T any] struct {
	matcher[T]
	fixedDescription string
}

func (c custom[T]) DescribeTo(description Description) {
	description.AppendText(c.fixedDescription)
}
