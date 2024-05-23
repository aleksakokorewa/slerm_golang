package main

import (
	"time"
)

type GoMetrClient struct {
	URL     string
	Timeout time.Duration
	Check   HealthCheck
}

// конструктор для гмклиент
func NewGoMetrClient(url string, timeout time.Duration) *GoMetrClient {
	return &GoMetrClient{
		URL:     url,
		Timeout: timeout,
		Check:   HealthCheck{}, //сюда можно потом добавить логику для реализации
	}
}

type HealthCheck struct {
	ServiceId string
	status    string
}

func (g *GoMetrClient) GetMetrics() string {
	return g.URL
}

func (g *GoMetrClient) Ping() error {
	return nil
}

func (g *GoMetrClient) GetID() string {
	return g.URL
}

func (g *GoMetrClient) Health() bool {
	return false
}

func (g *GoMetrClient) getHealth() HealthCheck {
	return HealthCheck{
		ServiceId: g.GetID(),
		status:    "pass",
	}
}
