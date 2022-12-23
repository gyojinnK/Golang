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
	fmt.Print("점수 입력 : ")
	reader := bufio.NewReader(os.Stdin) // 외부 데이터를 입력받을때
	in, err := reader.ReadString('\n')  // 들어온 데이터를 변수에 저장, ioError가 발생할 수 있는 부분에는 err변수 같이 반환
	if err != nil {
		log.Fatal(err)
	}
	in = strings.TrimSpace(in)               // 매개변수 in은 "55.9\r\n" 리턴 값은 "55.9"
	score, err := strconv.ParseFloat(in, 64) //float64로 파싱해서 변수에 저장, 마찬가지로 err변수 반환
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if score >= 60 {
		status = "합격"
		// status := "합격"
		// fmt.Println(in, "점은 ", status)
	} else {
		status = "불합격!"
		// status := "불합격!"
		// fmt.Println(in, "점은 ", status)
	}
	fmt.Println(in, "점은 ", status)
	//fmt.Println(i
	// fmt.Print("점수 입력 : ")
	// reader := bufio.NewReader(os.Stdin) // os.Stdin는 자바의 System.in . 즉 표준입력(키보드)
	// //in := reader.ReadString('\n') // assignment mismatch: 1 variable but reader.ReadString returns 2 valu
	// //in, err := reader.ReadString('\n') // err declared but not used
	// //fmt.Println(err) // 출력하려고 하는게 아닌 변수 사용위해서
	// //in, _ := reader.ReadString('\n') // 언더스코어를 통해 변수관련 문제를 해결. #1
	// in, err := reader.ReadString('\n') // if문으로 문제를 해결. #2 옵션
	// //log.Fatal(err) // 로그 기록 및 에러메세지 출력 후 프로그램 종료
	// if err != nil { // 에러가 발생하면
	// 	log.Fatal(err)
	//
	// fmt.Println(in)
}
