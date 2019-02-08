package verification

import (
	"fmt"
	"github.com/jackmrzhou/gc-ai/models"
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
	if err == nil && c.Code == code && c.Timestamp.Add(10 * time.Minute).After(time.Now()){
		models.DeleteVeriCode(&c)
		return true
	}else if err == nil{
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