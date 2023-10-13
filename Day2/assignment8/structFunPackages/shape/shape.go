package shape

import (
	"go-training/Day2/assignment8/structFunPackages/model"
	"math"
)

func AreaOfCircle(c model.Circle) float64 {

	return math.Pi * c.Radius * c.Radius
}
