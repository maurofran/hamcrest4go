package compare

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreaterThanOrEqualsTo_Matches(t *testing.T) {
	tests := []struct {
		name     string
		ref      int
		value    int
		expected bool
	}{
		{"equals", 5, 5, true},
		{"lesser", 5, 4, false},
		{"greater", 5, 6, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := GreaterThanOrEqualsTo(test.ref)
			// When
			actual := fixture.Matches(test.value)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestGreaterOrEqualsToThan_DescribeTo(t *testing.T) {
	// Given
	fixture := GreaterThanOrEqualsTo(5)
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `a value greater than or equals to 5`, description.String())
}

func TestGreaterThanOrEqualsTo_DescribeMismatch(t *testing.T) {
	// Given
	fixture := GreaterThanOrEqualsTo(5)
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch(8, description)
	// Then
	assert.Equal(t, `8 was greater than or equals to 5`, description.String())
}
