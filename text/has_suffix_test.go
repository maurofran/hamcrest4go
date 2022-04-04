package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasSuffix_Matches(t *testing.T) {
	tests := []struct {
		name     string
		substr   string
		test     string
		expected bool
	}{
		{"not has suffix", "foo", "has suffix", false},
		{"has suffix", "fix", "has suffix", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := HasSuffix(test.substr)
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestHasSuffix_DescribeTo(t *testing.T) {
	// Given
	fixture := HasSuffix("inside")
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `a string with suffix "inside"`, description.String())
}

func TestHasSuffix_DescribeMismatch(t *testing.T) {
	// Given
	fixture := HasPrefix("inside")
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
