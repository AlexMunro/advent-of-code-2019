package location

type Location struct {
	X, Y int
}

func Gradient(l1, l2 Location) float64 {
	return float64(l1.Y-l2.Y) / float64(l1.X-l2.X)
}

type LocationSet struct {
	contents map[Location]struct{}
}

func New(initial Location) *LocationSet {
	ls := LocationSet{}
	ls.contents = map[Location]struct{}{}
	ls.contents[initial] = struct{}{}
	return &ls
}

func NewEmptySet() *LocationSet {
	ls := LocationSet{}
	ls.contents = map[Location]struct{}{}
	return &ls
}

func FromSlice(locations []Location) *LocationSet {
	ls := NewEmptySet()
	for _, loc := range locations {
		ls.AddLoc(loc)
	}
	return ls
}

func (ls *LocationSet) Contains(loc Location) bool {
	_, contains := ls.contents[loc]
	return contains
}

func (ls *LocationSet) AddLoc(loc Location) {
	ls.contents[loc] = struct{}{}
}

func (ls *LocationSet) RemoveLoc(loc Location) {
	delete(ls.contents, loc)
}

func (ls *LocationSet) Size() int {
	return len(ls.contents)
}

func (ls *LocationSet) ToSlice() []Location {
	slice := make([]Location, 0, len(ls.contents))
	for l := range ls.contents {
		slice = append(slice, l)
	}
	return slice
}

// Difference mutates the existing set
func (ls *LocationSet) Difference(other *LocationSet) {
	for l, _ := range other.contents {
		ls.RemoveLoc(l)
	}
}

func (ls *LocationSet) Filter(predicate func(Location) bool) []Location {
	slice := []Location{}
	for loc, _ := range ls.contents {
		if predicate(loc) {
			slice = append(slice, loc)
		}
	}
	return slice
}
