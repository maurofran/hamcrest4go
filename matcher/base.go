package matcher

type Base struct {
}

func (Base) DescribeMismatch(actual string, description Description) {
	description.AppendText("was ")
	description.AppendValue(actual)
}
