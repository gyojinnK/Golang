package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}
}
