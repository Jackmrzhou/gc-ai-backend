package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	cfg *ini.File

	RunMode string

	//server
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JWTSecret string

	// database
	DBType   string
	DBUser   string
	DBPasswd string
	DBHost   string
	DBName   string

	// mail
	MailAddress string
	MailAuth	string
)

func LoadConf(path string) error{
	var err error
	cfg, err = ini.Load(path)
	if err != nil{
		return err
	}
	err = loadApp()
	err = loadServer()
	err = loadDatabase()
	err = loadMail()
	return err
}

func loadApp() error{
	RunMode = cfg.Section("App").Key("RUN_MODE").MustString("debug")
	JWTSecret = cfg.Section("App").Key("JWT_SECRET").MustString("googlecamp")
	return nil
}

func loadServer() error{
	sec, err := cfg.GetSection("server")
	if err != nil{
		log.Fatal(2, "Load server config error")
		return err
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = sec.Key("READ_TIMEOUT").MustDuration(time.Duration(60*time.Second))
	WriteTimeout = sec.Key("WRITE_TIMEOUT").MustDuration(time.Duration(60*time.Second))
	return nil
}

func loadDatabase() error{
	sec, err := cfg.GetSection("database")
	if err != nil{
		log.Fatal(2, "Load database error")
		return err
	}
	DBType = sec.Key("TYPE").MustString("mysql")
	DBUser = sec.Key("USER").String()
	DBPasswd = sec.Key("PASSWORD").String()
	DBHost = sec.Key("HOST").String()
	DBName = sec.Key("DB").String()
	return nil
}

func loadMail() error{
	sec, err := cfg.GetSection("mail")
	if err != nil{
		log.Fatal("Load mail config failed")
		return err
	}
	MailAddress = sec.Key("MAIL").String()
	MailAuth = sec.Key("AUTH").String()
	return nil
}