package main

import "fmt"

func main() {
    // Raw String Literal. 복수라인.
    rawLiteral := `테스트\n
테스트 입니다.\n
끝입니다.`

    // Interpreted String Literal
    interLiteral1 := "테스트\n테스트 입니다."
    // 아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
    interLiteral2 := "테스트\n" +
					 "테스트 입니다."

    fmt.Println("raw literal : " + rawLiteral)
    fmt.Println()
    fmt.Println("interpreted1 : " + interLiteral1)
    fmt.Println("interpreted2 : " + interLiteral2)
}
