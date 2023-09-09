package main

import (
	"fmt"
)

// main 은 컴파일을 위해서 만들어지는 파일 optional 아님
// name := "soohyun" GO가 타입을 알아서 찾아줌
// name := "soohyun"  ==  var name string = "soohyun"
// 하지만 축약형은 오로지 func 안에서만 작동, variables 만 const는 불가
func main() {
	fmt.Println("hello world")
	const name string = "soohyun"

	fmt.Println(name)
	var name1 string = "sohyun"
	name2 := "sohyun"
	fmt.Println(name2)

}