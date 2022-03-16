package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/plancer16/mall_end/config"
	"github.com/plancer16/mall_end/handler"
	"github.com/plancer16/mall_end/model"
	"github.com/plancer16/mall_end/repository"
	"github.com/plancer16/mall_end/service"
	"github.com/spf13/viper"
	"log"
)

var (
	DB *gorm.DB
	UserHandler handler.UserHandler
)

func init() {
	initViper()
	initDB()
	initHandler()
}

func initViper() {
	err := config.Init("")
	if err != nil {
		panic(err)
	}
}

func initDB() {
	conf := model.DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}
	var err error
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		"Local")
	DB, err = gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("connection err:%s",err)
	}
	DB.SingularTable(true)
	fmt.Println("DB初始化完成")
}

func initHandler() {
	UserHandler = handler.UserHandler{UserService: service.UserService{UserRepo: repository.UserRepo{DB: DB}}}
}