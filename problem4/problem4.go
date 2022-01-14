package main

import "fmt"

func FindaMaxSumSubArray(k int, arr []int) int {
	// your code here
}

func main() {
	fmt.Println(FindaMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})) //9
	fmt.Println(FindaMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))    //7
	fmt.Println(FindaMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))    //5
	fmt.Println(FindaMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))    //7
	fmt.Println(FindaMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))    //9
}
