package db

import (
	"fmt"
	"log"

	"github.com/ShawnRong/bento/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
)

var db *gorm.DB

type dbconfig struct {
	host   string
	port   int
	dbname string
	user   string
	pass   string
}

func Init() {
	var err error

	c := config.GetConfig()
	s := &dbconfig{
		host:   c.GetString("db.host"),
		port:   c.GetInt("db.port"),
		dbname: c.GetString("db.database"),
		user:   c.GetString("db.username"),
		pass:   c.GetString("db.password"),
	}
	connectInfo := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", s.user, s.pass, s.host, s.port, s.dbname)
	fmt.Println(connectInfo)

	db, err = gorm.Open("mysql", connectInfo)
	db.LogMode(c.GetBool("site.debug"))
	if err != nil {
		log.Fatal("error on database connection, info: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	defer db.Close()
}
