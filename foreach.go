package main

import "fmt"

func foreach_one() {

	myList := []string{"dog", "cat", "hedgehog"}

	// for {key}, {value} := range {list}
	for _, animal := range myList {
		fmt.Println("My animal is:", animal)
	}
}

func foreach_two() {
	
	myList := map[string]string{
		"minsu":	"27",
		"lion":		"43",
		"kimmin":	"30",
	}

	for name, age := range myList {
		fmt.Println(name, "age is", age)
	}
}

func main() {


	//foreach_one()
	foreach_two()


}
