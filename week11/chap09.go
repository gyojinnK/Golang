// 리시버 매개변수를 call by reference
package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n = *n * 2
	//사용자 지정 타입의 메소드를 호출할 때 포인터가 가리키는 값을 적어 줄 때 "&"를 생략한다.
}

func main() {
	number := Number(4)
	fmt.Println("원래 수 : ", number)
	number.Double()
	fmt.Println("두배가 된 수 : ", number)
}

// 리시버 매개변수가 call by value임
/* package main

import "fmt"

type Number int

func (n Number) Double() {
	n = n * 2
}

func main() {
	number := Number(4)
	fmt.Println("원래 수 : ", number)
	number.Double()
	fmt.Println("두배가 된 수 : ", number)
}
*/
