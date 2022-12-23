package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloat(fn string) ([]float64, error) {
	var numbers []float64 //길이를 정하지 않고 선언 -> ArrayList
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(f)
	//i := 0
	for sc.Scan() {
		//numbers[i], err = strconv.ParseFloat(sc.Text(), 64)
		number, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number) // append()메소드를 이용하여 리스트에 값 추가
		//i++
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	if sc.Err() != nil {
		return nil, sc.Err()
	}
	return numbers, nil
}
func main() {
	numbers, err := GetFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		fmt.Println(number)
		sum = sum + number
	}
	fmt.Printf("평균 : %.2f\n", sum/float64(len(numbers)))
}
