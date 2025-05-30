package gomail

import (
	"bytes"
	"html/template"
	"log"

	"github.com/ahargunyllib/thera-be/internal/infra/env"
	"github.com/ahargunyllib/thera-be/internal/infra/mailtmpl"
	"gopkg.in/gomail.v2"
)

type CustomGomailInterface interface {
	SendEmail(to string, subject string, body string) error
	SendEmailTemplate(to string, subject string, template string, data map[string]any) error
	SendEmailsTemplate(to []string, subject string, template string, data map[string]any) error
}

type CustomGomailStruct struct {
	Host        string
	SenderEmail string
	SenderName  string
	Port        int
	Username    string
	Password    string
	Templates   *template.Template
}

var Gomail = getGomail()

func getGomail() CustomGomailInterface {
	// Parse all templates at startup
	templates, err := template.ParseFS(mailtmpl.Templates, "*.html")
	if err != nil {
		log.Fatal(map[string]interface{}{
			"error": err.Error(),
		}, "[MAIL][NewMailDialer] failed to parse templates")
		return nil
	}

	return &CustomGomailStruct{
		Host:        env.AppEnv.GomailHost,
		SenderEmail: env.AppEnv.GomailSenderEmail,
		SenderName:  env.AppEnv.GomailSenderName,
		Port:        env.AppEnv.GomailPort,
		Username:    env.AppEnv.GomailUsername,
		Password:    env.AppEnv.GomailPassword,
		Templates:   templates,
	}
}

func (g *CustomGomailStruct) SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(g.SenderEmail, g.SenderName)},
		"To":      {to},
		"Subject": {subject},
	})
	m.SetBody("text/html", body)

	d := gomail.NewDialer(g.Host, g.Port, g.Username, g.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (g *CustomGomailStruct) SendEmailTemplate(to string, subject string, template string, data map[string]any) error {
	var body bytes.Buffer

	if err := g.Templates.ExecuteTemplate(&body, template, data); err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(g.SenderEmail, g.SenderName)},
		"To":      {to},
		"Subject": {subject},
	})
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(g.Host, g.Port, g.Username, g.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (g *CustomGomailStruct) SendEmailsTemplate(
	to []string,
	subject string,
	template string,
	data map[string]any,
) error {
	var body bytes.Buffer

	if err := g.Templates.ExecuteTemplate(&body, template, data); err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(g.SenderEmail, g.SenderName)},
		"To":      to,
		"Subject": {subject},
	})
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(g.Host, g.Port, g.Username, g.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
