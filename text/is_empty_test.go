package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty_Matches(t *testing.T) {
	tests := []struct {
		name     string
		test     string
		expected bool
	}{
		{"blank string", "   ", false},
		{"empty string", "", true},
		{"not empty string", "foo", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := IsEmpty()
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsEmpty_DescribeTo(t *testing.T) {
	// Given
	fixture := IsEmpty()
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `an empty string`, description.String())
}

func TestIsEmpty_DescribeMismatch(t *testing.T) {
	// Given
	fixture := IsEmpty()
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}

func TestIsNotEmpty_Matches(t *testing.T) {
	tests := []struct {
		name     string
		test     string
		expected bool
	}{
		{"blank string", "   ", true},
		{"empty string", "", false},
		{"not empty string", "foo", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := IsNotEmpty()
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsNotEmpty_DescribeTo(t *testing.T) {
	// Given
	fixture := IsNotEmpty()
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `not an empty string`, description.String())
}

func TestIsNotEmpty_DescribeMismatch(t *testing.T) {
	// Given
	fixture := IsNotEmpty()
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
