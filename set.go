package set

import (
	"errors"
)

type Set struct {
	items []interface{}
	size  int
	index int
	err   error
}

// Errors returned by Set.
var (
	ErrBadIteration = errors.New("set.Get: needs start interation with Next method")
)

// New returns a new Set containing the given parameters or empty.
func New(items ...interface{}) *Set {
	s := new(Set).Clear()
	for _, v := range items {
		s.Add(v)
		s.size++
	}
	return s
}

// Clear restarts Set with the default "zero" values.
func (s *Set) Clear() *Set {
	s.items = make([]interface{}, 0)
	s.size = 0
	s.index = -1
	return s
}

// Has checks if a given item exists in Set.
func (s *Set) Has(item interface{}) (exist bool, index int) {
	for i, v := range s.items {
		if item == v {
			return true, i
		}
	}

	return false, -1
}

// Add appends a new item in Set.
func (s *Set) Add(item interface{}) {
	exist, _ := s.Has(item)

	if !exist {
		s.items = append(s.items, item)
		s.size++
	}
}

// Delete deletes an existing item in Set.
func (s *Set) Delete(item interface{}) {
	exist, i := s.Has(item)

	if exist {
		s.items = append(s.items[:i], s.items[i+1:]...)
		s.size--
	}
}

// Size returns the number of items in Set.
func (s *Set) Size() int {
	return s.size
}

// Next allows to iterate over Set.
// Use in combination with Get().
func (s *Set) Next() bool {
	if s.index < len(s.items)-1 {
		s.index++

		return true
	}

	s.index = -1
	s.err = nil
	return false
}

// Get returns the current value when interating.
// Use in combination with Next().
func (s *Set) Get() interface{} {
	if s.index > -1 {
		return s.items[s.index]
	}

	s.err = ErrBadIteration
	return nil
}

// Error returns Set's last error.
func (s *Set) Error() error {
	return s.err
}
