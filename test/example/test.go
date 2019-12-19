package example

import (
	"fmt"
	"net/http"
)

type Request struct {
	From string
}

func CreateRequest(req *http.Request) *Request {
	return &Request{}
}

type EmailSender interface {
	Send(to, subject, body string) error
}

type SendEmail struct {
	From string
	Request *http.Request
}

func (sender *SendEmail) Send(to, subject, body string) error {
	return nil
}

func NewSendEmail() (*SendEmail, error) {
	return &SendEmail{}, nil
}

type CustomerWelcome struct {
	Emailer EmailSender
}

func NewCustomerWelcome(sender EmailSender) *CustomerWelcome {
	return &CustomerWelcome{
		Emailer: sender,
	}
}

func (welcomer *CustomerWelcome) Welcome(name, email string) error {
	body := fmt.Sprintf("Hi, %s!", name)
	subject := "Welcome"

	return welcomer.Emailer.Send(email, subject, body)
}
