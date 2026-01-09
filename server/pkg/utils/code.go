package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateVerificationCode membuat 6 digit code untuk password reset
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // 100000-999999
	return fmt.Sprintf("%06d", code)
}

// Generate4DigitCode membuat 4 digit code untuk account verification
func Generate4DigitCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9000) + 1000 // 1000-9999
	return fmt.Sprintf("%04d", code)
}

// Generate6DigitCode alias untuk GenerateVerificationCode
func Generate6DigitCode() string {
	return GenerateVerificationCode()
}
