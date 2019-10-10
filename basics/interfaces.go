package main
import ("fmt"
				"math")

//Define interface
type Shape interface {
	area() float32
}

type Circle struct {
	a,b,radius float32
}

type Rectangle struct {
	w,h float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float32 {
	return r.h * r.w
}

func getArea(s Shape) float32 {
	return s.area()
}

func main() {
	circle := Circle{a:0,b:0,radius:7}
	rectangle := Rectangle{w:5,h:10}

	fmt.Printf("Circle Area is %f\n", getArea(circle))
	fmt.Printf("Rectangle Area is %f\n", getArea(rectangle))
}
