package main

import "fmt"

func main() {
	alphabets := [2][3]string{{"a","b","c"},{"A","B","c"}}
	fmt.Println(alphabets)
	fmt.Println(alphabets[0][1])
	