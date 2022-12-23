package main

import (
	"fmt"
	"math"
)

type Cuboid struct {
	length float64
	width  float64
	height float64
}

type Sphere struct {
	radius float64
}

/*
// 직육면체의 부피 및 겉넓이
func (c Cuboid) Volume() float64 {
	return c.height * c.length * c.width
}
func (c Cuboid) Surface() float64 {
	return (c.height * c.length * 2) + (c.length * c.width * 2) + (c.height * c.width * 2)
}
*/

func (c Cuboid) Volume() float64 {
	return c.height * c.length * c.width
}
func (c Cuboid) Surface() float64 {
	return (c.height * c.length * 2) + (c.length * c.width * 2) + (c.height * c.width * 2)
}

// 구의 부피 및 겉넓이
func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(s.radius, 3)
}
func (s Sphere) Surface() float64 {
	return 4.0 * math.Pi * math.Pow(s.radius, 2)
}

// 메소드 오버로딩이 안돼서 인터페이스로 구현하기
type Shape3D interface {
	Volume() float64
	Surface() float64
}

/* func PrintInfo(c Cuboid) {
	fmt.Println(c)
	fmt.Printf("부피는 %0.2f 입니다.\n", c.Volume())
	fmt.Printf("겉넓이는 %0.2f 입니다.\n", c.Surface())
} */

func PrintInfo(s3 Shape3D) {
	fmt.Println(s3)
	fmt.Printf("부피는 %0.2f 입니다.\n", s3.Volume())
	fmt.Printf("겉넓이는 %0.2f 입니다.\n", s3.Surface())
}
func main() {
	c1 := Cuboid{length: 10.0, width: 2.0, height: 5.0}
	PrintInfo(c1)
	s1 := Sphere{radius: 10.0}
	PrintInfo(s1)
}

/* package main

import "fmt"

type CoffeePot string

// stringer interface의 메소드 String()을 구현
func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	coffeePot := CoffeePot("LG 에스프레소 머신")
	fmt.Println(coffeePot.String())
	fmt.Print(coffeePot, "\n")
	fmt.Print(coffeePot, "\n")
	fmt.Printf("%s", coffeePot)
}
*/
