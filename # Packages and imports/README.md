# Packages and imports

만약 우리가 go 에서 Hello World 를 출력하고 싶다면 어떻게 해야할까?

```
ackage main

import "fmt"

func main() {
	fmt.Print("Hello World")
}
```

위와 같은 방식으로 해야한다 JAVA 혹은 다른 언어들과 다르게 출력을 위해선 fmt 라는 package 를 import 해야한다.

fmt 는 기본적인 API(Println, Scanner..) 를 제공하는 패키지이다. 

그런데 fmt.Print 와 같이 우리가 import 해서 사용하는 함수들이 모두 대문자로 이루어져 있는것을 발견 할 수있다.

그 이유는 Go 는 만약 어떠한 패키지에서 export 하려한다면 함수의 첫 문자를 대문자로 지정해야 export가 가능하다.

즉, 소문자로 함수를 만들시 privet 함수가 만들어지게 된다.

```
파일명 : something.go

package something

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func SayHello() {
	fmt.Println("Hello")
}
```
```
파일명 : main.go

package main

import (
	"fmt"

	"github.com/DM/something"
)

func main() {
	fmt.Print("Hello World")
	something.SayHello()
	// something.sayHello() privet 함수이기 때문에 작동하지 않음
}
```

이러한 이유로 Go 에서는 대문자로 시작하는 함수들을 많이 볼수가 있다.


## 참고
[노마드 코더스](https://nomadcoders.co/go-for-beginners/lectures/1502)