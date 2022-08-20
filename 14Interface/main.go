package main

import (
	"fmt"
)

/*
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius *  c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64  {
	return s.area()
}
*/

type Movie interface {
	movieInfo() (string)
}

type Release struct {
	releaseDate string
}
type Price struct {
	price string
}

func (r Release) movieInfo() string{
	return r.releaseDate
}

func (p Price) movieInfo() string{
	return p.price
}

func getMovieInfo(m Movie)(string) {
	return m.movieInfo()
}
func main() {
	fmt.Println("interface")

	/* circle := Circle{radius:3}
	rectangle := Rectangle{width:2, height:3}

	fmt.Println("Circle area:",  getArea(circle))
	fmt.Println("Rectangle area:", getArea(rectangle)) */

	 release := Release{"15 Aug wed"}
	price := Price{"34.5"}
fmt.Println("Movie info -")
	fmt.Println("Relese Date: ",  getMovieInfo(release))
	fmt.Println("Price: ", getMovieInfo(price)) 
	
}