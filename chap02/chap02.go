package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	now := time.Now()
	year := now.Year()
	fmt.Println(year)
}
