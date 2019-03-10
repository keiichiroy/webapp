package lib

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Sess *gorm.DB

func init() {
	var err error
	dbinfo := os.Getenv("GODBINFO") //この形式で　→　host= port=5432 user= dbname= password= sslmode=require
	Sess, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
}
