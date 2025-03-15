package weather

import (
	"testing"
)

func TestWindDegrees(t *testing.T) {
	var values = map[int]string{
		100: "E",
		120: "ESE",
		31:  "NNE",
		270: "W",
		72:  "ENE",
		241: "WSW",
		229: "SW",
		336: "NNW",
	}

	for key, value := range values {
		direction := WindDegreesToDirection(key)
		if direction != value {
			t.Errorf("Failed to convert degrees into compass direction! Expected: %s , got %s", value, direction)
		}
	}
}
