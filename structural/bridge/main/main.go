package main

import (
	"fmt"
	// "strconv"
	"time"
)

func main() {
	fmt.Println("do it")
	var a string
	a = time.Now().Format(time.RFC3339)
	fmt.Println(a)

	// var a string

	// aa := time.Now().Unix()
	// a = strconv.FormatInt(aa, 16)
	// fmt.Println(aa)
	// fmt.Println(a)
}
