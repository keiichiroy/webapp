package models

import (
	"time"
)

type Member struct {
	Id         uint64    `db: "id" json:"id"`
	Username   string    `db: "username" json:"username"`
	Password   string    `db: "password" json:"password"`
	Last_name  string    `db: "last_name" json: "last_name"`
	First_name string    `db: "first_name" json:"first_name"`
	Birthday   string    `db: "birthday" json:"birthday"`
	Ken        int       `db: "ken" json: "ken"`
	Reg_date   time.Time `db: "reg_date" json: "reg_date"`
	Cancel     time.Time `db: "cancel" json: "cancel"`
}
