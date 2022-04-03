package matches

import "github.com/maurofran/hamcrest4go/matcher"

// AnyOf gets a matcher that assure that at least one of the given matchers is matched.
func AnyOf[T any](matchers ...matcher.Matcher[T]) matcher.Matcher[T] {
	return anyOf[T]{matchers: matchers}
}

type anyOf[T any] struct {
	matchers []matcher.Matcher[T]
}

func (n anyOf[T]) Matches(value T) bool {
	for _, m := range n.matchers {
		if m.Matches(value) {
			return true
		}
	}
	return false
}

func (n anyOf[T]) DescribeTo(description matcher.Description) {
	var params []interface{}
	for _, m := range n.matchers {
		params = append(params, m)
	}
	description.AppendValueList("(", " or ", ")", params...)
}

func (n anyOf[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
