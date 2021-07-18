package controller

import "go.sample.com/interface/repository"

type Controller struct {
	repository repository.Repository
}

type TestRepo struct {
	repository.Repository
}

func (t *TestRepo) Start(out string) (string, error) {
	return out, nil
}

func (c *Controller) Run(out string) (string, error) {
	c.repository = new(TestRepo)
	s, err := c.repository.Start(out)
	return s, err
}
