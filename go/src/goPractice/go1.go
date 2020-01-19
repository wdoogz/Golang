package main

import (
	"fmt"
	"os"
)

func main() {
	testPrint := "Its worth a shot ;)"
	fmt.Println("Heres to Learning Golang! ", testPrint)
	arg := 7
	thirdFunction(secondFunction(), arg)
}

func secondFunction() int {
	testfuncvar := 11
	return testfuncvar
}

func thirdFunction(testfuncvar int, arg string) {
	fmt.Println(testfuncvar, arg)
}
