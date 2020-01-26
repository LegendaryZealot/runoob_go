package main

import "fmt"

func main() {
	var str string = "string"
	fmt.Println(str)
	var a, b int = 12, 1
	fmt.Println(a, b)
	fmt.Println(&a, &b)
	autoType := "auto type"
	fmt.Println(autoType)
}
