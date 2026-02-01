package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// pg 建表语句
type Article struct {
	gorm.Model
	Title   string         `gorm:"type:varchar(255);not null;unique"`
	Content string         `gorm:"type:text;not null"`
	Tags    pq.StringArray `gorm:"type:varchar(255)[]"`
}
