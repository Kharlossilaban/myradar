package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EmailVerification stores OTP codes for email verification during registration
type EmailVerification struct {
	ID               string     `gorm:"type:varchar(36);primaryKey" json:"id"`
	UserID           string     `gorm:"type:varchar(36);not null" json:"user_id"`
	Email            string     `gorm:"type:varchar(255);not null;index:idx_ev_email" json:"email"`
	VerificationCode string     `gorm:"type:varchar(10);not null;index:idx_ev_code" json:"verification_code"` // Updated to 10 chars for "REG-XXXXXX"
	ExpiresAt        time.Time  `gorm:"not null;index:idx_ev_expires_at" json:"expires_at"`
	Used             bool       `gorm:"default:false" json:"used"`
	FailedAttempts   int        `gorm:"default:0" json:"failed_attempts"`                  // Track failed verification attempts
	LockedUntil      *time.Time `gorm:"index:idx_ev_locked" json:"locked_until,omitempty"` // Lockout for brute force protection
	CreatedAt        time.Time  `json:"created_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"-"`
}

// IsLocked checks if verification is locked due to too many failed attempts
func (e *EmailVerification) IsLocked() bool {
	if e.LockedUntil == nil {
		return false
	}
	return time.Now().Before(*e.LockedUntil)
}

// BeforeCreate hook untuk generate UUID
func (e *EmailVerification) BeforeCreate(tx *gorm.DB) error {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return nil
}
