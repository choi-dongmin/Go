# arrays and slices

go 에서 배열을 선언하고자 한다면 아래와 같은 방식으로 할 수있다.

```
package main

import "fmt"

func main() {
	names := [5]string{"DM", "BR", "BM", "GN", "YS"}
	fmt.Println(names)
}
```

names 배열의 길이,데이터타입 그리고 index 들을 선언함으로 작성할 수 있고

```
package main

import "fmt"

func main() {
	names := [5]string{"DM", "BR", "BM"}
	names[3] = "GN"
	names[4] = "YS"
	fmt.Println(names)
}
```

위와같은 형식으로 부족한 배열의 index 를 추가 시킬수 있다. 하지만 이 경우에는 index의 번호가 5 이상이면 선언한 배열의 길이보다 더 길기 때문에 오류가 발생한다.

## slice

slices 하는 법을 통해 우리는 배열의 길이가 없는 배열 를 만들수 있다

```
package main

import "fmt"

func main() {
	names := []string{"DM", "BR", "BM"}
	names = append(names, "GN")
	fmt.Println(names)
}
```

우선 배열을 선언하는데 있어 길이를 지정하지 않는다 그후 append(배열이름, 추가할 index값) 을 통해 기존의 index 값들과 추가할 index 를 합쳐 새로운 names를 return 한다.

## 키워드

slice : index 의 제한이 없는 배열을 만드는것

append : slice 에 item 을 추가하는 함수


## 참고
[노마드 코더스](https://nomadcoders.co/go-for-beginners/lectures/1509)