package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name        string `json:"name"`
	OwnerID     uint   `json:"owner_id"`
	Icon        string `json:"icon"`
	Type        int	   `json:"type"`
	Description string `json:"description"`
}

func (gb *GroupBasic) TableName() string {
	return "groups_basic"
}