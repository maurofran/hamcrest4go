package described

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"regexp"
	"strconv"
	"strings"
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
	parts := argPattern.Split(a.template, -1)
	for _, part := range parts {
		if strings.HasPrefix(part, "%") {
			idx, _ := strconv.Atoi(part[1:])
			description.AppendValue(a.args[idx])
		} else {
			description.AppendText(part)
		}
	}
}

func (a as[T]) DescribeMismatch(actual T, description matcher.Description) {
	a.matcher.DescribeMismatch(actual, description)
}
