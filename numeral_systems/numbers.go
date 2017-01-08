package main

import "fmt"

func main() {
	fmt.Println(42)
	fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)

	for i := 1000; i < 1010; i++ {
		fmt.Printf("%d \t %b \t %x", i, i, i)
	}
}
