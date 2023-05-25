package models

import (
	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Skill struct {
	Id          uuid.UUID `json:"id" gorm:"column:sk_id;type:uuid;default:gen_random_uuid();primarykey"`
	Name        string    `json:"name" gorm:"column:sk_name" validate:"nonzero"`
	Description string    `json:"description" gorm:"column:sk_description" validate:"nonzero"`
	Group       string    `json:"group" gorm:"column:sk_group" validate:"nonzero"`
	Category    string    `json:"category" gorm:"column:sk_category" validate:"nonzero"`
	Purpose     string    `json:"purpose" gorm:"column:sk_purpose" validate:"nonzero"`
	Priority    int       `json:"priority" gorm:"column:sk_priority" validate:"nonzero,min=0,max=10"`
	Level       int       `json:"level" gorm:"column:sk_level" validate:"nonzero,min=0,max=10"`
}

func ValidateSkillData(skill *Skill) error {
	if err := validator.Validate(skill); err != nil {
		return err
	}
	return nil
}
