package main

import (
	"testing"
)

type example struct {
	maze               []string
	shortestPathLength int
}

func TestShortestPath(t *testing.T) {
	examples := []example{
		example{
			maze: []string{
				"#########",
				"#b.A.@.a#",
				"#########",
			},
			shortestPathLength: 8,
		},
		example{
			maze: []string{
				"########################",
				"#f.D.E.e.C.b.A.@.a.B.c.#",
				"######################.#",
				"#d.....................#",
				"########################",
			},
			shortestPathLength: 86,
		},
		example{
			maze: []string{
				"########################",
				"#...............b.C.D.f#",
				"#.######################",
				"#.....@.a.B.c.d.A.e.F.g#",
				"########################",
			},
			shortestPathLength: 132,
		},
		example{
			maze: []string{
				"#################",
				"#i.G..c...e..H.p#",
				"########.########",
				"#j.A..b...f..D.o#",
				"########@########",
				"#k.E..a...g..B.n#",
				"########.########",
				"#l.F..d...h..C.m#",
				"#################",
			},
			shortestPathLength: 136,
		},
		example{
			maze: []string{
				"########################",
				"#@..............ac.GI.b#",
				"###d#e#f################",
				"###A#B#C################",
				"###g#h#i################",
				"########################",
			},
			shortestPathLength: 81,
		},
	}

	for _, example := range examples {
		actual := shortestPath(example.maze)
		if example.shortestPathLength != actual {
			t.Errorf("Expected maze to have length %d but had length %d\n", example.shortestPathLength, actual)
		}
	}
}

type encodeExample struct {
	keys     map[rune]struct{}
	encoding int
}

func TestEncodeKeys(t *testing.T) {
	examples := []encodeExample{
		encodeExample{keys: map[rune]struct{}{'a': struct{}{}}, encoding: 0b1},
		encodeExample{keys: map[rune]struct{}{'a': struct{}{}, 'c': struct{}{}}, encoding: 0b101},
		encodeExample{keys: map[rune]struct{}{'a': struct{}{}, 'e': struct{}{}}, encoding: 0b10001},
	}

	for _, example := range examples {
		actual := encodeKeys(example.keys)
		if actual != example.encoding {
			t.Errorf("Expected keys to be encoded as %v but were instead encoded as %v", example.encoding, actual)
		}
	}
}
