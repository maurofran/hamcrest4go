package text

import (
	"github.com/maurofran/hamcrest4go/is"
	"github.com/maurofran/hamcrest4go/matcher"
	"strings"
)

// IsBlank matches blank strings.
func IsBlank() matcher.Matcher[string] {
	return isBlankInstance
}

// IsNotBlank matches not blank strings.
func IsNotBlank() matcher.Matcher[string] {
	return is.Not(IsBlank())
}

var isBlankInstance = isBlank{}

type isBlank struct {
	matcher.Base[string]
}

func (isBlank) Matches(value string) bool {
	return strings.TrimSpace(value) == ""
}

func (isBlank) DescribeTo(description matcher.Description) {
	description.AppendText("a blank string")
}
