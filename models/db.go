package models

import (
	"fmt"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func OpenDB() error {
	if db != nil{
		return fmt.Errorf("Database connection has already opened!")
	}
	var err error
	db, err = gorm.Open(conf.DBType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBUser,
		conf.DBPasswd,
		conf.DBHost,
		conf.DBName))
	if err != nil{
		return err
	}
	initialRun()
	return nil
}

func initialRun() {
	if !db.HasTable(&User{}){
		db.CreateTable(&User{})
	}
	if !db.HasTable(&Profile{}){
		db.CreateTable(&Profile{})
	}
	if !db.HasTable(&VerificationCode{}){
		db.CreateTable(&VerificationCode{})
	}
	if !db.HasTable(&Game{}){
		db.CreateTable(&Game{})
	}
	if !db.HasTable(&Rank{}){
		db.CreateTable(&Rank{})
	}
	if !db.HasTable(&SourceCode{}){
		db.CreateTable(&SourceCode{})
	}
	if !db.HasTable(&Battle{}){
		db.CreateTable(&Battle{})
	}
}

func CloseDB() {
	defer db.Close()
}