package view

import "app/models"

type Tag struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}

func NewTag(tag *models.Tag) *Tag {
	t := &Tag{
		ID:  tag.ID,
		Tag: tag.Tag,
	}

	return t
}
