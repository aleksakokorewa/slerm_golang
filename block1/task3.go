package main

import (
	"fmt"
	"os"
)

func main() {
	str1 := []string{"apple", "banana", "cherry"}
	str2 := []string{"apple", "banana", "Cherry"}
	var str3 []string
	fmt.Println(isSorted(str1))
	fmt.Println(isSorted(str2))
	fmt.Println(isSorted(str3))
}

func isSorted(str []string) bool {
	if len(str) < 1 {
		fmt.Println("Empty string")
		os.Exit(1)
	}
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] > str[j] {
				return false
			}
		}
	}
	return true
}
