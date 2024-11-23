package utils

import "fmt"

// Set is a collection based on map
type Set struct {
	items map[string]bool
}

// Create a new collection
func NewSet() *Set {
	return &Set{
		items: make(map[string]bool),
	}
}

// Add an element to a collection
func (s *Set) Add(item string) {
	s.items[item] = true
}

// Remove an element from a collection
func (s *Set) Remove(item string) {
	delete(s.items, item)
}

// Checks whether a collection contains an element
func (s *Set) Contains(item string) bool {
	_, exists := s.items[item]
	return exists
}

// Returns the number of elements in a collection
func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) GetSet() *map[string]bool {
	return &s.items
}

// Returns a string representation of the collection
func (s *Set) String() string {
	elements := []string{}
	for item := range s.items {
		elements = append(elements, item)
	}
	return fmt.Sprintf("%v", elements)
}
