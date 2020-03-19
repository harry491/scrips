package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/**
邮件发送
*/
func SendEmailTo(to string, code string) {
	e := email.NewEmail()
	e.From = "zhangnan <zhangnan910404@sina.cn>"
	e.To = []string{to}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Scrips验证码"
	e.Text = []byte("你的验证码为:")
	e.HTML = []byte("<h1>" + code + "</h1>")
	err := e.Send("smtp.sina.cn:25", smtp.PlainAuth("", "zhangnan910404@sina.cn", "zhangnan110", "smtp.sina.cn"))
	if err != nil {
		fmt.Println(err)
	}
}
