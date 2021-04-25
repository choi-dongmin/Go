package main

import (
	"fmt"
	"strings"
)

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }
// 아래와 같은 코드
func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("끝")
	lenght = len(name)                // 위에서 만든 변수 lenght 을 재정의
	uppercase = strings.ToUpper(name) // 위에서 만든 변수 uppercase 을 재정의
	return
}
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	len, upperName := lenAndUpper("dm")
	fmt.Println(len, upperName)
	repeatMe("me", "you", "we", "they")

}
