package entity

import "math"

type Location struct {
	X int8
	Y int8
}

func (l Location) Distance(otherLocation Location) int {
	v1 := math.Pow(float64(l.X-otherLocation.Y), 2)
	v2 := math.Pow(float64(l.Y-otherLocation.Y), 2)

	return int(math.Round(math.Sqrt(v1 + v2)))
}
