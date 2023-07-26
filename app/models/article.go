package models

import "goGIn/app/models/init"

type Article struct {
	init.ID
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	State         int    `json:"state"`
	init.Timestamps
}
