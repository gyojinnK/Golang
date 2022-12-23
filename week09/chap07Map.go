package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// map을 사용하여 구현

func GetStrings(fn string) ([]string, error) {
	var lines []string //길이를 정하지 않고 선언 -> ArrayList
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(f)
	//i := 0
	for sc.Scan() {
		line := sc.Text()           // 한 줄 읽기
		lines = append(lines, line) // lines 슬라이스에 line 저장
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	if sc.Err() != nil {
		return nil, sc.Err()
	}
	return lines, nil
}

func main() {
	lines, err := GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
		// 초기값인 zero값을 이용하여 count 값 변환
	}

	for name, count := range counts { //map자료형 for돌릴 때 -> for key, value
		fmt.Printf("%s: %d\n", name, count)
	}
}
