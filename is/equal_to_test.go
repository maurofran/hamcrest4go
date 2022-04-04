package is

import (
	"fmt"
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

type equalerImpl string

func (e equalerImpl) Equal(other equalerImpl) bool {
	return string(e) == string(other)
}

func (e equalerImpl) String() string {
	return fmt.Sprintf("%q", string(e))
}

func TestEqualTo_Matches(t *testing.T) {
	tests := []struct {
		name     string
		value    equalerImpl
		other    equalerImpl
		expected bool
	}{
		{"not equals", "foo", "baz", false},
		{"equals", "foo", "foo", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			fixture := EqualTo(test.other)
			// When
			actual := fixture.Matches(test.value)
			// Then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestEqualTo_DescribeTo(t *testing.T) {
	// Given
	fixture := EqualTo(equalerImpl("foo"))
	description := matcher.StringDescription()
	// When
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `equal to "foo"`, description.String())
}

func TestEqualTo_DescribeMismatch(t *testing.T) {
	// Given
	fixture := EqualTo(equalerImpl("baz"))
	description := matcher.StringDescription()
	// When
	fixture.DescribeMismatch("foo", description)
	// Then
	assert.Equal(t, `was "foo"`, description.String())
}
