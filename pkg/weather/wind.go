package weather

import "math"

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

// WindDegreesToDirection transforms given in degrees into a direction
func WindDegreesToDirection(degrees int) string {
	temp := math.Floor((float64(degrees) / 22.5) + 0.5)
	index := int(temp) % 16

	return directions[index]
}
