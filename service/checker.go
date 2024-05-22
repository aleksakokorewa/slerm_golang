package service

import "fmt"

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

type Checker struct {
	items []Checkable
}

func (c Checker) String() string {
	var ids []string
	for _, item := range c.items {
		ids = append(ids, item.GetID())
	}
	return fmt.Sprintf("ID's: %v", ids)
}

func (c *Checker) Add(item Checkable) {
	c.items = append(c.items, item)
}

func (c Checker) Check() {
	for _, item := range c.items {
		if !item.Health() {
			fmt.Printf("идентификатор %s элемента не работает: ", item.GetID())
		}
	}
}
