package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade : ")
	reader := bufio.NewReader(os.Stdin) // os.Stdin은  자바의  System.in과 같음
	//input := reader.ReadString('\n')    // \n이 입력되기 전의 문자를 입력값으로 하겠다는 뜻
	// ^assignment mismatch: 1 variable but reader.ReadString returns 2 values
	//input, err := reader.ReadString('\n) => err에 <nil>을 받아 원치않는 출력 발생
	/*
		<nil>을 가리는 법
			1. _를 에러변수명으로 하여 모든 에러를 무시한다.
			2. 조건문에 의해서 걸러낸다.
	*/
	//input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')

	if err != nil { //에러가 발생한다면
		log.Fatal(err) //log.Fatal()은 실행되는 순간 log를 남기고 에러메세지 띄우고 강제종료
	}
	input = strings.TrimSpace(input)            // 매개변수 input은 "00.0\r\n" TrimSpace(input)의 리턴값은 "00.0"
	score, err := strconv.ParseFloat(input, 64) //string을 형변환 할 때
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "점 => 합격"
	} else {
		status = "점 => 불합격"
	}
	fmt.Println(input, status)

	fmt.Println(input)
}
