package main

import (
	"fmt"

	hr "github.com/robmurgai/learning_Golang/hackerrank"
)

func main() {

	fmt.Printf("\n########## Starting Exercise ##########\n\n")

	// Simple Test Cases.
	// arr := []int64{1, 2, 2, 4}
	// r := int64(2)

	// {1,3,9,9,27,81} with r=3
	// arr := []int64{1, 3, 9, 9, 27, 81}
	// r := int64(3)

	// n and r, the size of arr and the common ratio.
	// n space-seperated integers arr[i].

	// Test Contraint i < j < k
	//{1, 2, 1, 2, 4} with r=2, the answer is 3 because indeces i < j < k
	// arr := []int64{1, 2, 1, 2, 4}
	// r := int64(2)

	// Test Constraint 1 <= n <= 10^5 where n is the size of arr
	// arr := make([]int64, 1e5)
	// for i := 0; i < len(arr); i++ {
	// 	arr[i] = int64((i + 1) * 1e4)
	// }
	// r := int64(1e2)

	// Test Constraint 1 <= arr[i] <= 10^9 where arr[i] is the value of the integers in the input array.
	//{1, 10^3, 10^6, 10^9} with r=10^3, the answer is 2 (1, 10^3, 10^6) (10^3, 10^6, 10^9)
	// arr := []int64{1, 1e3, 1e6, 1e9}
	// r := int64(1e3)

	// Test Constraint 1 <= r <= 10^9 where r is the common ratio.
	// r=10^9 use case
	//{1, 10^3, 10^6, 10^9} with r=10^9, the answer is 0 because at 10^9, you can only have 1 number at that size so you can't form a tripe, only a double
	// arr := []int64{1, 1e3, 1e6, 1e9}
	// r := int64(1e9)

	// Test Constraint 1 <= r <= 10^9 where r is the common ratio.
	// r=1 use case.
	// arr := []int64{5, 5, 5}
	// r := int64(1)

	// [2, 2, 2, 2] = 4 (0, 1, 2) (0, 1, 3) (0, 2, 3) (1, 2, 3) = 3^0 + 3*1
	// arr := []int64{2, 2, 2, 2}
	// r := int64(1)

	// [2, 2, 2, 2, 2] = 10 (0, 1, 2) (0, 1, 3) (0, 2, 3) (1, 2, 3) (0, 1, 4) (0, 2, 4) (0, 3, 4) (1, 2, 4) (1, 3, 4) (2, 3, 4) =  = 3^0 + 3*1 +3*2 = 1 + 3((len-3)!)
	// [2, 2, 2, 2, 2], Answer is 10
	// index 2; Len: 3;	(0, 1, 2)  = 1
	// index 3; Len: 4;	(0, 1, 3) (0, 2, 3)      (1, 2, 3) = 2 + 1
	// index 4; Len: 5;	(0, 1, 4) (0, 2, 4) (0, 3, 4)       (1, 2, 4) (1, 3, 4)      (2, 3, 4) =  3 + 2 + 1
	// arr := []int64{2, 2, 2, 2, 2}
	// r := int64(1)

	// [2, 2, 2, 2, 2]
	// Len: 3;	(0, 1, 2)  = 1
	// Len: 4;	(0, 1, 3) (0, 2, 3)      (1, 2, 3) = 2 + 1
	// Len: 5;	(0, 1, 4) (0, 2, 4) (0, 3, 4)       (1, 2, 4) (1, 3, 4)      (2, 3, 4) =  3 + 2 + 1
	// Len: 6;	(0, 1, 5) (0, 2, 5) (0, 3, 5) (0, 4, 5)         (1, 2, 5) (1, 3, 5) (1, 4, 5)        (2, 3, 5) (2, 4, 5)     (3, 4, 5) = 4 + 3 + 2 + 1

	// Test Case #3, Answer is 166661666700000
	arr := make([]int64, 1e5)
	for i := 0; i < len(arr); i++ {
		arr[i] = int64(1234)
	}
	r := int64(1)

	fmt.Printf("\n\nNum of triplets: %v\n", hr.CountTriplets(arr, r))

	fmt.Printf("\n########## Ending Exercise ##########\n")
}
