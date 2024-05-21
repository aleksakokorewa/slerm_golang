package main

import "fmt"

func main() {
	str := "съешь ещё этих мягких французских булок, да выпей чаю"
	str1 := "старт"
	res := countOfLetters(str)
	res1 := countOfLetters(str1)
	printResult(res)
	fmt.Println("")
	printResult(res1)
}

func printResult(m map[rune]int) {
	for k, v := range m {
		fmt.Printf("%c:%d\n", k, v)
	}
}

func countOfLetters(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}
