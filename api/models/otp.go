package models

import "time"

// OTP model
type OTP struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
	UsedAt    *time.Time `json:"used_at,omitempty"`
}

// CreateOTPRequest model
type CreateOTPRequest struct {
	UserID string `json:"user_id"`
}

// VerifyOTPRequest model
type VerifyOTPRequest struct {
	UserID string `json:"user_id"`
	Code   string `json:"code"`
}
