package main

import (
	"fmt"
	// "reflect"
	// "strconv"
)

func Caesar(offset int, input string) string {
	// your code here

	var output string
	
	// constrain offset
	if offset > (122 - 97) {
		offset = offset % (122 - 96)
	}

	number := 0

	for i := 0; i < len(input); i++ {
		if input[i] >= 97 && input[i] <= 122 {
			number = int(input[i]) + offset
			if number < 97 {
				number = 122 - (97 - number)
			} else if number > 122 {
				number = 97 + (number - 123)
			}
			output += fmt.Sprint(string(rune(number)))
		} else {
			output += string(input[i])
		}
	}

	return output
}

func main() {
	fmt.Println(Caesar(3, "abc"))                         //def
	fmt.Println(Caesar(2, "alta"))                        // cnvc
	fmt.Println(Caesar(10, "alterraacademy"))             //kvdobbkkmknowi
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))  //bcdefghijklmnopqrstuvwxyza
	fmt.Print(Caesar(1000, "abcdefghijklmnopqrstuvwxyz")) //mnopqrstuvwxyzabcdefghijkl
}
