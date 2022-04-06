package matcher

type Base[T any] struct {
}

func (Base[T]) DescribeMismatch(actual T, description Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
