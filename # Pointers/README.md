# Pointers

포인터는 값이 저장된 메모리 주소에 접근하는 것을 말한다.

```
package main

import "fmt"

func main() {
	a := 2
	b := a
	a = 10
	fmt.Println(a, b)
}


결과값 : 10 2
```

위 경우에 b 가 선언될때 a의 값은 2 였기 떄문에 b 가 선언된 후에 a 의 값을 변화 시켜도 b에는 영향을주지 못한다.

하지만 포인터를 이용해 b의 값을 a의 주소로 선언한다면 뒤에서 a 의 값이 바뀐다면 b도 같이 바뀔것이다.

```
package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(&a, b)
}
``` 

위 코드를 보면 변수 b 에 &a(a 변수의 데이터 주소값) 을 넣어주고 있고 

&a 와 b 를 출력하고 있다. 이렇게 출력할시 같은 주소값을 출력한다.

이 기능을 이용해 만약 이전 코드처럼 변수가 뒤에서 변경이 되더라도 그 변수값을 그대로 가져올수 있게 만들어보자.

```
package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 10
	fmt.Println(a, *b)
}

결과값 : 10 10 
```

위 처럼 &a 를 b 변수에 담아주고나서 a 를 수정하고 출력하는곳에서 *b(*는 들어온 &:주소의 데이터를 본다는뜻)하면 원하는 값이 출력이 된다.

즉 우리는 변수 b 를 보고 있지만 변수 b 는 &a 이고 &a 는 10 이라는것을 *b 를 통해 알 수 있다.

또한 b 는 &a 를 가르키고 있기 때문에 값을 b에서도 값을 변경 할 수 있다.

```
package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 10
	*b = 20
	fmt.Println(a, *b)
}

결과값 : 20 20
```

위 코드는 *b 즉 b가 가르키는 주소의 데이터를 변경하고 있음으로 a 의 데이터 또한 변경된다.

이럴때 우리는 pointer b 를 통해 a를 업데이트 하였다고 한다.  

## 키워드

& : 변수가 데이터를 저장하기 위한 주소를 나타내어 주는 기호

* : 변수 주소 안에 어떠한 데이터가 있는지 확인하는 기호

## 참고 
[노마드 코더스](https://nomadcoders.co/go-for-beginners/lobby)