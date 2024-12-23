package models

import "gorm.io/gorm"

type DBUser struct {
	gorm.Model
	Name string
}
