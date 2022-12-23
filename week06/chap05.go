package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloat(fn string) ([4]float64, error) {
	var number [4]float64

	f, err := os.Open(fn) // 파일 열기
	if err != nil {
		//log.Fatal(err)
		return number, err
	}

	sc := bufio.NewScanner(f) // 스케너 객체 생성
	i := 0
	for sc.Scan() { // 파일을 스캔
		number[i], err = strconv.ParseFloat(sc.Text(), 64) // 읽어서 float으로 변환
		if err != nil {
			return number, err
		}
		i++
	}

	err = f.Close() // 열었던 파일 닫기
	if err != nil {
		return number, err
	}

	return number, nil
}

func main() {
	numbers, err := GetFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for _, number := range numbers { // _는 인덱스 인자를 생략
		sum = sum + number
	}
	fmt.Printf("평균 : %.2f\n", sum/float64(len(numbers))) // len()메소들 이용하여 처리
}
