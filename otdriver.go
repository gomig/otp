package otp

import (
	"net/url"

	"github.com/dgryski/dgoogauth"
	"rsc.io/qr"
)

type googleOTP struct {
	issuer   string
	username string
	secret   string
	driver   *dgoogauth.OTPConfig
}

func (otp *googleOTP) Validate(code string) (bool, error) {
	return otp.driver.Authenticate(code)
}

func (otp *googleOTP) RAW() (string, error) {
	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		return "", err
	}
	URL.Path += "/" + url.PathEscape(otp.username)
	params := url.Values{}
	params.Add("secret", otp.secret)
	params.Add("issuer", otp.issuer)
	params.Add("digits", "6")
	URL.RawQuery = params.Encode()
	return URL.String(), nil
}

func (otp *googleOTP) QR() ([]byte, error) {
	if raw, err := otp.RAW(); err != nil {
		return nil, err
	} else if code, err := qr.Encode(raw, qr.Q); err != nil {
		return nil, err
	} else {
		return code.PNG(), nil
	}
}
