package gomail

import (
	"bytes"
	"text/template"

	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
	"gopkg.in/gomail.v2"
)

type Gomail struct {
	message  *gomail.Message
	dialer   *gomail.Dialer
	htmlPath string
}

func NewGomail(env env.Env) *Gomail {
	return &Gomail{
		message:  gomail.NewMessage(),
		dialer:   gomail.NewDialer(env.EmailHost, env.EmailPort, env.EmailSender, env.EmailPassword),
		htmlPath: env.HtmlPath,
	}
}

func (g *Gomail) SetSender(sender string) {
	g.message.SetHeader("From", sender)
}

func (g *Gomail) SetReciever(to ...string) {
	g.message.SetHeader("To", to...)
}

func (g *Gomail) SetSubject(subject string) {
	g.message.SetHeader("Subject", subject)
}

func (g *Gomail) SetBodyHTML(path string, data interface{}) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(g.htmlPath + path)
	if err != nil {
		return errx.New().WithCode(iris.StatusInternalServerError).WithMessage("Failed to parse template").WithError(err)
	}

	err = t.Execute(&body, data)
	if err != nil {
		return errx.New().WithCode(iris.StatusInternalServerError).WithMessage("Failed to execute template").WithError(err)
	}

	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) Send() error {
	if err := g.dialer.DialAndSend(g.message); err != nil {
		return errx.New().WithCode(iris.StatusInternalServerError).WithMessage("Failed to send email").WithError(err)
	}
	return nil
}
