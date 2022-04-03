package is

import "github.com/maurofran/hamcrest4go/matcher"

// Equaler is the constraint interface for an object that has the Equal method.
type Equaler[O any] interface {
	Equal(other O) bool
}

// EqualTo gets a matcher for nil value.
func EqualTo[T Equaler[T]](ref T) matcher.Matcher[T] {
	return equalTo[T]{ref: ref}
}

type equalTo[T Equaler[T]] struct {
	ref T
}

func (e equalTo[T]) Matches(value T) bool {
	return value.Equal(e.ref)
}

func (e equalTo[T]) DescribeTo(description matcher.Description) {
	description.AppendText("equal to ")
	description.AppendValue(e.ref)
}

func (equalTo[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
