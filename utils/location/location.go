package location

import "fmt"

type Location struct {
	X, Y int
}

type Direction = int

const (
	North Direction = 1
	South Direction = 2
	West  Direction = 3
	East  Direction = 4
)

// Head in the given direction to return a new Location
func (loc Location) Head(dir Direction) Location {
	switch dir {
	case North:
		return Location{X: loc.X, Y: loc.Y + 1}
	case South:
		return Location{X: loc.X, Y: loc.Y - 1}
	case East:
		return Location{X: loc.X + 1, Y: loc.Y}
	case West:
		return Location{X: loc.X - 1, Y: loc.Y}
	default:
		panic(fmt.Sprintf("%d is not a known direction", dir))
	}
}

func Gradient(l1, l2 Location) float64 {
	return float64(l1.Y-l2.Y) / float64(l1.X-l2.X)
}

func MinX(locations []Location) int {
	minX := locations[0].X
	for _, loc := range locations {
		if minX > loc.X {
			minX = loc.X
		}
	}
	return minX
}

func MaxX(locations []Location) int {
	maxX := locations[0].X
	for _, loc := range locations {
		if maxX < loc.X {
			maxX = loc.X
		}
	}
	return maxX
}

func MinY(locations []Location) int {
	minY := locations[0].Y
	for _, loc := range locations {
		if minY > loc.Y {
			minY = loc.Y
		}
	}
	return minY
}

func MaxY(locations []Location) int {
	maxY := locations[0].Y
	for _, loc := range locations {
		if maxY < loc.Y {
			maxY = loc.Y
		}
	}
	return maxY
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
