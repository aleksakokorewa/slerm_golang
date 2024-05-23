package main

import (
	"fmt"
	"time"
)

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

func main() {
	fmt.Println("Тут будет выведен идентификатор")
	client1 := NewGoMetrClient("https://host1.com", 10*time.Second)
	client2 := NewGoMetrClient("https://host2.com", 10*time.Second)

	checker := Checker{}
	checker.Add(client1, client2)
	fmt.Println(checker)
	checker.Check()
}
