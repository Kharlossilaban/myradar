package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// OTP Types for different purposes
const (
	OTPTypeRegistration  = "REG"
	OTPTypePasswordReset = "PWD"
)

// GenerateVerificationCode generates a 6-digit code for password reset
// Returns: "PWD-XXXXXX" format
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // 100000-999999
	return fmt.Sprintf("%s-%06d", OTPTypePasswordReset, code)
}

// GenerateRegistrationCode generates a 6-digit code for email verification
// Returns: "REG-XXXXXX" format
func GenerateRegistrationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // 100000-999999
	return fmt.Sprintf("%s-%06d", OTPTypeRegistration, code)
}

// Generate6DigitCode alias for GenerateVerificationCode (backward compatibility)
func Generate6DigitCode() string {
	return GenerateVerificationCode()
}

// Generate4DigitCode - DEPRECATED: Use GenerateRegistrationCode instead
// Kept for backward compatibility during transition
func Generate4DigitCode() string {
	return GenerateRegistrationCode()
}

// ExtractOTPDigits extracts the 6-digit code from formatted OTP (e.g., "REG-123456" -> "123456")
func ExtractOTPDigits(formattedOTP string) string {
	if len(formattedOTP) >= 10 && formattedOTP[3] == '-' {
		return formattedOTP[4:]
	}
	return formattedOTP // Return as-is if not in expected format
}

// IsValidOTPFormat validates OTP format (with or without prefix)
func IsValidOTPFormat(code string, allowPrefix bool) bool {
	if allowPrefix {
		// Format: XXX-XXXXXX (10 chars)
		if len(code) == 10 && code[3] == '-' {
			prefix := code[:3]
			if prefix != OTPTypeRegistration && prefix != OTPTypePasswordReset {
				return false
			}
			digits := code[4:]
			for _, c := range digits {
				if c < '0' || c > '9' {
					return false
				}
			}
			return true
		}
	}

	// Format: XXXXXX (6 digits only)
	if len(code) == 6 {
		for _, c := range code {
			if c < '0' || c > '9' {
				return false
			}
		}
		return true
	}

	return false
}

// GetOTPType extracts the type from formatted OTP
func GetOTPType(formattedOTP string) string {
	if len(formattedOTP) >= 4 && formattedOTP[3] == '-' {
		return formattedOTP[:3]
	}
	return ""
}
