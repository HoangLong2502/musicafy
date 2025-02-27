package usermodels

import (
	"errors"
	"time"

	"example.com/musicafy_be/common"
)

type User struct {
	ID                int        `json:"id" gorm:"primaryKey;column:id"`
	Username          string     `json:"username" gorm:"column:username"`
	HashedPassword    string     `json:"hashed_password"`
	FullName          string     `json:"full_name" gorm:"column:full_name"`
	Email             *string    `json:"email,omitempty" gorm:"column:email"`
	IsVerify          bool       `json:"is_verify" gorm:"column:is_verify"`
	IsActive          bool       `json:"is_active" gorm:"column:active"`
	Role              *int       `json:"role,omitempty" gorm:"column:role"`
	Gender            *string    `json:"gender,omitempty"`
	Licence           *string    `json:"licence,omitempty"`
	Dob               *time.Time `json:"dob,omitempty"`
	Avatar            string     `json:"avatar,omitempty" gorm:"column:active"`
	PasswordChangedAt time.Time  `json:"password_changed_at"`
}

func (User) TableName() string {
	return "users"
}

type LoginReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	UserAgent string `json:"user_agent"`
	ClientIp  string `json:"client_ip"`
}

type LoginRes struct {
	SessionId             string    `json:"session_id,omitempty"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User                  User      `json:"data"`
}

type AccountQueries struct {
	ID       *int    `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"email hoặc mật khẩu không đúng",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"email đã tồn tại",
		"ErrEmailExisted",
	)
)
