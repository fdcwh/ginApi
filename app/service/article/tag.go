package article

import (
	"goGIn/app/models"
)

type Tag struct {
	ID    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State)
}
