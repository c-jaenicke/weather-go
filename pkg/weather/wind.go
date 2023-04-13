package weather

import "math"

// directions contains is a slice of compass designations
var directions = []string{
	"N",
	"NNE",
	"NE",
	"ENE",
	"E",
	"ESE",
	"SE",
	"SSE",
	"S",
	"SSW",
	"SW",
	"WSW",
	"W",
	"WNW",
	"NW",
	"NNW",
}

// WindDegreesToDirection transforms degrees (0 - 360)  into the corresponding compass designation.
func WindDegreesToDirection(degrees int) string {
	temp := math.Floor((float64(degrees) / 22.5) + 0.5)

	return directions[(int(temp) % 16)]
}
