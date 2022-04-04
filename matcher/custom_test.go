package matcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustom_Matches(t *testing.T) {
	// Given
	fn := func(value string) bool {
		return "foo" == value
	}
	fixture := Custom("description", fn)
	// When
	actual := fixture.Matches("foo")
	// Then
	assert.True(t, actual)
}

func TestCustom_DescribeTo(t *testing.T) {
	// Given
	fn := func(value string) bool {
		return false
	}
	fixture := Custom("description", fn)
	// When
	description := StringDescription()
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, "description", description.String())
}

func TestCustom_DescribeMismatch(t *testing.T) {
	// Given
	fn := func(value string) bool {
		return false
	}
	fixture := Custom("description", fn)
	// When
	description := StringDescription()
	fixture.DescribeMismatch("actual", description)
	// Then
	assert.Equal(t, `was "actual"`, description.String())
}
