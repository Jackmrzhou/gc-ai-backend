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
	IsAdmin	int `gorm:"type:tinyint"`
	Ranks []Rank `gorm:"PRELOAD:false"`
	SourceCodes []SourceCode `gorm:"PRELOAD:false"`
	Attacks []Battle `gorm:"foreignkey:AttackerID; PRELOAD:false"`
	Defenses []Battle `gorm:"foreignkey:DefenderID; PRELOAD:false"`
}

func (self User) String() string {
	return fmt.Sprintf("<id: %d, email: %s, is_banned: %d>", self.ID, self.Email, self.IsBanned)
}

type Profile struct {
	gorm.Model `json:"-"`
	//Id int `gorm:"PRIMARY_KEY"`
	User User `json:"-"`
	UserID uint `json:"-"`
	// referencing back
	Nickname string `gorm:"type:varchar(20); UNIQUE_INDEX" json:"nickname"`
	Avatar string `gorm:"type:varchar(200)" json:"avatar"`
	Introduction string `gorm:"type:varchar(200)" json:"introduction"`
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
	SourceCodes []SourceCode `gorm:"PRELOAD:false" json:"-"`
	BattleHistory []Battle `gorm:"PRELOAD:false" json:"-"`
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

type SourceCode struct {
	gorm.Model `json:"-"`
	UserID uint `json:"user_id"`
	GameID uint	`json:"game_id"`
	CodeType int `gorm:"type:tinyint" json:"code_type"`
	Language string `gorm:"type:varchar(20)" json:"language"`
	Content string `gorm:"type:TEXT" json:"source_code"`
}

// two const for CodeType in SourceCode
const (
	DEFAULT = 0
	ATTACK = 1
	DEFEND = 2
)

func (self *SourceCode) MarshalJSON() ([]byte, error) {
	type Alias SourceCode
	return json.Marshal(&struct {
		ID uint `json:"id"`
		*Alias
	}{
		self.ID,
		(*Alias)(self),
	})
}

func (self SourceCode) String() string {
	return fmt.Sprintf("SourceCode<UserID: %d, GameID: %d>", self.UserID, self.GameID)
}

type Rank struct {
	gorm.Model
	GameID uint
	UserID uint
	Score  uint
}

func (self Rank) String() string {
	return fmt.Sprintf("Rank<UserID: %d, GameID: %d>", self.UserID, self.GameID)
}

type Battle struct {
	// only support 1V1
	gorm.Model `json:"-"`
	Status uint `json:"status"`
	AttackerID uint `json:"attacker_id"`
	DefenderID uint `json:"defender_id"`
	GameID uint `json:"game_id"`
	Detail string `gorm:"type:TEXT" json:"detail"`
	WinnerID uint `json:"winner_id"`
	RewardScore uint `json:"reward_score"`
	PenaltyScore uint `json:"penalty_score"`
}

func CreateBattle(atkID, defID uint, game *Game) (*Battle, error) {

	// todo:check atkID defID game

	battle := Battle{
		Status:Suspending,
		AttackerID:atkID,
		DefenderID:defID,
		GameID:game.ID,
	}

	err := db.Create(&battle).Error
	return &battle,err
}

func UpdateBattle(battle *Battle) error {
	err := db.Save(battle).Error
	return err
}

func QueryBattleByID(id uint) (*Battle, error) {
	battle := new(Battle)
	err := db.Where("id = ?", id).First(battle).Error
	return battle, err
}

func UpdateBattleStatusByID(id uint, status uint) error {
	var battle Battle
	err := db.Where("id = ?", id).First(&battle).Error
	if err != nil{
		return err
	}

	err = db.Model(&battle).Update("status", status).Error
	return err
}

const (
	Suspending = 1
	Judeging = 2
	Finished = 3
)

func (self Battle) String() string {
	return fmt.Sprintf("Battle<ATKID: %d, DEFID: %d>", self.AttackerID, self.DefenderID)
}

func IsEngagedInBattle(userID uint) bool{
	// check whether a user has battles not finished
	var battles []Battle
	db.Where("attacker_id = ? AND status <> ?", userID, Finished).Find(&battles)
	return len(battles) != 0
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

func QueryUserByID(ID uint) (*User, error) {
	user := new(User)
	err := db.Where("id = ?", ID).First(user).Error
	return user, err
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
	//log.Println("Starting cleaning verification codes...")
	var codes []VerificationCode
	db.Find(&codes)
	for _, code := range codes{
		if code.Timestamp.Add(conf.CodeActiveTime).Before(time.Now()){
			// this record is not active
			db.Delete(&code)
		}
	}
	//log.Println("Cleaning verification codes finished.")
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
	// todo:check user, game
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

func CreateRankByID(userID, gameID, score uint) (*Rank, error) {
	rank := new(Rank)
	// todo:check user, game
	if err:= db.Where("user_id = ? AND game_id = ? AND score = ?",
		userID, gameID, score).First(rank).Error; err == nil{
		return rank, fmt.Errorf("UserID: %d, GameID: %d. Score already exists.", userID, gameID)
	}
	rank.UserID = userID
	rank.GameID = gameID
	rank.Score = score
	err := db.Create(rank).Error
	return rank, err
}

func DeleteRankByIDDebug(ID uint) {
	if conf.RunMode == "debug" {
		db.Unscoped().Delete(&Rank{}, "ID = ?", ID)
	}
}

func QueryRankByGameID(gameID uint) ([]Rank, error) {
	var ranks []Rank
	err := db.Where("game_id = ?", gameID).Find(&ranks).Error
	return ranks, err
}

func QueryRankByUserID(userID uint) ([]Rank, error) {
	var ranks []Rank
	err := db.Where("user_id = ?", userID).Find(&ranks).Error
	return ranks, err
}

func QueryRankByUserAndGameID(userID, gameID uint) (*Rank, error) {
	rank := new(Rank)
	err := db.Where("user_id = ? AND game_id = ?", userID, gameID).First(rank).Error
	return rank, err
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

func UpdateRankByID(id uint, scoreDelta int) error {
	rank := new(Rank)
	if err := db.Where("id = ?", id).First(rank).Error; err != nil{
		return err
	}
	sc := int(rank.Score)
	sc += scoreDelta
	if sc < 0{
		// prevent overflow
		rank.Score = 0
	}else{
		rank.Score = uint(sc)
	}
	err := db.Model(rank).Update("score", rank.Score).Error
	return err
}

func CreateSourceCode(user *User, game *Game, codeType int, language,content string) (*SourceCode, error) {
	sourceCode := new(SourceCode)
	// todo:check user, game
	sourceCode.GameID = game.ID
	sourceCode.UserID = user.ID
	sourceCode.CodeType = codeType
	sourceCode.Language = language
	sourceCode.Content = content
	err := db.Create(sourceCode).Error
	return sourceCode, err
}

func SetSrcRole(userID, srcID uint, role int) error {
	var sourceCodes []SourceCode
	db.Where("user_id = ?", userID).Find(&sourceCodes)
	for _, src := range sourceCodes{
		if src.ID == srcID{
			src.CodeType = role
			db.Save(&src)
		}else if src.CodeType == role{
			src.CodeType = DEFAULT
			db.Save(&src)
		}
	}
	return nil
}

func QuerySourceCodeByUserID(userID uint) ([]SourceCode, error) {
	var codes []SourceCode
	err := db.Where("user_id = ?", userID).Find(&codes).Error
	return codes, err
}

func QuerySourceCodeByUserGameIDs(userID, gameID uint) ([]SourceCode, error) {
	var codes []SourceCode
	err := db.Where("user_id = ? AND game_id = ?", userID, gameID).Find(&codes).Error
	return codes, err
}

func QueryATKSrcByUserID(userID uint) (*SourceCode, error) {
	src := new(SourceCode)
	err := db.Where("user_id = ? AND code_type = ?", userID, ATTACK).First(src).Error
	return src, err
}

func QueryDEFSrcByUserID(userID uint) (*SourceCode, error) {
	src := new(SourceCode)
	err := db.Where("user_id = ? AND code_type = ?", userID, DEFEND).First(src).Error
	return src, err
}

func UpdateSourceCode(code *SourceCode) error {
	err := db.Save(code).Error
	return err
}

func QueryProfileByUserID(userID uint) (*Profile, error) {
	profile := new(Profile)
	err := db.Where("user_id = ?", userID).First(profile).Error
	return profile, err
}

func UpdateProfile(profile *Profile) error {
	return db.Save(profile).Error
}

func QueryProfileByNickname(nickname string) (*Profile, error) {
	profile := new(Profile)
	err := db.Where("nickname = ?", nickname).First(profile).Error
	return profile, err
}

func QueryBattlesByUserID(userID uint) ([]Battle, error) {
	var battles []Battle
	err := db.Where("attacker_id = ? OR defender_id = ?", userID, userID).Find(&battles).Error
	return battles, err
}