package something

import "fmt"


func sayBye(){
	fmt.Println("bye")
}

// Exprot해주고 싶으면 func를 대문자로 시작하면됨 그래서 println이 Println 임
func SayHello(){
	fmt.Println("hello")
}