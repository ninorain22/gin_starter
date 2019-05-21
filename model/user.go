package model

import "time"

type User struct {
	Id int64	`json:"id"`
	Name string	`xorm:"varchar(50)" json:"name"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}
