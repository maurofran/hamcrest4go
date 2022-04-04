package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBlank_Matches(t *testing.T) {
	tests := []struct {
		name     string
		test     string
		expected bool
	}{
		{"blank string", "   ", true},
		{"empty string", "", true},
		{"not blank string", "foo", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := IsBlank()
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsBlank_DescribeTo(t *testing.T) {
	// Given
	fixture := IsBlank()
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `a blank string`, description.String())
}

func TestIsBlank_DescribeMismatch(t *testing.T) {
	// Given
	fixture := IsBlank()
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}

func TestIsNotBlank_Matches(t *testing.T) {
	tests := []struct {
		name     string
		test     string
		expected bool
	}{
		{"blank string", "   ", false},
		{"empty string", "", false},
		{"not blank string", "foo", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := IsNotBlank()
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsNotBlank_DescribeTo(t *testing.T) {
	// Given
	fixture := IsNotBlank()
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `not a blank string`, description.String())
}

func TestIsNotBlank_DescribeMismatch(t *testing.T) {
	// Given
	fixture := IsNotBlank()
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
