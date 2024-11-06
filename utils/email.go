package utils

import (
	"GinVueA/define"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendEmail 发送邮件
func SendEmail(to string, subject string, text string) {
	e := email.NewEmail()
	// 设置发送方的邮箱
	e.From = define.EmailFrom
	// 设置接收方的邮箱
	e.To = []string{to}
	// 设置标题
	e.Subject = subject
	// 设置邮件发送的内容
	e.Text = []byte(text)
	// 设置服务器相关的配置
	emailAddr := define.EmailHost + ":" + define.EmailPort
	err := e.Send(emailAddr, smtp.PlainAuth("", define.EmailFrom, define.EmailPassWord, define.EmailHost))
	if err != nil {
		fmt.Println(err)
	}
}
