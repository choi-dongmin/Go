# 함수 (Functions)

go 에서 함수를 만들때는 func 함수이름() 리턴 데이터형 {} 의 구성으로 작성된다.

```
package main

import "fmt"

func multiply(a int, b int) int { //매개변수의 타입이 같다면 fanc multiply(a , b int) int {} 같은 형태로도 가능 
	return a * b
}


func main() {
	fmt.Println(multiply(2, 2))
}
```

곱셈을 해주는 multiply 함수를 만들었고 그 함수는 매개변수를 a int ,b int 를 받고 있고 리턴의 형태도 int(정수) 이다.

JAVA 와 비슷한 구조이지만 매개변수의 데이터 타입과 리턴의 데이터 타입이 보다 뒤쪽에 위치한다.   

## 다중 return

go가 매력적인 이유중 하나는 다중 return 값을 가질수 있다는것이다.

```
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	nameLen, upperName := lenAndUpper("dm")
	fmt.Println(nameLen, upperName)
}
```

위 lenAndUpper 함수는 문자열이 인자값으로 들어 왔을때 그 문자열의 길이와 그 문자열을 대문자화 시켜 return 하는 기능을 가진다.

여기서 흥미로운것이 다른 언어들과 다르게 go 는 return 값을 int , string 으로 한번에 반환 한다는 것이다.


```
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
```
len(name) : 매개변수(문자열)의 길이를 구하여 반환 (int type)

strings.ToUpper(name) : strings 라는 패키지의 ToUpper 을 통해 매개변수(문자열)의 대문자화 (string type)

위 두가지의 다른 데이터 타입의 다른 값을 한 함수에서 출력해 낼수 있다.

```
func main() {
	nameLen, upperName := lenAndUpper("dm")
	fmt.Println(nameLen, upperName)
}
```

nameLen, upperName := lenAndUpper("dm") : lenAndUpper 함수에 인자값은 ("dm") 이다. 그리고 return 받은 값은 매개변수와 일치하여 nameLen = int, upperName = string (만약 이 두 변수의 앞뒤가 바뀐다면 nameLen = string, upperName = int이다. )

* 만약 다중 return 을 선언했다면 반듯이 그 수만큼을 사용해야 한다. (2 매개변수, 1 return 불가 / 2 매개변수, 2 return, 1 var 불가)

그러나 무시된 변수를 통해 그 값을 무시해버릴수 있다.

```
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	_, upperName := lenAndUpper("dm")
	fmt.Println(upperName)
}
``` 

위와 같이 _ 를 통해 무시되는 변수를 만들어 출력시 제외시킬수 있다. 

## 다중 인자값 

만약 인자값이 즉, 매개변수로 들어오는 값의 수가 불분명 하다면 매개변수를 할 수없다 그럴때는 아래 코드와 같이 함수를 작성한다.

```
package main

import "fmt"

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("me", "you", "we", "they")
}

결과값 : [me you we they]
```

repeatMe 함수는 매개변수(문자열)로 받는 값들을 그대로 다시 출력하는 함수이다.

```
func repeatMe(words ...string) { //return 값이 없기 때문에 return 의 데이터 타입을 쓰지 않았다. 
	fmt.Println(words)
}
```

중요한 부분은 바로 func repeatMe(words ...string) 부분이다. repeatMe 는 words 라는 매개변수를 가지는데 이 매개변수는 인자값의 개수가 불분명한 문자열이다.(...String) 이다.

이렇게 ... 을 사용해 입력된 인자값을 매개변수에 배열처럼 저장한다.


## naked return

```
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	nameLen, upperName := lenAndUpper("dm")
	fmt.Println(nameLen, upperName)
}
```
위 코드에서 lenAndUpper 함수는 len(name)과 strings.ToUpper(name) 를 return 하고 있다.

그러나 return 을 다른 방식으로 작성할 수 있다.

```
package main

import (
	"fmt"
	"strings"
)

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (lenght int, uppercase string) {
	lenght = len(name) // 위에서 만든 변수 lenght 을 재정의
	uppercase = strings.ToUpper(name) // 위에서 만든 변수 uppercase 을 재정의
	return
}

func main() {
	nameLen, upperName := lenAndUpper("dm")
	fmt.Println(nameLen, upperName)
}
```

위 lenAndUpper 은 return의 데이터 값을 정의 해주는곳에서 변수를 지정하고 있고 코드 안쪽에서 재정의 함으로 같은 이전과 같은 출력값을 가지지만 다른 작성방식으로 코드를 작성했다.

즉 return 값을 코드 안에서 변수를 정의하지 않고 return 의 데이터 타입을 정의하는 곳에서 변수를 정의하는 방식이다.

## defer
- defer 기능은 함수가 그 기능을 다하였을때 추가적으로 어떤 동작을 하도록 하는것이다.

```
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
	lenght = len(name) // 위에서 만든 변수 lenght 을 재정의
	uppercase = strings.ToUpper(name) // 위에서 만든 변수 uppercase 을 재정의
	return
}
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	_, upperName := lenAndUpper("dm")
	fmt.Println(upperName)
	repeatMe("me", "you", "we", "they")

}

결과값 :
끝
DM
[me you we they]
```
만약 lenAndUpper 함수가 그 안의 기능을 완료한 다음 "끝"이라는 단어를 출력하도록 한 코드이다. 

```
defer fmt.Println("끝")
```
defer 를 이용해 함수가 끝날시 실행하고 싶은 기능을 추가한다.

* 끝이 DM 보다 먼저 나오는것은 lenAndUpper 의 return 을 하고 defer 을 실행하고 그 return 값을 
main 함수(fmt.Println(upperName)) 에서 출력하고 있기 때문이다.

## 키워드

- 다중 retrun : go 는 다른 언어와는 다르게 하나의 함수에서 여러개의 return 값을 가질수 있다.

- 다중 인자값 : 입력될 인자값의 개수가 불분명 하고 그 타입이 같을때 ... 을 통해 매개변수에 배열처럼 할당.

- naked return : return 의 데이터 타입을 선언하는 곳에서 변수와 그 데이터 타입을 정의하고 코드 안쪽에서 재정의를 통한다.

- defer : defer 가 속한 함수가 return 을 실행 하고 그 함수의 기능이 끝날고 실행하는 기능  

## 참고 
[노마드 코더스](https://nomadcoders.co/go-for-beginners/lectures/1503) 