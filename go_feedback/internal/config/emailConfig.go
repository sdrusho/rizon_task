package config

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

type DriverWelcomeEmailData struct {
	DriverName       string
	FleetName        string
	PasswordSetupUrl string
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

func (e *EmailService) GeneratePasswordSetupEmailBody(setupURL string) (string, error) {
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

func (e *EmailService) GenerateDriverWelcomeEmailBody(data DriverWelcomeEmailData) (string, error) {
	tmpl, err := template.ParseFiles("pkg/templates/driver-welcome-email.html")
	if err != nil {
		return "", fmt.Errorf("failed to parse welcome email template: %v", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute welcome email template: %v", err)
	}

	return body.String(), nil
}

func (e *EmailService) SendPasswordSetupEmail(email, passwordSetupBaseURL, token string) error {
	setupURL := fmt.Sprintf("%s/%s?token=%s", passwordSetupBaseURL, email, token)

	body, err := e.GeneratePasswordSetupEmailBody(setupURL)
	if err != nil {
		log.Printf("Failed to generate email body: %v", err)
		return err
	}

	err = e.SendEmail(email, "Set up your password", body)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	return nil
}

func (e *EmailService) SendDriverWelcomeEmail(email string, data DriverWelcomeEmailData) error {
	body, err := e.GenerateDriverWelcomeEmailBody(data)
	if err != nil {
		log.Printf("Failed to generate welcome email body: %v", err)
		return err
	}

	err = e.SendEmail(email, "Welcome to Rooya's PolyDrive App!", body)
	if err != nil {
		log.Printf("Failed to send welcome email: %v", err)
		return err
	}

	return nil
}
