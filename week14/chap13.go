package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	//fmt.Println(url, "주소에 접근 중...")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//channel <- len(body)
	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{
		"http://www.inhatc.ac.kr",
		"http://www.naver.com",
		"http://www.daum.net",
		"http://www.nate.com",
	}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("주소 %s에서 받은 문자 개수는 %d입니다.\n", page.URL, page.Size)
	}
}

/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, channel chan int) {
	fmt.Println(url, "주소에 접근 중...")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //메인함수가 종료되기 직전까지 네트워크 커넥션 유지하다 연결 해제
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
	//fmt.Println(string(body))
}

func main() {
	sizes := make(chan int)
	urls := []string{
		"http://www.inhatc.ac.kr",
		"http://www.naver.com",
		"http://www.daum.net",
		"http://www.nate.com",
	}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
} */

/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, channel chan int) {
	fmt.Println(url, "주소에 접근 중...")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //메인함수가 종료되기 직전까지 네트워크 커넥션 유지하다 연결 해제
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
	//fmt.Println(string(body))
}

func main() {
	sizes := make(chan int)
	go responseSize("http://www.inhatc.ac.kr", sizes) // go키워드를 붙여 스레드 할당
	go responseSize("http://www.naver.com", sizes)
	go responseSize("http://www.daum.net", sizes)
	go responseSize("http://www.nate.com", sizes)
	//time.Sleep(5 * time.Second) // 스레드의 지연시간 설정
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
}
*/
/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println(url, "주소에 접근 중...")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //메인함수가 종료되기 직전까지 네트워크 커넥션 유지하다 연결 해제
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
	//fmt.Println(string(body))
}

func main() {
	go responseSize("http://www.inhatc.ac.kr") // go키워드를 붙여 스레드 할당
	go responseSize("http://www.naver.com")
	go responseSize("http://www.daum.net")
	go responseSize("http://www.nate.com")
	time.Sleep(5 * time.Second) // 스레드의 지연시간 설정
} */

/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //메인함수가 종료되기 직전까지 네트워크 커넥션 유지하다 연결 해제
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
	//fmt.Println(string(body))
}

func main() {
	responseSize("https://www.inhatc.ac.kr")
	responseSize("https://www.naver.com")
	responseSize("https://www.daum.net")
	responseSize("https://www.nate.com")
}
*/
/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.inhatc.ac.kr")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //메인함수가 종료되기 직전까지 네트워크 커넥션 유지하다 연결 해제
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	fmt.Println(len(body))
}
*/
