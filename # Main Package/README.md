# Main Package 

- Go는 main.go 라는 이름의 약속된 파일만 컴파일이 가능하고 다른 이름의 파일들은 예컨데 learn.go / study.go 같은 개인이 이름을 붙이는 파일들은 컴파일이 불가하고 함수의 모음이나 다른 사람이 컴파일 하도록 코드를 작성하는것은 가능 등 같은 기능에만 가능하다 즉, 나 자신이 컴파일을 요구 할 떄는 main.go 라는 이름을 해야한다. 

```
package main

import "fmt"

func main() {
	fmt.Print("Hello World")
}
```

Go 파일은 컴파일을 진행할때 main package 안에 main.go 를 찾아서 컴파일을 한다.

## 참고
[노마드 코더스](https://nomadcoders.co/go-for-beginners/lectures/1501)

