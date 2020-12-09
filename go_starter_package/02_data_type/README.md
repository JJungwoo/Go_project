# 데이터 타입

## Go 데이터 타입

Go 프로그래밍 언어는 다음과 같은 기본적인 데이터 타입들을 갖고 있다.

- 1. bool 타입
- 2. 문자열 타입
  - string: 문자열은 한번 생성되면 수정될 수 없는 Immutable 타입이다.
- 3. 정수형 타입
  - int, int8, int16, int32, int64
  - uint, uint8, uint16, uint32, uint64, uinptr
- 4. Float 및 복소수 타입
  - float32, float64, complex64, complex128
- 5. 기타 타입
  - byte: uint8과 동일하며 바이트 코드에 사용
  - rune: int32과 동일하며 유니코드 코드포인터에 사용

## 문자열

문자열 리터럴은 Back Quote('') 혹은 이중인용부호("")를 사용하여 표현할 수 있다.

- 1. Back Quote('')로 둘러 싸인 문자열은 Raw String Literal 이라 부르는데, 이 안에 있는 문자열은 별도로 해석되지 않고 Raw String 그대로의 값을 갖는다. 예를 들어, 문자열 안에 `\n` 이 있을 경우 이는 New Line 즉 개행으로 해석되지 않는다. Back Quote는 복수 라인의 문자열을 표현할 때 자주 사용된다.
- 2. 이중인용부호("")로 둘러 싸인 문자열은 Interpreted String Literal이라 부르는데, 복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석된다. 위에서 `\n` 는 이중인용부호의 경우에는 개행으로 처리된다. 이중인용부호를 이용하여 문자열을 여러 라인에 걸쳐쓰기 위해서는 `+` 연산자를 이용해 결합하여 사용한다.

```go
	rawLiteral := `테스트\n
테스트 입니다.\n
끝입니다.`

    interLiteral1 := "테스트\n테스트 입니다."
    interLiteral2 := "테스트\n" +
					 "테스트 입니다."

    fmt.Println("raw literal : " + rawLiteral)
    fmt.Println()
    fmt.Println("interpreted1 : " + interLiteral1)
    fmt.Println("interpreted2 : " + interLiteral2)
```

- 출력 화면

```
raw literal : 테스트\n
테스트 입니다.\n
끝입니다.

interpreted1 : 테스트
테스트 입니다.
interpreted2 : 테스트
테스트 입니다.
```

## 데이터 타입 변환

하나의 데이터 타입에서 다른 데이터 타입으로 변환하기 위해서는 T(v)와 같이 표현하고 이를 Type Conversion이라 부르는데, 여기서 T는 변환하고자 하는 타입을 표시하고, v는 변환될 값(value)을 지정한 것이다. 

```go
var i int = 100
var u uint = uint(i)
var f float32 = float32(i)

str := "ABC"
bytes := []byte(str)
str2 := string(bytes)
```

> Go에서는 타입간 변환이 명시적으로 지정해야 한다. 간단한 int에서 uint로 변환할 때, 암묵적 변환이 일어나지 않고 런타임 에러가 발생한다.


## 참조

- [Go 변수와 상수](http://golang.site/go/article/4-Go-%EB%B3%80%EC%88%98%EC%99%80-%EC%83%81%EC%88%98)
- [디스커버리 Go언어 - 한빛미디어](https://www.hanbit.co.kr/store/books/look.php?p_code=B5279497767)
