package util

// Set represents a collection of items with no duplicates.
type Set[T comparable] map[T]struct{}

// Contains returns true if the given item is in the set.
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

// Insert adds the given item to the set.
// It returns true if the item was newly inserted.
func (s Set[T]) Insert(item T) bool {
	if s.Contains(item) {
		return false
	}
	s[item] = struct{}{}
	return true
}

// Removes item from the set.
// Returns true if the value was present in the set.
func (s Set[T]) Remove(item T) bool {
	if !s.Contains(item) {
		return false
	}
	delete(s, item)
	return true
}
