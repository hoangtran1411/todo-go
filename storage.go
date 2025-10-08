package main

type Stogare[T any] struct {
	FileName string
}

func NewStogare[T any](fileName string) *Stogare[T] {
	return &Stogare[T]{
		FileName: fileName,
	}
}
