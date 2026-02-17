package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	dialer *gomail.Dialer
	from   string
}

func NewEmailService(host string, port int, username, password, from string) *EmailService {
	dialer := gomail.NewDialer(host, port, username, password)
	return &EmailService{
		dialer: dialer,
		from:   from,
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := e.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %v", err)
	}

	return nil
}

func (e *EmailService) GenerateSignupEmailBody(setupURL string) (string, error) {
	tmpl, err := template.ParseFiles("pkg/templates/user_setup_email.html")
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, struct{ SetupURL string }{SetupURL: setupURL})
	if err != nil {
		return "", err
	}

	return body.String(), nil
}

func (e *EmailService) SendSignupEmail(email string) error {
	setupURL := fmt.Sprint("exp://192.168.110.11:8081/--/enjoying", email)

	body, err := e.GenerateSignupEmailBody(setupURL)
	if err != nil {
		log.Printf("Failed to generate email body: %v", err)
		return err
	}

	err = e.SendEmail(email, "Set up your feedback user", body)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	return nil
}
