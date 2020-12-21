# 반복문 

## for

go언어는 일반적인 언어들의 for문 구조와 실행 동작을 처리한다. 
for문의 형태는 "for 초기값; 조건식; 증감 { ... }" 의 형식을 따른다.
수행과정은 초기 값을 먼저 조건식과 비교하여 반복을 진행할지 체크하고 증감 처리를 통해 값을 증가시킨다.

go언어에서는 다음과 같이 for문을 사용할 수 있다.

- 일반적인 구조

```go
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
```

- 조건식만 사용하는 구조

```go
n := 1
for n < 100 {
	n *= 2
	//if n > 90 {
	//   break
	//}
}
println(n)
```

- 조건없이 무한 루프 처리

```go
for {
	println("Infinite Loop")
}
```

array, slice, string, map 등 자료구조를 for문에서 처리할 때는 다음과 같이 사용한다.

```go
for key, value := range oldMap {
    newMap[key] = value
}
```


## 참조

- [Go 변수와 상수](http://golang.site/go/article/4-Go-%EB%B3%80%EC%88%98%EC%99%80-%EC%83%81%EC%88%98)
- [디스커버리 Go언어 - 한빛미디어](https://www.hanbit.co.kr/store/books/look.php?p_code=B5279497767)
- [effective_go - for](https://golang.org/doc/effective_go.html#for)
