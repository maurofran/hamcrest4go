package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"strings"
)

// Contains check if given text contains the given substring.
func Contains(substr string) matcher.Matcher[string] {
	return contains{substr: substr}
}

type contains struct {
	substr string
}

func (m contains) Matches(value string) bool {
	return strings.Contains(value, m.substr)
}

func (m contains) DescribeTo(description matcher.Description) {
	description.AppendText("a string containing '")
	description.AppendText(m.substr)
	description.AppendText("'")
}

func (contains) DescribeMismatch(actual string, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
