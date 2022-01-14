package main

import "fmt"

func Compare(a, b string) string {
	// your code here
	var hasil string

	if len(a) > len(b) {
		i := 0
		for i < (len(a) - len(b) + 1) {
			if b == a[i:len(b)] {
				hasil = b
				i = len(a)
			}
			i++
		}
	} else {
		i := 0
		for i < (len(b) - len(a) + 1) {
			if a == b[i:len(a)] {
				hasil = a
				i = len(b)
			}
			i++
		}
	}

	return hasil
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
