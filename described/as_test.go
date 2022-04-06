package described

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockMatcher[T any] struct {
	mock.Mock
}

func (m *mockMatcher[T]) DescribeTo(description matcher.Description) {
	m.Called(description)
}

func (m *mockMatcher[T]) Matches(value T) bool {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *mockMatcher[T]) DescribeMismatch(actual T, description matcher.Description) {
	m.Called(actual, description)
}

func TestAs_Matches(t *testing.T) {
	// Given
	mm := &mockMatcher[string]{}
	fixture := As[string]("%0 = %1", mm, "Foo", "baz")
	// When
	mm.On("Matches", "aValue").Return(true)
	actual := fixture.Matches("aValue")
	// Then
	assert.True(t, actual)
	mm.AssertExpectations(t)
}

func TestAs_DescribeTo(t *testing.T) {
	// Given
	mm := &mockMatcher[string]{}
	fixture := As[string]("%0 = %1", mm, "Foo", "baz")
	// When
	description := matcher.StringDescription()
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `"Foo" = "baz"`, description.String())
	mm.AssertExpectations(t)
}

func TestAs_DescribeMismatch(t *testing.T) {
	// Given
	mm := &mockMatcher[string]{}
	fixture := As[string]("%0 = %1", mm, "Foo", "baz")
	// When
	description := matcher.StringDescription()
	mm.On("DescribeMismatch", "aValue", description)
	fixture.DescribeMismatch("aValue", description)
	// Then
	mm.AssertExpectations(t)
}
