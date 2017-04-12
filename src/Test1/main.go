// Test1 project main.go
package main

import (
	"fmt"

	"./pack1"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is :%f\n", ms.f1)
	fmt.Printf("The string is :%s\n", ms.str)
	fmt.Println(ms)
}

func main22() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
}
