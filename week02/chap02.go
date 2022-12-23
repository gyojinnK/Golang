package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//rand.Seed(97)
	// seconds := time.Now().Unix()
	// rand.Seed(seconds)
	rand.Seed(time.Now().Unix()) //Seed() 재설정

	target := rand.Intn(100) + 1 // 1 ~ 100
	//fmt.Println(target)
	fmt.Println("1에서 100사이의 수를 추측하세요")

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("남은 추측 횟수 : ", 10-guesses, "회")

		reader := bufio.NewReader(os.Stdin) // 외부 데이터 받기
		in, err := reader.ReadString('\n')  // 외부 데이터 읽기
		if err != nil {
			log.Fatal(err) //애러 출력 후 프로그램 종료
		}
		in = strings.TrimSpace(in)
		number, err := strconv.Atoi(in)
		if err != nil {
			log.Fatal(err)
		}

		if number == target {
			success = true
			fmt.Println("정답입니다. 축하해요~")
			break
		} else if number > target {
			fmt.Println("정답은 입력한 수 보다 작습니다. Down!")
		} else {
			fmt.Println("정답은 입력한 수 보다 큽니다. Up!")
		}
	}
	if !success {
		fmt.Println("게임 종료. 정답은 ", target)
	}
}
