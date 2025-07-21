package main

import "fmt"

func main() {

	arr := [2][2]int{{2, 13}, {5, 79}}
	// arr.length --> len(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
}
