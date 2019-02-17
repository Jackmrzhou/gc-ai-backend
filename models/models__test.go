package models

import (
	"fmt"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"os"
	"testing"
)

type pair struct {
	email string
	passwd string
}

var users = []pair{
	{"jack@email.com", "123456"},
	{"bob@email.com", "123456"},
}

var profiles = []Profile{
	{
		Nickname:"test nickname",
		Avatar:"test avatar uri",
		Introduction:"test introduction",
	},
	{
		Nickname:"test nickname 2",
		Avatar:"test avatar uri 2",
		Introduction:"test introduction 2",
	},
}

var games = []Game{
	{
		Name:"game 1",
		Introduction:"Introduction 1",
	},
	{
		Name:"game 2",
		Introduction:"Introduction 2",
	},
}

var usersCreated []User

func TestCreateUser(t *testing.T) {
	var err error
	var _user User
	for i, user := range users{
		_user, err = CreateUser(user.email, user.passwd)
		if err != nil{
			t.Error(err)
		}
		fmt.Println(_user)
		usersCreated = append(usersCreated, _user)
		profiles[i].UserID = _user.ID
	}
	_user, err = CreateUser(users[0].email, users[0].passwd)
	if err == nil{
		t.Error("Create duplicate records should throw an error!")
	}
}

func TestCreateProfile(t *testing.T) {
	var err error
	for _, profile := range profiles{
		err = CreateProfile(&profile)
		if err != nil{
			t.Error(err)
		}
		db.Debug().Model(&profile).Related(&profile.User)
		fmt.Printf("%s\n%s\n",profile, profile.User)
	}
	var fakeProfile = profiles[0]
	fakeProfile.UserID = 1
	err = CreateProfile(&fakeProfile)
	if err == nil{
		t.Error("A profile must have a corresponding user!")
	}
}

func clear() {
	for _, user := range usersCreated{
		var profile Profile
		db.Model(&user).Related(&profile)
		fmt.Println(profile)
		db.Unscoped().Delete(&user)
	}
	for _, profile := range profiles{
		db.Unscoped().Delete(&profile)
	}
	for _, game := range games{
		db.Unscoped().Delete(&game)
	}
}

func TestMain(m *testing.M) {
	err := conf.LoadConf("../conf/app.conf")
	if err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	err = OpenDB()
	if err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}

	m.Run()

	clear()
}