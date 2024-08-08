package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
	"time"

	"github.com/google/uuid"
)

type otpService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewOTPService(storage storage.IStorage, log logger.ILogger) otpService {
	return otpService{
		storage: storage,
		log:     log,
	}
}

// CreateOTP creates a new OTP and saves it to the database
func (o otpService) CreateOTP(ctx context.Context, request models.CreateOTPRequest) (models.OTP, error) {
	otpID := uuid.New().String()
	code := generateRandomOTPCode()               // Implement this function to generate OTP code
	expiresAt := time.Now().Add(15 * time.Minute) // OTP expires in 15 minutes

	otp := models.OTP{
		ID:        otpID,
		UserID:    request.UserID,
		Code:      code,
		ExpiresAt: expiresAt,
	}

	err := o.storage.OTP().Create(ctx, otp)
	if err != nil {
		o.log.Error("error while creating OTP", logger.Error(err))
		return models.OTP{}, err
	}

	return otp, nil
}

// VerifyOTP verifies the OTP code
func (o otpService) VerifyOTP(ctx context.Context, request models.VerifyOTPRequest) (bool, error) {
	otp, err := o.storage.OTP().GetByUserIDAndCode(ctx, request.UserID, request.Code)
	if err != nil {
		o.log.Error("error while getting OTP", logger.Error(err))
		return false, err
	}

	if otp.ID == "" || otp.UsedAt != nil || otp.ExpiresAt.Before(time.Now()) {
		return false, nil
	}

	// Mark OTP as used
	now := time.Now()
	err = o.storage.OTP().UpdateUsedAt(ctx, otp.ID, now)
	if err != nil {
		o.log.Error("error while updating OTP usage", logger.Error(err))
		return false, err
	}

	return true, nil
}

// generateRandomOTPCode generates a random OTP code
func generateRandomOTPCode() string {
	// Implement a function to generate a random OTP code
	return "123456" // Placeholder
}
