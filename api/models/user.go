package models

import "time"

type User struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Cash      uint      `json:"cash"`
	UserType  string    `json:"user_type"`
	BranchID  string    `json:"branch_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Login     string    `json:"login"`
}

type CreateUser struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

type UpdateUser struct {
	ID       string `json:"-"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

type UpdateUserPassword struct {
	ID          string `json:"-"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

