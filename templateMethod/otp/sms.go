package main

import "fmt"

type sms struct {}

func (s *sms) generateRandomOTP(len int) string {
	randomOTP := generateNumber(len)
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}
func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP is " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *sms) publicMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}

