package main

import "errors"

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(data string) error
}

// this does not implement the Handler interface, but is used for composition for other instances of Handler
type handler struct {
	next Handler
}

func (h *handler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}
func (h *handler) NextHandle(data string) error {
	if h.next != nil {
		return h.next.Handle(data)
	}
	return nil
}

// To check if empty data
type emptyHandler struct {
	handler
}

func (e *emptyHandler) Handle(data string) error {
	if data == "" {
		return errors.New("Empty")
	}
	return e.NextHandle(data)
}

func NewEmptyHandler() Handler {
	return &emptyHandler{}
}

// To validate length of the data
type LengthHandler struct {
	handler
}

func (e *LengthHandler) Handle(data string) error {
	if len(data) < 5 {
		return errors.New("Short")
	}
	return e.NextHandle(data)
}

func NewLengthHandler() Handler {
	return &LengthHandler{}
}

// To save the data in the repo after validation
type RepoHandler struct {
	handler
	repo Repository
}

func (e *RepoHandler) Handle(data string) error {
	e.repo.PutData(data)
	return nil
}

func NewRepositoryHandler(repo Repository) Handler {
	return &RepoHandler{repo: repo}
}
