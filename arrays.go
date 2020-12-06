package main

import "fmt"

func main() {

	//배열의 크기가 5인 int 형 배열
	//초기값 0
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	//배열 선언시 초기값 지정
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//2차원 배열 선언
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
