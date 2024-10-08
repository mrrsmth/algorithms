package main

import (
	f "algorithms/first"
	sw "algorithms/slidingWindow"
	tp "algorithms/twoPointers"

	"fmt"
)

func main() {
	fmt.Println("Hello algorithm!")

	// First
	fmt.Println(f.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(f.Fib(5))
	fmt.Println(f.PlusOne([]int{4, 3, 2, 1}))

	//Two Pointers
	reverseStr := []byte{'h', 'e', 'l', 'l', 'o'}
	tp.ReverseString(reverseStr)
	fmt.Printf("Reverse string: %s\n", string(reverseStr))
	fmt.Println(tp.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(tp.TwoSum([]int{1, 2, 4}, 6))
	fmt.Println(tp.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))

	// Sliding Window
	fmt.Println(sw.FindMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
