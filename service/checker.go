package main

import "fmt"

type Checker struct {
	items []Checkable
}

func NewChecker() *Checker {
	return &Checker{
		items: make([]Checkable, 0),
	}
}

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

func (c *Checker) Add(items ...Checkable) {
	c.items = append(c.items, items...)
}

func (c Checker) String() string {
	var ids []string
	for _, item := range c.items {
		ids = append(ids, item.GetID())
	}
	return fmt.Sprintf("ID's: %v\n", ids)
}

func (c Checker) Check() {
	for _, item := range c.items {
		if !item.Health() {
			fmt.Printf("идентификатор элемента %s\n не работает\n", item.GetID())
		}
	}
}
