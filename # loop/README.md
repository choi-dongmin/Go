# loop

## for 

go 에서 반복문은 오직 for 를 통해서 가능하다. 

```

import (
	"fmt"
)
func supperAdd(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i)
	}
}
func main() {
	supperAdd(1, 2, 3, 4, 5, 6)
}

결과값
0
1
2
3
4
5 
```
이렇게 사용할 수 있다 만약 매개변수로 받는 값을 사용하고 싶다면 number 에 [index] 값을 이용할 수 있다.

```
func supperAdd(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
``` 

## range 
- range 는 배열을 반복 시킬때 조금 더 잘 적용할 수 있도록 하는 기능이다.

for each / for in 과 같은 형태

```
import (
	"fmt"
)
func supperAdd(numbers ...int){
	for number := range numbers {
		fmt.Println(number)
	}
}
func main() {
	supperAdd(1, 2, 3, 4, 5, 6)
}

결과값 
0
1
2
3
4
5
```    

for 와 range 를 이용해서 값을 출력하고 있다 그런데 결과값은 1,2,3,4,5,6 이 아닌 0,1,2,3,4,5 이다.

그 이유는 range 는 index 를 사용한다.

```
import (
	"fmt"
)
func supperAdd(numbers ...int) {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
}
func main() {
	supperAdd(1, 2, 3, 4, 5, 6)
}

결과값 
0 1
1 2
2 3
3 4
4 5
5 6

* 우리는 index 값을 사용하고 싶지 않다면 _ 를 통해 무시할 수 있다. 
```
위처럼 range 는 for 안에서만 사용할 수 있다.

그럼 인자값을 모두 더하여 값을 받고 싶다면 어떠한 방식으로 할 수 있을까?

```
func supperAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number //total = total + number
	}
	return total
}
func main() {
	result := supperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

결과값
21
```
위와 같은 방식으로 매개변수로 받는 인자값 모두를 더 할수 있다