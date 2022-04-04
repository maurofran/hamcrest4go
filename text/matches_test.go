package text

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchesPattern_Matches(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		test     string
		expected bool
	}{
		{"blank string", "\\d+", "   ", false},
		{"empty string", "\\d+", "", false},
		{"not empty string", "\\d+", "foo", false},
		{"matching string", "\\d+", "123", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := MatchesPattern(test.pattern)
			// When
			actual := fixture.Matches(test.test)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestMatchesPattern_DescribeTo(t *testing.T) {
	// Given
	fixture := MatchesPattern("\\d+")
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `a string matching the pattern "\d+"`, description.String())
}

func TestMatchesPattern_DescribeMismatch(t *testing.T) {
	// Given
	fixture := MatchesPattern("\\d+")
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
