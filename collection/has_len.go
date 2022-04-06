package collection

import (
	"github.com/maurofran/hamcrest4go/compare"
	"github.com/maurofran/hamcrest4go/matcher"
)

// HasLen gets a matcher for the length matching the provided value.
func HasLen[T any](length int) matcher.Matcher[[]T] {
	return HasLenMatching[T](compare.EqualsTo(length))
}

// HasLenMatching gets a matcher for the length matching the provided matcher.
func HasLenMatching[T any](lenMatcher matcher.Matcher[int]) matcher.Matcher[[]T] {
	return hasLen[T]{
		lenMatcher: lenMatcher,
	}
}

type hasLen[T any] struct {
	lenMatcher matcher.Matcher[int]
}

func (h hasLen[T]) Matches(value []T) bool {
	return h.lenMatcher.Matches(len(value))
}

func (h hasLen[T]) DescribeTo(description matcher.Description) {
	description.AppendText("a slice or array with length ")
	description.AppendDescriptionOf(h.lenMatcher)
}

func (h hasLen[T]) DescribeMismatch(value []T, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(value)
}

func HasMapLen[K comparable, V any](length int) matcher.Matcher[map[K]V] {
	return HasMapLenMatching[K, V](compare.EqualsTo(length))
}

// HasMapLenMatching gets a mapper checking for map length.
func HasMapLenMatching[K comparable, V any](lenMatcher matcher.Matcher[int]) matcher.Matcher[map[K]V] {
	return hasMapLen[K, V]{lenMatcher: lenMatcher}
}

type hasMapLen[K comparable, V any] struct {
	lenMatcher matcher.Matcher[int]
}

func (h hasMapLen[K, V]) Matches(value map[K]V) bool {
	return h.lenMatcher.Matches(len(value))
}

func (h hasMapLen[K, V]) DescribeTo(description matcher.Description) {
	description.AppendText("a slice or array with length ")
	description.AppendDescriptionOf(h.lenMatcher)
}

func (h hasMapLen[K, V]) DescribeMismatch(value map[K]V, description matcher.Description) {
	description.AppendText("was ")
	description.AppendValue(value)
}
