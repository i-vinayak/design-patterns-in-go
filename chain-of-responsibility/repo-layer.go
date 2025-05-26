package main

type Repository interface {
	GetData() string
	PutData(string)
}

type repository struct {
	data string
}

func (r *repository) GetData() string {
	return r.data
}

func (r *repository) PutData(data string) {
	r.data = data
}

func NewRepository() Repository {
	return &repository{}
}
