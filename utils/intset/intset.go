package intset

// IntSet wraps the map-to-empty-struct idiom in Golang for sets
type IntSet struct {
	contents map[int]struct{}
}

// New creates and returns an empty contents map
func New() *IntSet {
	is := IntSet{}
	is.contents = make(map[int]struct{})
	return &is
}

// Add an integer to the set if it's not already present
func (is *IntSet) Add(i int) {
	is.contents[i] = struct{}{}
}

// Remove deletes i from the set if it is present
func (is *IntSet) Remove(i int) {
	delete(is.contents, i)
}

// Contains returns whether i is already in the set
func (is *IntSet) Contains(i int) bool {
	if is == nil {
		return false
	}
	_, exists := is.contents[i]
	return exists
}

// ToSlice returns the elements contained in the set as a new slice
func (is *IntSet) ToSlice() []int {
	slice := make([]int, 0, len(is.contents))
	for i := range is.contents {
		slice = append(slice, i)
	}
	return slice
}
