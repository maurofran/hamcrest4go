package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains_Matches(t *testing.T) {
	tests := []struct {
		name     string
		substr   string
		test     string
		expected bool
	}{
		{"not contains", "foo", "search inside me", false},
		{"contains", "inside", "search inside me", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := Contains(test.substr)
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestContains_DescribeTo(t *testing.T) {
	// Given
	fixture := Contains("inside")
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `a string containing "inside"`, description.String())
}

func TestContains_DescribeMismatch(t *testing.T) {
	// Given
	fixture := Contains("inside")
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
