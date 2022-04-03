package text

import "github.com/maurofran/hamcrest4go/matcher"

// IsEmpty matches empty strings
func IsEmpty() matcher.Matcher[string] {
	return isEmptyInstance
}

var isEmptyInstance = isEmpty{}

type isEmpty struct {
}

func (isEmpty) Matches(value string) bool {
	return value == ""
}

func (isEmpty) DescribeTo(description matcher.Description) {
	description.AppendText("an empty string")
}

func (isEmpty) DescribeMismatch(actual string, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
