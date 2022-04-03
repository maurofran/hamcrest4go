package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"regexp"
)

// MatchesRegExp check if given text matches the given regular expression.
func MatchesRegExp(pattern *regexp.Regexp) matcher.Matcher[string] {
	return matches{pattern: pattern}
}

// MatchesPattern check if given text matches the given pattern string.
func MatchesPattern(pattern string) matcher.Matcher[string] {
	return MatchesRegExp(regexp.MustCompile(pattern))
}

type matches struct {
	pattern *regexp.Regexp
}

func (m matches) Matches(value string) bool {
	return m.pattern.MatchString(value)
}

func (m matches) DescribeTo(description matcher.Description) {
	description.AppendText("a string matching the pattern '")
	description.AppendText(m.pattern.String())
	description.AppendText("'")
}

func (matches) DescribeMismatch(actual string, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
