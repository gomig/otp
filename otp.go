package otp

import (
	"crypto/md5"
	"encoding/base32"
	"strings"

	"github.com/dgryski/dgoogauth"
)

type OTP interface {
	// Validate check if otp code is valid
	Validate(code string) (bool, error)
	// RAW get raw setup key
	RAW() (string, error)
	// QR get setup key QR Code image
	QR() ([]byte, error)
}

func NewGoogleOTP(issuer, username string, keys ...string) OTP {
	otp := new(googleOTP)
	otp.issuer = issuer
	otp.username = username

	// Generate Secret
	secret := md5.Sum([]byte(issuer + username + strings.Join(keys, "")))
	_len := len(secret)
	_mid := _len / 2
	otp.secret = base32.StdEncoding.EncodeToString([]byte{secret[0], secret[1], secret[2], secret[_mid-1], secret[_mid], secret[_mid+1], secret[_mid+2], secret[_len-3], secret[_len-2], secret[_len-1]})

	// generate driver
	otp.driver = &dgoogauth.OTPConfig{
		Secret:      otp.secret,
		WindowSize:  3,
		HotpCounter: 0,
		UTC:         true,
	}
	return otp
}
