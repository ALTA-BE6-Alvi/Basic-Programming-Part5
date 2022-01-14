package main

import (
	"fmt"
)

func Caesar(offset int, input string) string {
	// your code here
}

func main() {
	fmt.Println(Caesar(3, "abc"))                         //def
	fmt.Println(Caesar(2, "alta"))                        // cnvc
	fmt.Println(Caesar(10, "alterraacademy"))             //kvdobbkkmknowi
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))  //bcdefghijklmnopqrstuvwxyza
	fmt.Print(Caesar(1000, "abcdefghijklmnopqrstuvwxyz")) //mnopqrstuvwxyzabcdefghijkl
}
