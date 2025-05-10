package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
	EmailMessage string
}

type SMSNotifier struct {
	SMSMessage string
}

func (e EmailNotifier) Send(msg string) {
	fmt.Println("Email Message :", msg)
}

func (s SMSNotifier) Send(msg string) {
	fmt.Println("SMS Message :", msg)
}

func NotifyUser(n Notifier, msg string) {
	n.Send(msg)
}

func main() {
	email := EmailNotifier{
		EmailMessage: "Hello World !!!",
	}
	NotifyUser(email, email.EmailMessage)
	sms := SMSNotifier{
		SMSMessage: "1234567890",
	}
	NotifyUser(sms, sms.SMSMessage)
}
