package models

import (
	"goGIn/app/models/init"
	"goGIn/kernel"
)

type Tag struct {
	init.ID
	Name  string `json:"name"`
	State int    `json:"state"`
	init.Timestamps
}

func AddTag(name string, state int) error {
	tag := Tag{
		Name:  name,
		State: state,
	}
	if err := kernel.FdGorm.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}
