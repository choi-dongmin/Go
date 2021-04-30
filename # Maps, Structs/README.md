# Maps, structs

map 은 key 값과  value 값을 통해 데이터를 나타내는데 즉, key 값의 데이터는 value 이다.

```
package main

import "fmt"

func main() {
	DM := map[string]string{"name": "DM", "age": "12"} // int 12 value 값을 string 으로 지정했기 떄문에 받을수 없다
	fmt.Println(DM)
}
```

위에서 key 값 name 의 변수 데이터는 DM, key age 의 변수 데이터는 12(str) 이다.


그렇다면 map 을 range 를 통해서 출력해보자.

```
package main

import "fmt"

func main() {
	DM := map[string]string{"name": "DM", "age": "12"}
	for key, velue := range DM {
		fmt.Println(key, velue)
	}
}

결과값 :
name DM
age 12
```

# structs 

만약 map 에서 데이터타입을 혼용해서 쓰고 싶을때 사용 가능한 것이 "structs" 이다.

```
package main

import "fmt"

type ME struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "mandu"}
	DM := ME{"DM", 21, favFood}
	fmt.Println(DM)
}

```

위와 같이 함수 밖에서 struct 이름과 key와 value 타입을 정의해주고 함수에서 변수값을 입력함으로 struct 를 이용 할 수있다.

그러나 위처람 순서에의해 변수를 정의하는것은 코드상으로 보기 좋지 않다 어떤 변수가 어떤 key 값인지 알기 힘들다.

```
package main

import "fmt"

type ME struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "mandu"}
	DM := ME{name: "DM", favFood: favFood, age: 21}
	fmt.Println(DM)
}

```
위와 같이 key과 같이 정의해주면 훨씬 직접적이게 key 값과 value 값을 확인할 수 있다.

## 키워드 

- map : key 와 value 를 선언하고 입력을 통해 key 와 velue 값을 연결시키는것. 

- struct : map 을 사용하는데 있어 데이터 타입이 한정적이지만 struct 은 데이터 타입을 혼용해서 선언할 수 있다.


## 참고 
[노마드 코더스](https://nomadcoders.co/go-for-beginners/lobby)