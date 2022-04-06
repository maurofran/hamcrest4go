package text

import (
	"github.com/maurofran/hamcrest4go/is"
	"github.com/maurofran/hamcrest4go/matcher"
)

// IsEmpty matches empty strings.
func IsEmpty() matcher.Matcher[string] {
	return isEmptyInstance
}

// IsNotEmpty matches not empty strings.
func IsNotEmpty() matcher.Matcher[string] {
	return is.Not(IsEmpty())
}

var isEmptyInstance = isEmpty{}

type isEmpty struct {
	matcher.Base[string]
}

func (isEmpty) Matches(value string) bool {
	return value == ""
}

func (isEmpty) DescribeTo(description matcher.Description) {
	description.AppendText("an empty string")
}
