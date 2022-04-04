package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasExactLength_Matches(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		test     string
		expected bool
	}{
		{"equal length", 3, "foo", true},
		{"different length", 3, "fooBar", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := HasExactLength(test.length)
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestHasExactLength_DescribeTo(t *testing.T) {
	// Given
	fixture := HasExactLength(3)
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `a string with length with a value equal to 3`, description.String())
}

func TestHasExactLength_DescribeMismatch(t *testing.T) {
	// Given
	fixture := IsBlank()
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
