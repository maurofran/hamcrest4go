package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"strings"
)

// HasPrefix check if given text starts with the given substring.
func HasPrefix(prefix string) matcher.Matcher[string] {
	return hasPrefix{prefix: prefix}
}

type hasPrefix struct {
	matcher.Base
	prefix string
}

func (m hasPrefix) Matches(value string) bool {
	return strings.HasPrefix(value, m.prefix)
}

func (m hasPrefix) DescribeTo(description matcher.Description) {
	description.AppendText("a string with prefix \"")
	description.AppendText(m.prefix)
	description.AppendText("\"")
}
