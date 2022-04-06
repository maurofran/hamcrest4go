package described

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"regexp"
	"strconv"
)

var argPattern = regexp.MustCompile("%([0-9]+)")

// As gets a description decorator matcher.
func As[T any](template string, matcher matcher.Matcher[T], args ...interface{}) matcher.Matcher[T] {
	return as[T]{
		matcher:  matcher,
		template: template,
		args:     args,
	}
}

type as[T any] struct {
	matcher  matcher.Matcher[T]
	template string
	args     []interface{}
}

func (a as[T]) Matches(value T) bool {
	return a.matcher.Matches(value)
}

func (a as[T]) DescribeTo(description matcher.Description) {
	result := argPattern.ReplaceAllStringFunc(a.template, func(s string) string {
		idx, _ := strconv.Atoi(s)
		description := matcher.StringDescription()
		description.AppendValue(a.args[idx])
		return description.String()
	})
	description.AppendText(result)
}

func (a as[T]) DescribeMismatch(actual T, description matcher.Description) {
	a.matcher.DescribeMismatch(actual, description)
}
