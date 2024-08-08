package service

import (
	"test/pkg/logger"
	"test/storage"
)

type IServiceManager interface {
	User() userService

	AuthService() authService
	OTPService() otpService
}

type Service struct {
	userService userService
	authService authService
	otpService  otpService
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}

	services.userService = NewUserService(storage, log)
	services.authService = NewAuthService(storage, log)
	services.otpService = NewOTPService(storage, log)

	return services
}

func (s Service) User() userService {
	return s.userService
}

func (s Service) AuthService() authService {
	return s.authService
}
func (s Service) OTPService() otpService {
	return s.otpService
}
