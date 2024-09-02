package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerID uint
	ToID    uint
	Type    int
	Description string
}

func (c *Contact) TableName() string {
	return "contacts"
}