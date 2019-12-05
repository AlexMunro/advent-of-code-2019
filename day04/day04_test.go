package main

import(
	"testing"
)

func TestValidPassword(t *testing.T) {
	examples := map[int]bool {
		111111: true,
		223450: false,
		123789: false,
	}

	for k, v := range examples {
		result := validPassword(k)
		if result != v {
			t.Errorf("Expected to get %t from %d but got %t", v, k, result)
		}
	}
}

func TestExtraValidPassword(t *testing.T) {
	examples := map[int]bool {
		112233: true,
		123444: false,
		111122: true,
		// extra cases
		112345: true,
		122345: true,
		112223: true,
	}

	for k, v := range examples {
		result := extraValidPassword(k)
		if result != v {
			t.Errorf("Expected to get %t from %d but got %t", v, k, result)
		}
	}
}