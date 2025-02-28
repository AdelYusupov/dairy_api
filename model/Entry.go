package model

import (
	"diary_api/database"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserId  uint
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return entry, nil
	}

	return entry, nil
}
func (entry *Entry) Delete() (*Entry, error) {
	err := database.Database.Delete(&entry).Error
	if err != nil {
		return entry, nil
	}
	return entry, nil
}
