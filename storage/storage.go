package storage

import (
	"context"
	"test/api/models"
	"time"
)

type IStorage interface {
	Close()
	User() IUserStorage
	OTP() IOTPStorage // Add this line

}

type IUserStorage interface {
	Create(context.Context, models.CreateUser) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.User, error)
	GetList(context.Context, models.GetListRequest) (models.UsersResponse, error)
	Update(context.Context, models.UpdateUser) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	GetPassword(context.Context, string) (string, error)
	UpdatePassword(context.Context, models.UpdateUserPassword) error
	GetUserCredentialsByLogin(context.Context, string) (models.User, error)
}
type IOTPStorage interface {
	Create(context.Context, models.OTP) error
	GetByUserIDAndCode(context.Context, string, string) (models.OTP, error)
	UpdateUsedAt(context.Context, string, time.Time) error
}
