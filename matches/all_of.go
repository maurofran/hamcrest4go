package matches

import "github.com/maurofran/hamcrest4go/matcher"

// AllOf gets a matcher that assure that all the given matchers are matched.
func AllOf[T any](matchers ...matcher.Matcher[T]) matcher.Matcher[T] {
	return allOf[T]{matchers: matchers}
}

type allOf[T any] struct {
	matchers []matcher.Matcher[T]
}

func (n allOf[T]) Matches(value T) bool {
	for _, m := range n.matchers {
		if !m.Matches(value) {
			return false
		}
	}
	return true
}

func (n allOf[T]) DescribeTo(description matcher.Description) {
	var params []interface{}
	for _, m := range n.matchers {
		params = append(params, m)
	}
	description.AppendValueList("(", " and ", ")", params...)
}

func (n allOf[T]) DescribeMismatch(actual T, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
