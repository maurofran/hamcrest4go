package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"strings"
)

// HasSuffix check if given text starts with the given substring.
func HasSuffix(suffix string) matcher.Matcher[string] {
	return hasSuffix{suffix: suffix}
}

type hasSuffix struct {
	matcher.Base
	suffix string
}

func (m hasSuffix) Matches(value string) bool {
	return strings.HasSuffix(value, m.suffix)
}

func (m hasSuffix) DescribeTo(description matcher.Description) {
	description.AppendText("a string with suffix \"")
	description.AppendText(m.suffix)
	description.AppendText("\"")
}
