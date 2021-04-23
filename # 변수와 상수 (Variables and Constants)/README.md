# 변수와 상수 (Variables and Constants)

- 변수 : 값을 담고 수정할 수 있는 일종의 상자와 같은 개념
- 상수 : 값을 바꿀수 없는 변수  

기본적으로 go 에서 상수를 지정 할 떄는 const , 변수를 지정 할 때는 var 를 쓴다. 또한 go 는 JAVA 와 같이 데이터 타입또한 지정 해주어야 한다.

## 상수(Constants)

```
package main

import "fmt"

func main() {
	const me string = "DM"
	// me = "YOU" 상수이기 때문에 값 할당 불가.
	fmt.Println(me)
}

```

## 변수 (Variables)
```
package main

import "fmt"

func main() {
	var me string = "DM"
	me = "Choi"
	fmt.Println(me)
}
```

변수 me 에 문자열 DM 을 변수값으로 할당해 주고 바로 밑에서 변수값 DM 을 Choi 로 변경하고 있다. 그렇기 때문에 만약 go 를 실행하면 Choi 라는 값이 출력이 된다.

이처럼 변수는 값을 한번 할당한 후에도 계속해서 바꾸어줄 수 있다. 

코드를 작성할 때 변수는 정말 많이 쓰이기 때문에 go 에서는 조금 더 간결하게 변수를 선언할 수 있다.
```
package main

import "fmt"

func main() {
	me := "DM"
	me = "Choi"
	fmt.Println(me)
}
``` 

위의 코드는 var me string ="DM" 을 me := "me" 로 간결하게 나타내고 있다 또한 이때는 데이터 타입을 선언하지 않아도 go에서 데이터 타입을 찾아 선언한다.

그리고 축약형을 통해 go 가 데이터 타입을 지정했다면 그 값은 변경이 불가하다.

## 참고 
[노마드 코더스](https://nomadcoders.co/go-for-beginners/lectures/1502)

