package main

import "sort"

func AfinderEfficient(arr []int, x int) [][]int {
	sort.Ints(arr)
	result := [][]int{}
	for i, ei := range arr { //[1,2,3]
		for _, ej := range arr[(i + 1):] { //[2,3]
			if ei+ej == x {
				result = append(result, []int{ei, ej})
			}
		}
	}
	return result
}

func AfinderBeckup(arr []int, x int) [][]int {
	result := [][]int{}
	for i, ei := range arr { //[1,2,3]
		for _, ej := range arr[(i + 1):] { //[2,3]
			if ei+ej == x {
				result = append(result, []int{ei, ej})
			}
		}
	}
	return result
}
