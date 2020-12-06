package main

//go 표준 lib
import "fmt"

//프로그램의 entry point
func main() {
	//문자열 concat?
	fmt.Println("go" + "lang")

	//javascript는 이거 된다고..
	//fmt.Println("1" * "2")

	//정수 연산
	fmt.Println("1+1 = ", 1+1)

	//부동소숫점
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	//Boolean 논리 연산
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
