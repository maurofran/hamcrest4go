package compare

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualsTo_Matches(t *testing.T) {
	tests := []struct {
		name     string
		ref      int
		value    int
		expected bool
	}{
		{"equals", 5, 5, true},
		{"lesser", 5, 4, false},
		{"greater", 5, 6, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := EqualsTo(test.ref)
			// When
			actual := fixture.Matches(test.value)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestEqualsTo_DescribeTo(t *testing.T) {
	// Given
	fixture := EqualsTo(5)
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `a value equal to 5`, description.String())
}

func TestEqualsTo_DescribeMismatch(t *testing.T) {
	// Given
	fixture := EqualsTo(5)
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch(8, description)
	// Then
	assert.Equal(t, `8 was equal to 5`, description.String())
}
