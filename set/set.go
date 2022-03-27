// Package set implements an algebraic set and operations for sets.
package set

import "fmt"

// Set represents a set of elements. The elements are integers in range [0..63].
//
// Examples:
// {0, 1, 2, 63} is a valid set.
// {63, 64} is an invalid set.
type Set uint64

// Empty represents an empty set: {}.
const Empty = Set(0)

// String returns a string representation of the set, e.g. "{0, 1, 2}".
func String(s Set) string {
	return ""
}

// IsEmpty returns true iff the set is empty.
//
// Examples:
// IsEmpty({}) -> true
// IsEmpty({5}) -> false.
func IsEmpty(s Set) bool {
	return false
}

// Len returns the number of elements in the set.
//
// Examples:
// Len({}) -> 0
// Len({0, 1, 2}) -> 3
func Len(s Set) int {
	return 0
}

// Elements returns a slice of set elements.
//
// Examples:
// Elements({0, 1, 2}) -> []int{0, 1, 2}
func Elements(s Set) []int {
	return nil
}

// Add returns a new set that contains the integer `n`.
// An error must be return if the integer is not in range [0..63].
//
// Examples:
// Add({0, 1, 2}, 5) -> {0, 1, 2, 5}, nil
// Add({0, 1, 2}, 64) -> {}, error
func Add(s Set, n int) (Set, error) {
	if n > 63 || n < 0 {
		error := fmt.Errorf("%v is not in range [0...63]", n)
	}
	return (1 << n) | s, error

}

// Contains returns true iff the element `n` exists in the set.
//
// Examples:
// Contains({0, 1, 2}, 2) -> true
// Contains({0, 1, 2}, 3) -> false
func Contains(s Set, n int) bool {
	return s&(1<<n) > 0
}

// Remove returns a new set that does not contain the element `n`.
//
// Examples:
// Remove({0, 1, 2}, 1) -> {0, 2}
// Remove({0, 1, 2}, 4) -> {0, 1, 2}
func Remove(s Set, n int) Set {
	return s &^ (1 << n)
}

// Union returns a new set that is a union of two sets.
// The union contains elements that are present in *either* set.
//
// Examples:
// Union({0, 1, 2}, {1, 3, 4}) -> {0, 1, 2, 3, 4}
// Union({0, 1, 2}, {}) -> {0, 1, 2}
func Union(s1, s2 Set) Set {
	return s1 | s2
}

// Intersection returns a new set that is an intersection of two sets.
// The intersection contains elements that are present in *both* sets.
//
// Examples:
// Intersection({0, 1, 2}, {2, 3, 4}) -> {2}
// Intersection({0, 1, 2}, {}) -> {}
func Intersection(s1, s2 Set) Set {
	return Empty
}

// Difference returns a new set that is a difference of two sets.
// The difference contains elements that are not present in *both* sets.
//
// Examples:
// Difference({0, 1, 2}, {1, 5}) -> {0, 2, 5}
func Difference(s1, s2 Set) Set {
	return Empty
}

// Subtract returns a new set that contains elements that are present
// in the first set, but not present in the second set.
//
// Examples:
// Subtract({0, 1, 2}, {1, 4}) -> {0, 2}
func Subtract(s1, s2 Set) Set {
	return Empty
}
