# OTP

**Google Authenticator** client for golang.

## Usage

```go
package main

import "github.com/gomig/otp"

// Create Instance
driver := otp.NewGoogleOTP("mysite.com", "john@doe.com", "extra key 1", "extra key 2")

// Get Setup Key
if setupKey, err := driver.Raw(); err != nil {
    log.Fatal("OTP key failed!")
} else {
    log.Logf("Enter this setup key in Google Authenticator App: %s\n", setupKey)
}

// Get Setup QR
if setupQr, err := driver.QR(); err != nil {
    log.Fatal("OTP QRCode failed!")
} else {
    // Send Otp base64 image data to user
}

// Validate OTP
if ok, err := driver.Validate("123456"); err != nil {
    log.Fatal("OTP validation failed!")
} else if !ok {
    log.Fatal("OTP invalid!")
} else {
    log.Log("OTP is valid")
}
```
