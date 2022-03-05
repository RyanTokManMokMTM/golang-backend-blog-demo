//Package mail - send email
package mail

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

//Email using SMTP Protocol
type Email struct {
	*SMTP
}

//SMTP define SMTP/Email server info
type SMTP struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTP) *Email {
	return &Email{
		SMTP: info,
	}
}

/*
SendMail
@param to : to a group of email address
@param subject : email subject
@param body :content of the email
*/
func (e *Email) SendMail(to []string, subject, body string) error {
	msg := gomail.NewMessage() //new email message and set required info including from address,to address etc
	//email header format
	msg.SetHeader("From", e.From)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	//Connect to SMTP server by given info
	dia := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dia.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dia.DialAndSend(msg) //open connection and send the message
}
