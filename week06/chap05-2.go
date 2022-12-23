package main

import "fmt"

func test(pa *bool) {
	fmt.Println(*pa, pa)
}

func main() {
	b := [3]string{
		"a",
		"b",
		"c",
	}
	a := true
	//test(a) // 포인터 변수는 주소를 받아야 한다.
	test(&a)
	fmt.Println(b[2])
}
