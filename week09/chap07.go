package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// map을 사용하지 않고 구현

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
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
