package firebase

var otpStore = make(map[string]string)

func SaveOTP(email, otp string) {
	otpStore[email] = otp
}

func GetOTP(email string) string {
	return otpStore[email]
}

func DeleteOTP(email string) {
	delete(otpStore, email)
}
