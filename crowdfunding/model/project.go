package model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	User      User `gorm:"ForeignKEY:Uid"`
	Uid       uint `gorm:"not null"`
	Title     string
	Content   string `gorm:"type:longtext"`
	Fund      int64
	IsPass    string //unknown//pass//fail
	PicPath   string
	Telephone string
}
