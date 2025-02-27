package token

import (
	"errors"

	"example.com/musicafy_be/common"
	usermodels "example.com/musicafy_be/modules/user/models"
)

type TokenMaker interface {
	CreateToken(data usermodels.User) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"token không tồn tại",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"lỗi giải mã token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("token invalid"),
		"token invalid",
		"token không hợp lệ",
		"ErrInvalidToken",
	)
)
