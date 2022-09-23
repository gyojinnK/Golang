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
	//rand.Seed(98)	//rand은 Seed를 사용하는데 초기 Seed가 있어 변경하지 않으면 같은 수를 return한다. => rand.Seed()로 변경하여 사용.
	seconds := time.Now().Unix() // 난수를 발생시키위해 time.Now()를 활용하여
	rand.Seed(seconds)           // Seed의 값을 현재시간으로 지속적으로 변경하여 난수를 발생시킨다.
	target := rand.Intn(100) + 1

	fmt.Println("1~100사이의 수를 추측해보세용~!")

	success := false
	for i := 0; i < 10; i++ {
		fmt.Println("남은 기회 : ", 10-i, "회")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input) //strconv.Atoi() => 알파벳을 int형으로 변환
		if err != nil {
			log.Fatal(err)
		}

		if num == target {
			success = true
			fmt.Println("정답이에요~! 축하합니다!!!")
			break
		} else if num >= target {
			fmt.Println("정답은 더 작은 수에요 [Down!]")
		} else {
			fmt.Println("정답은 더 큰 수에요 [Up!]")
		}
	}
	if !success {
		fmt.Println("완전 멍청하네요 풉ㅋ 정답은 ", target, "임ㅋ")
	}
}

/*
	rand은 초기 시드가 있어서 Seed를 변경하기 않으면 같은 Seed와 결과가 동일하게 나온다.
	time.Now()를 사용하여 각기 다른 난수를 얻을 수 있다.
*/
