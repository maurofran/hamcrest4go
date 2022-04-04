package matcher

// Matcher is the interface describing a generic matcher.
type Matcher[T any] interface {
	SelfDescribing
	Matches(value T) bool
	DescribeMismatch(actual T, description Description)
}
