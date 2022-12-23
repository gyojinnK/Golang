package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o") // NewReplacer()메소드를 이용하여 해당 문자를 지정 문자로 변환
	fixed := replacer.Replace(broken)         // replacer.Replace()메소드로 변수에 저장 꼭 윗줄과 함께 사용
	fmt.Println(fixed)
	now := time.Now()  // var now time.Time = time.Now()
	year := now.Year() // var year int = now.Year()
	fmt.Println(year)
}
