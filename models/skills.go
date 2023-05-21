package models

import (
	"github.com/google/uuid"
)

type Skill struct {
	Id          uuid.UUID `json:"id" gorm:"column:sk_id;type:uuid;default:gen_random_uuid();primarykey"`
	Name        string    `json:"name" gorm:"column:sk_name"`
	Description string    `json:"description" gorm:"column:sk_description"`
	Group       string    `json:"group" gorm:"column:sk_group"`
	Category    string    `json:"category" gorm:"column:sk_category"`
	Purpose     string    `json:"purpose" gorm:"column:sk_purpose"`
	Priority    int       `json:"priority" gorm:"column:sk_priority"`
	Level       int       `json:"level" gorm:"column:sk_level"`
}
