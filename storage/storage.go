package storage

import (
	"context"
	"test/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
}

type IUserStorage interface {
	Create(context.Context, models.CreateUser) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.User, error)
	GetList(context.Context, models.GetListRequest) (models.UsersResponse, error)
	Update(context.Context, models.UpdateUser) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	GetPassword(context.Context, string) (string, error)
	UpdatePassword(context.Context, models.UpdateUserPassword) error
	GetAdminCredentialsByLogin(context.Context, string) (models.User, error)
}
