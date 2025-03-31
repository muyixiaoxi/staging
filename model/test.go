package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.DB
	Name string
	Time time.Time
}
