package main

type iOtp interface {
	generateRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publicMetric()
}

type otp struct {
	iOtp iOtp
}

func (o *otp) generateAndSendOTP(otpLength int) error {
	otp := o.iOtp.generateRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publicMetric()
	return nil
}

func (o *otp) setIotp(iOtp iOtp) {
	o.iOtp = iOtp
}
