package main

import "fmt"

func main() {
	smsOTP := &sms{}
	o := otp{iOtp: smsOTP}

	err := o.generateAndSendOTP(6)
	if err != nil {
		panic(err)
	}

	fmt.Println("")

	emailOTP := &email{}
	o.setIotp(emailOTP)
	err = o.generateAndSendOTP(6)
	if err != nil {
		panic(err)
	}
}