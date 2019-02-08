package models

import (
	"crypto/sha256"
	"fmt"
	"github.com/jackmrzhou/gc-ai/conf"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	//Id int `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	Email string `gorm:"type:varchar(320); UNIQUE_INDEX"`
	Password string `gorm:"type:char(64)"`
	IsBanned int `gorm:"type:tinyint"`
}

func (self User) String() string {
	return fmt.Sprintf("<id: %d, email: %s, is_banned: %d>", self.ID, self.Email, self.IsBanned)
}

type Profile struct {
	gorm.Model
	//Id int `gorm:"PRIMARY_KEY"`
	User User
	UserID int
	// referencing back
	Nickname string `gorm:"type:varchar(20)"`
	Avatar string `gorm:"type:varchar(200)"`
	Introduction string `gorm:"type:varchar(200)"`
}

func (self Profile) String() string {
	return fmt.Sprintf("<user_id: %d, nickname: %s, avatar: %s, introduction: %s>",
			self.UserID,
			self.Nickname,
			self.Avatar,
			self.Introduction,
		)
}

type VerificationCode struct {
	ID uint `gorm:"PRIMARY_KEY"`
	Email string `gorm:"type:varchar(320); UNIQUE_INDEX"`
	Code string `gorm:"type:char(6)"`
	Timestamp time.Time
}

func (self VerificationCode) String() string {
	return fmt.Sprintf("<email: %s, code: %s, timestamp: %s>",
		self.Email,
		self.Code,
		self.Timestamp)
}

func CreateUser(email, passwd string) (User, error) {
	passwdHashed := fmt.Sprintf("%x", sha256.Sum256([]byte(passwd)))
	user := User{
		Email:email,
		Password:passwdHashed,
		IsBanned:0,
	}
	err := db.Create(&user).Error
	return user, err
}

func CreateProfile(profile *Profile) error {
	if err := db.Where("id = ?", profile.UserID).First(&User{}).Error; err!=nil{
		return err
	}
	err := db.Create(profile).Error
	return err
}

func QueryUser(email, passwd string) (User, error){
	passwdHashed := fmt.Sprintf("%x", sha256.Sum256([]byte(passwd)))
	var user User
	if err := db.Where("email = ? AND password = ?", email, passwdHashed).First(&user).Error; err !=nil{
		return user, err
	}
	return user, nil
}

func IsBanned(user *User) bool {
	return user.IsBanned == 1
}

func DeleteUser(email string) {
	if conf.RunMode == "debug" {
		db.Unscoped().Delete(User{}, "email = ?", email)
	}
}

func CreateVeriCode(email, code string) error {
	err := db.Create(&VerificationCode{
		Email:email,
		Code:code,
		Timestamp:time.Now(),
	}).Error
	return err
}

func QueryVeriCode(email string) (VerificationCode,error) {
	var code VerificationCode
	if err := db.Where("email = ?", email).First(&code).Error; err != nil{
		return code, err
	}
	return code, nil
}

func DeleteVeriCode(code *VerificationCode) {
	db.Delete(code)
}

func DeleteVeriCodeDebug(email string) {
	if conf.RunMode == "debug"{
		db.Delete(VerificationCode{}, "email = ?", email)
	}
}

func CreateVeriCodeDebug(code *VerificationCode) {
	if conf.RunMode == "debug"{
		db.Create(code)
	}
}