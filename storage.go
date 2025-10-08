package main

import (
	"encoding/json"
	"os"
)

type Stogare[T any] struct {
	FileName string
}

func NewStogare[T any](fileName string) *Stogare[T] {
	return &Stogare[T]{
		FileName: fileName,
	}
}

func (s *Stogare[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(s.FileName, fileData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *Stogare[T]) Load(data *T) error {

	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, data)
	if err != nil {
		return err
	}

	return nil
}
