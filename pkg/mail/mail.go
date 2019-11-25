package mail

import (
	"gin-project-template/pkg/util"
	"github.com/matcornic/hermes"
	"gopkg.in/gomail.v2"
)

func Generate(email hermes.Email) (body string) {
	appConf, err := util.FetchAppConf()
	if err != nil {
		panic(err)
	}
	mailProduct := appConf.MailProduct
	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      mailProduct.Name,
			Link:      mailProduct.Link,
			Logo:      mailProduct.Logo,
			Copyright: mailProduct.Copyright,
		},
	}
	h.Theme = new(hermes.Default)
	body, err = h.GenerateHTML(email)
	if err != nil {
		panic(err)
	}
	return body
}

func Send(subject string, body string, to ...string) {
	appConf, err := util.FetchAppConf()
	if err != nil {
		panic(err)
	}
	mailConf := appConf.Mail
	m := gomail.NewMessage()
	m.SetHeader("From", mailConf.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mailConf.Host, mailConf.Port, mailConf.From, mailConf.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
