# 조건문

## 조건문

### if 문

Go의 조건문은 `괄호()` 로 둘러 싸지 않아도 된다. if 문 내부 코드는 조건이 맞으면 `블럭{}` 안에 내용을 실행한다.
그리고 반드시 조건과 블럭 시작 `{` 을 같은 라인에 두어야 한다. 같은 라인에 두지 않으면 에러가 발생한다.
<br>

Go의 조건문은 **반드시 Boolean 식**으로 표현되어야 한다.

- if 예시

```go
if k == 1 {		// 	같은 라인
	println("true")
}
```

- if - else if - else 예시

```go
if k == 1 {
	println("if")
} else if k == 2 {	// 같은 라인
	println("else if")
} else {			// 같은 라인
	println("else")
}
``` 	

- if 조건에 변수 선언

Go의 if 문에서는 조건식을 사용하기 이전에 간단한 코드를 실행할 수 있다. 
대신 if 조건문에서 사용한 변수는 해당 블럭 내부에서만 사용이 가능하다. 
조건문 뿐만 아니라 for문, switch문  또한 이러한 방식이 가능하다.

```go
tmp := 2
if val := tmp * 2; val < max {
	println(var)
}

var++ // if 블럭을 벗어나 사용하면 에러 발생
```

### switch 문

여러 값을 비교해야 하는 경우 다수의 조건식 방법으로 switch문이 있다. 
Go의 switch문은 break를 따로 하지 않아도 자동으로 case를 종료한다. 
일반적인 switch처럼 default를 사용할 수 있다.

- switch 예시

```go
var num = 1
switch num {
case 1:
	println("1")
case 2:
	println("2")
case 3, 4:
	println("3, 4")
default:
	println("default")
}
```

- switch 조건에 변수 선언

```go
var num = 1
switch x:= num << 2; x - 1 {
	// ...
}
```

- 추가 사용 예시

```go
switch v.(type) {
case int:
	println("int")
case bool:
	println("bool")
case string:
	println("string")
default:
	println("unknown")
}
```

만약 case를 종료하지 않으려면 해당 case 끝에 아래와 같이 `fallthrough`를 추가해주면 된다.

```go
val := 1
switch val {
case 1:
	println("1 이하")
	fallthrough
case 2:
	println("2 이하")
	fallthrough
case 3:
	println("3 이하")
	fallthrough
default:
	println("default")
}
```

## 참조

- [Go 변수와 상수](http://golang.site/go/article/4-Go-%EB%B3%80%EC%88%98%EC%99%80-%EC%83%81%EC%88%98)
- [디스커버리 Go언어 - 한빛미디어](https://www.hanbit.co.kr/store/books/look.php?p_code=B5279497767)
