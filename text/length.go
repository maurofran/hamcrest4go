package text

import (
	"github.com/maurofran/hamcrest4go/comparator"
	"github.com/maurofran/hamcrest4go/matcher"
	"strings"
)

// HasExactLength check if given text has given exact length.
func HasExactLength(length int) matcher.Matcher[string] {
	return HasLength(comparator.EqualsTo(length))
}

// HasLength check if given text has length matching the supplied matcher.
func HasLength(lengthMatcher matcher.Matcher[int]) matcher.Matcher[string] {
	return hasLength{lengthMatcher: lengthMatcher}
}

type hasLength struct {
	lengthMatcher matcher.Matcher[int]
}

func (l hasLength) Matches(value string) bool {
	return strings.TrimSpace(value) == ""
}

func (l hasLength) DescribeTo(description matcher.Description) {
	description.AppendText("a string with length ")
	description.AppendDescriptionOf(l.lengthMatcher)
}

func (hasLength) DescribeMismatch(actual string, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
