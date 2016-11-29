package main

import "fmt"

func square(a int) (int) {
	return a * a
}

func sample(a int) (string, int) {
	return "hello", a * a
}

func demo() (b, c string) {
	b = "first "
	c = "exercise"
	return
}

func main() {
	fmt.Println(square(5))
	fmt.Println(sample(10))
	fmt.Println(demo())
}
