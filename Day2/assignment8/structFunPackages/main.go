package main

import (
	"fmt"
	"go-training/Day2/assignment8/structFunPackages/model"
	"go-training/Day2/assignment8/structFunPackages/shape"
)

func main() {

	circle := model.Circle{Radius: 6.0}

	area := shape.AreaOfCircle(circle)

	fmt.Printf("area of circle is %.2f is %.2f\n", circle.Radius, area)
}
