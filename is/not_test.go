package is

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockMatcher struct {
	mock.Mock
}

func (m *mockMatcher) Matches(value string) bool {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *mockMatcher) DescribeTo(description matcher.Description) {
	_ = m.Called(description)
}

func (m *mockMatcher) DescribeMismatch(actual string, description matcher.Description) {
	_ = m.Called(actual, description)
}

func TestNot_Matches(t *testing.T) {
	// Given
	m := new(mockMatcher)
	fixture := Not[string](m)
	// When
	m.On("Matches", "foo").Return(true)
	actual := fixture.Matches("foo")
	// Then
	assert.Equal(t, false, actual)
	m.AssertExpectations(t)
}

func TestNot_DescribeTo(t *testing.T) {
	// Given
	m := new(mockMatcher)
	fixture := Not[string](m)
	description := matcher.StringDescription()
	// When
	m.On("DescribeTo", description)
	fixture.DescribeTo(description)
	// Then
	assert.Equal(t, `not `, description.String())
	m.AssertExpectations(t)
}

func TestNot_DescribeMismatch(t *testing.T) {
	// Given
	m := new(mockMatcher)
	fixture := Not[string](m)
	description := matcher.StringDescription()
	// When
	m.On("DescribeMismatch", "foo", description)
	fixture.DescribeMismatch("foo", description)
	// Then
	m.AssertExpectations(t)
}
