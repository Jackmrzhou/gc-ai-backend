package verification

import (
	"fmt"
	"github.com/jackmrzhou/gc-ai/conf"
	"gopkg.in/gomail.v2"
)

type MailSender interface {
	SendMail(email, body string) error
}

type CodeMailSender struct {

}

func (mailSender *CodeMailSender) SendMail(email, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.MailAddress)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "AI Battle 注册验证码")
	m.SetBody("text/html", fmt.Sprintf("%s, 10分钟内有效。", code))

	d := gomail.NewDialer("smtp.qq.com", 587, conf.MailAddress, conf.MailAuth)

	err := d.DialAndSend(m)
	return err
}