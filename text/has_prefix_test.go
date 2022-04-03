package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPrefix_Matches(t *testing.T) {
	tests := []struct {
		name     string
		substr   string
		test     string
		expected bool
	}{
		{"not has prefix", "foo", "has prefix", false},
		{"has prefix", "has", "has prefix", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := HasPrefix(test.substr)
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestHasPrefix_DescribeTo(t *testing.T) {
	// Given
	fixture := HasPrefix("inside")
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `a string with prefix "inside"`, description.String())
}

func TestHasPrefix_DescribeMismatch(t *testing.T) {
	// Given
	fixture := HasPrefix("inside")
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
