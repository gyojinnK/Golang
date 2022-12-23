package main

import "fmt"

func main() {
	d := make([]int, 3) // 배열을 초기화하여 생성하는 방법
	var e []int
	fmt.Println(d, len(d))
	fmt.Println(e, len(e))
	fmt.Printf("%#v %#v\n", d, e)

	fmt.Println(d, len(d))

	fmt.Println(e, len(e))

	a := []int{1, 3, -9, 7}
	b := a[1:4]
	fmt.Println(b)
	//c := a[1:5] // 파이썬과 다르기 런타임 panic 발생
	//fmt.Println(c)
	c := a[1:] // 파이썬과 동일
	fmt.Println(c)
}
