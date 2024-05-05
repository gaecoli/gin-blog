package utils

import (
	g "gin-blog/internal/global"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func SendEmail(from string, to, cc, bcc []string, subject string, message []byte, emailType string) error {
	emailSender := email.NewEmail()

	emailSender.From = from
	emailSender.To = to
	//emailSender.Cc = cc
	//emailSender.Bcc = bcc

	emailSender.Subject = subject

	switch emailType {
	case "html":
		emailSender.HTML = message
	default:
		emailSender.Text = message
	}

	err := emailSender.Send(g.GetConfig().Email.EmailHost, smtp.PlainAuth("", g.GetConfig().Email.Username, g.GetConfig().Email.Password, g.GetConfig().Email.Host))

	if err != nil {
		log.Printf("发送邮件异常，异常信息：%s", err)
	}

	return err
}
