package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockSelfDescribing struct {
	mock.Mock
}

func (m *MockSelfDescribing) DescribeTo(description Description) {
	m.Called(description)
}

func TestStringDescription_AppendText(t *testing.T) {
	// Given
	fixture := StringDescription()
	// When
	fixture.AppendText("a text")
	// Then
	assert.NoError(t, fixture.Err())
	assert.Equal(t, "a text", fixture.String())
}

func TestStringDescription_AppendDescriptionOf(t *testing.T) {
	// Given
	fixture := StringDescription()
	value := new(MockSelfDescribing)
	value.On("DescribeTo", fixture)
	// When
	fixture.AppendDescriptionOf(value)
	// Then
	value.AssertExpectations(t)
}

func TestStringDescription_AppendValue(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected string
	}{
		{"nil value", nil, "nil"},
		{"empty string", "", `""`},
		{"string with value", "a value", `"a value"`},
		{"a rune", rune(65), `"A"`},
		{"an empty slice", []interface{}{}, `[]`},
		{"a slice with values", []int{1, 2, 3}, `[1, 2, 3]`},
		{"an empty array", [...]interface{}{}, `[]`},
		{"an array with values", [...]interface{}{"a", "b"}, `["a", "b"]`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := StringDescription()
			// When
			fixture.AppendValue(test.value)
			// Then
			assert.NoError(t, fixture.Err())
			assert.Equal(t, test.expected, fixture.String())
		})
	}

	t.Run("a self describing value", func(t *testing.T) {
		// Given
		fixture := StringDescription()
		value := new(MockSelfDescribing)
		// When
		value.On("DescribeTo", fixture)
		fixture.AppendValue(value)
		// Then
		value.AssertExpectations(t)
	})

	t.Run("a slice of self describing values", func(t *testing.T) {
		// Given
		fixture := StringDescription()
		value := new(MockSelfDescribing)
		// When
		value.On("DescribeTo", fixture)
		fixture.AppendValue([]interface{}{value})
		// Then
		value.AssertExpectations(t)
	})
}
