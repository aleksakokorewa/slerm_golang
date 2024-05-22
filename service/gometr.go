package service

import "errors"

type GoMetrClient struct {
	URL     string
	Timeout int
}

type HealthCheck struct {
	Id     string
	Status string
}

func (c *GoMetrClient) GetHealth() HealthCheck {
	return HealthCheck{
		Id: c.URL,
	}
}

func (g *GoMetrClient) Ping() error {
	return errors.New("ping not implemented")
}

func (g *GoMetrClient) GetID() string {
	return g.ID
}

func (g *GoMetrClient) Health() bool {
	return true
}
