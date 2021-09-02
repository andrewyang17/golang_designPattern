package main

import "fmt"

type email struct {}

func (e *email) generateRandomOTP(len int) string {
	randomOTP := generateNumber(len)
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}
func (e *email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (e *email) getMessage(otp string) string {
	return "EMAIL OTP is " + otp
}

func (e *email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

func (e *email) publicMetric() {
	fmt.Printf("EMAIL: publishing metrics\n")
}

