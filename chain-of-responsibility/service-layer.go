package main

import "errors"

type Service interface {
	ValidateAndSet(string) error
	GetData() string
}

type service struct {
	repo Repository
}

// handle the responsibility
func (s *service) ValidateAndSet(data string) error {
	if data != "" {
		s.repo.PutData(data)
		return nil
	}

	return errors.New("Empty data")
}

// repo is responsible to send the required data to client
func (s *service) GetData() string {
	return s.repo.GetData()
}

// creating a chain of responsibility
func NewService(repository Repository) Service {
	return &service{repo: repository}
}
