package main

import "fmt"

type Notify interface {
	Execute() string
}

type Code struct {
	code string
}

func (c *Code) Execute() string {
	return c.code
}

type EmailNotification struct {
	Notify
}

func (en *EmailNotification) Execute() string {
	fmt.Println("Sending email...")
	return fmt.Sprintf("Your verification is %s\n", en.Notify.Execute())
}

type SmsNotification struct {
	Notify
}

func (es *SmsNotification) Execute() string {
	fmt.Println("Sending sms...")
	return fmt.Sprintf("Your verification is %s\n", es.Notify.Execute())
}

func main() {
	newCode := &Code{"123123"}
	sms := &SmsNotification{newCode}
	fmt.Println(sms.Execute())
}