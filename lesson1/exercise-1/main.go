package main

import "fmt"

func main() {
	var i int;
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &i)
	fmt.Println("Square of a number:",square(i));
	number, word := demo(i, "Someword")
	fmt.Printf("number: %v, string: %v\n", number, word)
	fmt.Println("Naked Return:",sample(i));


}

func square(num int) int {
	return num * num
}

func demo(num int, word string) (int, string) {
	return num, word
}

func sample(num int) (a int) {
	a = num + num
	return
}

