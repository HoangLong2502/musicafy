package usermodels

import (
	"example.com/musicafy_be/common"
)

type Verify struct {
	common.SQLModel `json:",inline"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	SecretCode      string `json:"secret_code"`
}

func (Verify) TableName() string {
	return "verifies"
}
