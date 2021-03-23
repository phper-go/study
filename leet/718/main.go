package main

import (
	"fmt"
)

func main() {

	// var arr1 = []int{1, 2, 3, 2, 1, 3, 2, 1, 2, 1, 2, 1, 2, 2, 2, 4, 7, 3, 2}
	// var arr2 = []int{3, 2, 1, 4, 7, 3, 2}

	var arr1 = []int{1, 1, 0, 0, 1}
	var arr2 = []int{1, 1, 0, 0, 0, 2}

	// var arr1 = []int{0, 0, 0, 0, 0}
	// var arr2 = []int{0, 0, 0, 0, 0}

	// var arr1 = []int{1, 2}
	// var arr2 = []int{3, 2}

	fmt.Println(findLength(arr1, arr2))
}

func findLength(A []int, B []int) int {
	var mapB = make(map[int][]int)
	for k, v := range B {
		mapB[v] = append(mapB[v], k)
	}
	var arrMax []int
	var lenA = len(A)
	for i := 0; i < lenA; i++ {
		for _, key := range mapB[A[i]] {
			var max []int
			for t := 0; t < lenA-i; t++ {
				var Anv = A[i+t]
				var Bnk = mapB[Anv]
				if !in_array(key+t, Bnk) {
					break
				}
				max = B[key : key+t+1]
			}

			if len(max) > len(arrMax) {
				arrMax = max
			}
		}
	}
	fmt.Println(arrMax)
	return len(arrMax)
}

func in_array(v int, arr []int) bool {
	for _, val := range arr {
		if v == val {
			return true
		}
	}
	return false
}
