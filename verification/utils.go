package verification

import (
	"fmt"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"math/rand"
	"time"
)

type codePair struct {
	veriCode string
	timestamp time.Time
}

var source = rand.NewSource(time.Now().Unix())

var mailSender CodeMailSender

func SendCode(email string) (string, error) {
	code := generateCode(email)
	var err error

	vericode, err := models.QueryVeriCode(email)
	if err == nil{
		// already exists
		// todo:write into config
		if vericode.Timestamp.Add(time.Minute).Before(time.Now()) {
			models.DeleteVeriCode(&vericode)
		}else{
			return "", fmt.Errorf("Sending email too fast. %s", email)
		}
	}

	err = models.CreateVeriCode(email, code)
	if err != nil{
		// storing into database failed
		return "", err
	}
	err = mailSender.SendMail(email,code)
	if err != nil{
		// send mail failed
		return "", err
	}
	return code, nil
}

func CheckAndDelCode(email, code string) bool {
	c,err := models.QueryVeriCode(email)
	if err == nil && c.Code == code && c.Timestamp.Add(conf.CodeActiveTime).After(time.Now()){
		// valid
		models.DeleteVeriCode(&c)
		return true
	}else if err == nil && c.Code == code{
		// time expired
		models.DeleteVeriCode(&c)
	}
	return false
}

func generateCode(email string) string{
	r := rand.New(source)
	num := r.Intn(900000) + 100000
	// map to [10 0000, 100 0000)
	return fmt.Sprint(num)
}