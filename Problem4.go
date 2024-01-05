package main

import "fmt"

// Implementation a: mathematical formula
// Time complexity: O(1)
// Space complexity: O(1)
func sum_to_n_a (n int) int {
	return n * (n + 1) / 2
}

// Implementation b: iterative
// Time complexity: O(n)
// Space complexity: O(1)
func sum_to_n_b (n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// Implementation c: recursive
// Time complexity: O(n)
// Space complexity: O(n)
func sum_to_n_c (n int) int {
	if n == 0 {
		return 0
	}
	return n + sum_to_n_c(n - 1)
}

// Test the functions
func main() {
    n := 10
    fmt.Println("Sum to", n, "(Method A):", sum_to_n_a(n))
    fmt.Println("Sum to", n, "(Method B):", sum_to_n_b(n))
    fmt.Println("Sum to", n, "(Method C):", sum_to_n_c(n))
}