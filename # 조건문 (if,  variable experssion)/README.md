# 조건문 (if, if else)

어떠한 조건에 만족하면 코드를 실행시키는 조건문

```
import "fmt"

func canIDrink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}
func main() {
	fmt.Println(canIDrink(1))
}

결과값 : false 
```

다른 언어들과는 다르게 조건을 () 가 아니라 바로 옆에서 작성할 수 있다.

## variable experssion

- variable experssion이란 기능은 if 가 실행되는 순간에 변수를 만들수 있는 기능이다. 

```
import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}
func main() {
	fmt.Println(canIDrink(16))
}

결과값 : true 
```
if 뒤에 바로 변수를 선언하고 if문 안에서 그 변수만을 사용하고 있다.


```
package main

import "fmt"

func canIDrink(age int) bool {
	koreanAge := age + 2
	if koreanAge < 18 {
		return false
	} else {
		return true
	}
}
func main() {
	fmt.Println(canIDrink(16))
}
```

위와 똑같은 문장이다 그러나 다른점은 if 뒤에 변수를 선언하면 오직 if/eles 안에서만 작동 한다는것을 알수있다.


##키워드

- variable experssion : if 문을 선언 할때 if 문 내에서만 동작할 변수를 선언하는것. 

## 참고 
[노마드 코더스](https://nomadcoders.co/go-for-beginners/lectures/1506)