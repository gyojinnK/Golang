package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// var l, w float64
	// l = 10.5
	// w = 5.7
	// fmt.Println("%.1f * %.1f = %2f\n", l, w, 1*w)

	// var l, w = 10.5, 5.7
	// fmt.Printf("%.1f * %.1f = %.2f\n", l, w, 1*w)

	l, w := 10.5, 5.7
	fmt.Printf("%.1f * %.1f = %.2f\n", l, w, 1*w)

	//ZeroVaule
	var inha string
	var testBool bool
	fmt.Println(inha, testBool)

	var a int
	var b float64
	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(a))

	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("헤드 퍼스트 고"))
}
