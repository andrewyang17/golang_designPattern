package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMail(email *email) {
	fmt.Printf("Email has successfully sent to %v!\n", email.to)
}

type build func(builder *EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMail(&builder.email)
}

func main() {
	SendEmail(func(eb *EmailBuilder) {
		eb.From("yangzhisiangg@gmail.com").
			To("therock@gmail.com").
			Subject("Meeting").
			Body("Hello do you want to meet?")
	})
}
