package text

import (
	"github.com/maurofran/hamcrest4go/compare"
	"github.com/maurofran/hamcrest4go/matcher"
)

// HasExactLength check if given text has given exact length.
func HasExactLength(length int) matcher.Matcher[string] {
	return HasLength(compare.EqualsTo(length))
}

// HasLength check if given text has length matching the supplied matcher.
func HasLength(lengthMatcher matcher.Matcher[int]) matcher.Matcher[string] {
	return hasLength{lengthMatcher: lengthMatcher}
}

type hasLength struct {
	matcher.Base[string]
	lengthMatcher matcher.Matcher[int]
}

func (l hasLength) Matches(value string) bool {
	return l.lengthMatcher.Matches(len(value))
}

func (l hasLength) DescribeTo(description matcher.Description) {
	description.AppendText("a string with length with ")
	description.AppendDescriptionOf(l.lengthMatcher)
}
