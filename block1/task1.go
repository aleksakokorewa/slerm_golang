package main

import "fmt"

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

func main() {
	fmt.Println("Тут будет выведен идентификатор")
	check := GenerateCheck()
	for _, item := range check {
		if item.status == PassStatus {
			fmt.Println(item.Id)
		}
	}
}

type HealthCheck struct {
	Id     int
	status string
}

func GenerateCheck() []HealthCheck {
	check := make([]HealthCheck, 6)
	for i := 0; i < 6; i++ {
		check[i].Id = i + 1
		if (i+1)%2 == 0 {
			check[i].status = PassStatus
		} else {
			check[i].status = FailStatus
		}
	}
	return check
}
