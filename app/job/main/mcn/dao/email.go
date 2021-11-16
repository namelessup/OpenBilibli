package dao

import (
	"os"

	"github.com/namelessup/bilibili/app/job/main/mcn/conf"
	"github.com/namelessup/bilibili/library/log"

	"gopkg.in/gomail.v2"
)

// SendMail send the email.
func (d *Dao) SendMail(body string, subject string, send []string) (err error) {
	log.Info("send mail send:%v", send)
	msg := gomail.NewMessage()
	msg.SetHeader("From", conf.Conf.MailConf.Username)
	msg.SetHeader("To", send...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body, gomail.SetPartEncoding(gomail.Base64))
	if err = d.email.DialAndSend(msg); err != nil {
		log.Error("s.email.DialAndSend error(%v)", err)
		return
	}
	return
}

// SendMailAttach send the email.
func (d *Dao) SendMailAttach(filename string, subject string, send []string) (err error) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", conf.Conf.MailConf.Username)
	msg.SetHeader("To", send...)
	msg.SetHeader("Subject", subject)
	msg.Attach(filename)
	if err = d.email.DialAndSend(msg); err != nil {
		log.Error("s.email.DialAndSend error(%v)", err)
		return
	}
	err = os.Remove(filename)
	return
}
