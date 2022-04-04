package is

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilOrZero_Matches(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{"zero", "", true},
		{"not zero", "foo", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := NilOrZero[string]()
			// When
			actual := fixture.Matches(test.value)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestNilOrZero_DescribeTo(t *testing.T) {
	// Given
	fixture := NilOrZero[string]()
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `nil or zero value`, description.String())
}

func TestNilOrZero_DescribeMismatch(t *testing.T) {
	// Given
	fixture := NilOrZero[string]()
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}

func TestNotNilOrZero_Matches(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{"zero", "", false},
		{"not zero", "foo", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := NotNilOrZero[string]()
			// When
			actual := fixture.Matches(test.value)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestNotNilOrZero_DescribeTo(t *testing.T) {
	// Given
	fixture := NotNilOrZero[string]()
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `not nil or zero value`, description.String())
}

func TestNotNilOrZero_DescribeMismatch(t *testing.T) {
	// Given
	fixture := NilOrZero[string]()
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
