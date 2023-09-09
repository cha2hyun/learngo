package main

import (
	"fmt"
	"strings"
)

// main 은 컴파일을 위해서 만들어지는 파일 optional 아님
// name := "soohyun" GO가 타입을 알아서 찾아줌
// name := "soohyun"  ==  var name string = "soohyun"
// 하지만 축약형은 오로지 func 안에서만 작동, variables 만 const는 불가
// func main() {
// 	fmt.Println("hello world")
// 	const name string = "soohyun"

// 	fmt.Println(name)
// 	var name1 string = "sohyun"
// 	name2 := "sohyun"
// 	fmt.Println(name2)
// }

// func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

// naked return
func lenAndUpper2(name string) (length int, uppercase string){
	// defer = return하고 바로 실행됨 
	defer fmt.Println("done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return 
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	// array 에 loop 할 수 있게 해줌
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	for i:=0 ; i<len(numbers) ; i++	{
		fmt.
	}
	return 1
}

func canIDring(age int) bool {
	// if 조건문 안에 variables를 생성할 수 있다.
	// if 조건문에서만 사용되는 변수라고 유추할 수 있음
	if koreanAge := age +2; koreanAge < 20 {
		return false
	}
	return true
}

func main() {
	fmt.Println(multiply(2, 2))
	fmt.Println(lenAndUpper("cha2hyun"))
	totalLength, upperName := lenAndUpper("cha2hyun")
	fmt.Println(totalLength, upperName )
	repeatMe("1","2","3","4")
	fmt.Println("hellloo")
	superAdd(1,2,3,4,5,6,7)
}