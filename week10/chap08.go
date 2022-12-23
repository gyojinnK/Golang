package main

import (
	"fmt"
	"magazine"
)

func main() {
	e1 := magazine.Employee{Name: "최인하", Salary: 30000}
	// e1.HomeAddress.Street = "인하로"
	addr := magazine.Address{Street: "인하로", City: "수원시"}
	//e1.HomeAddress = addr
	e1.Address = addr
	fmt.Printf("%s에 사는 %s님의 연봉은 %.1f달러 입니다\n", e1.City, e1.Name, e1.Salary)

	var s1 magazine.Subscriber
	s1.Name = "김인하"
	// s1.HomeAddress.City = "안산시"
	// s1.HomeAddress.State = "경기도"
	s1.City = "안산시"
	s1.State = "경기도"
	s1.Rate = 4.99
	s1.Active = true
	fmt.Printf("%s %s에 사는 %s님의 월 구독료는 %.1f달러 입니다", s1.State, s1.City, s1.Name, s1.Rate)
}

// package main

// import "fmt"

// type subscriber struct {
// 	name   string
// 	rate   float64
// 	active bool
// }

// func main() {
// 	var s1 subscriber
// 	s1.name = "김인하"
// 	s1.rate = 4.99
// 	s1.active = true
// 	fmt.Printf("%s님의 월 구독료는 %.1f달러 입니다", s1.name, s1.rate)
// }
