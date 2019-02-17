package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	//Id int `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	Email string `gorm:"type:varchar(320); UNIQUE_INDEX"`
	Password string `gorm:"type:char(64)"`
	IsBanned int `gorm:"type:tinyint"`
	Ranks []Rank `gorm:"PRELOAD:false"`
}

func (self User) String() string {
	return fmt.Sprintf("<id: %d, email: %s, is_banned: %d>", self.ID, self.Email, self.IsBanned)
}

type Profile struct {
	gorm.Model
	//Id int `gorm:"PRIMARY_KEY"`
	User User
	UserID uint
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

type Game struct {
	gorm.Model `json:"-"`
	Name string `gorm:"type:varchar(100); UNIQUE_INDEX" json:"name"`
	Introduction string `gorm:"type:TEXT" json:"introduction"`
	Ranks []Rank `gorm:"PRELOAD:false" json:"-"`
}

func (self *Game) MarshalJSON() ([]byte, error) {
	type Alias Game
	return json.Marshal(&struct {
		ID uint `json:"id"`
		*Alias
	}{
		self.ID,
		(*Alias)(self),
	})
}

func (self Game) String() string {
	return fmt.Sprintf("<Name: %s, Introduction: %s>", self.Name, self.Introduction)
}

type Rank struct {
	gorm.Model
	GameID uint
	UserID uint
	Score  uint
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

func QueryProfile(userID uint) (*Profile, error) {
	profile := new (Profile)
	err := db.Where("user_id = ?", userID).First(profile).Error
	return profile, err
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

func DeleteUserDebug(email string) {
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

func MaintainVeriCode() {
	var codes []VerificationCode
	db.Find(&codes)
	for _, code := range codes{
		if code.Timestamp.Add(conf.CodeActiveTime).Before(time.Now()){
			// this record is not active
			db.Delete(&code)
		}
	}
}

func CreateGame(game *Game) error {
	err := db.Create(game).Error
	return err
}

func DeleteGameByIDDebug(ID uint) {
	if conf.RunMode == "debug" {
		db.Unscoped().Delete(&Game{}, "ID = ?", ID)
	}
}

func DeleteGameByNameDebug(name string) {
	if conf.RunMode == "debug"{
		db.Unscoped().Delete(&Game{}, "Name = ?", name)
	}
}

func QueryAllGames() []Game {
	var games []Game
	db.Find(&games)
	return games
}

func QueryGameByID(id uint) (*Game, error) {
	game := new(Game)
	err := db.Where("id = ?", id).First(game).Error
	return game, err
}

func CreateRank(user *User, game *Game, score uint) (*Rank, error) {
	rank := new(Rank)
	if err:= db.Where("user_id = ? AND game_id = ? AND score = ?",
		user.ID, game.ID, score).First(rank).Error; err == nil{
		return rank, fmt.Errorf("User: %s, Game: %s. Score already exists.", user.Email, game.Name)
	}
	rank.UserID = user.ID
	rank.GameID = game.ID
	rank.Score = score
	err := db.Create(rank).Error
	return rank, err
}

func DeleteRankByIDDebug(ID uint) {
	if conf.RunMode == "debug" {
		db.Unscoped().Delete(&Rank{}, "ID = ?", ID)
	}
}

func QueryRankByGameID(gameID uint) ([]*Rank, error) {
	var ranks []*Rank
	err := db.Where("game_id = ?", gameID).Find(&ranks).Error
	return ranks, err
}

func QueryRankByUserID(userID uint) ([]*Rank, error) {
	var ranks []*Rank
	err := db.Where("user_id = ?", userID).Find(&ranks).Error
	return ranks, err
}

func UpdateRank(user *User, game *Game, score uint) error {
	rank := new(Rank)
	if err := db.Where("user_id = ? AND game_id = ?",
		user.ID, game.ID).First(rank).Error; err != nil{
			return err
	}
	err := db.Model(rank).Update("score", score).Error
	return err
}

func UpdataRankByID(id, score uint) error {
	rank := new(Rank)
	if err := db.Where("id = ?", id).First(rank).Error; err != nil{
		return err
	}
	err := db.Model(rank).Update("score", score).Error
	return err
}