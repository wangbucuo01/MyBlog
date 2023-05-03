package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

// 用于传递发送邮箱所必需的信息
type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	// 创建一个消息实例
	m := gomail.NewMessage()
	// 发件人、收件人、主题、正文
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// NewDialer创建一个新的SMTP拨号实例，设置对应的拨号信息
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	// 打开与SMTP服务器的连接并发送电子邮件
	return dialer.DialAndSend(m)
}
