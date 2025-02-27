package userbiz

import (
	"example.com/musicafy_be/common"
	usermodels "example.com/musicafy_be/modules/user/models"
	"example.com/musicafy_be/utils"
)

type RegisterStore interface {
	FindUser(arg usermodels.AccountQueries) (usermodels.User, error)
	CreateAccount(data usermodels.User) (*int, error)
}

type registerBiz struct {
	store RegisterStore
}

type RegisterReq struct {
	Username *string `json:"user_name"`
	Email    *string
	Password *string `json:"password"`
	FullName *string `json:"full_name"`
}

func NewRegisterBiz(store RegisterStore) (*registerBiz, error) {
	return &registerBiz{
		store: store,
	}, nil
}

func (biz *registerBiz) Register(req RegisterReq) (*int, error) {
	_, err := biz.store.FindUser(usermodels.AccountQueries{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		if err.Error() != "record not found" {
			return nil, common.NewCustomError(err, "find user with info", "Thông tin đăng ký đã tồn tại", "REGISTER")
		}
	}

	hassPassword, err := utils.HashedPassword(*req.Password)
	if err != nil {
		return nil, common.NewCustomError(err, "error hash password", "Lỗi hashpassword", "REGISTER")
	}

	id, err := biz.store.CreateAccount(usermodels.User{
		Username:       *req.Username,
		HashedPassword: hassPassword,
		FullName:       *req.FullName,
		Email:          req.Email,
	})
	if err != nil {
		return nil, common.NewCustomError(err, "error create user", "Lỗi tạo bản ghi User", "REGISTER")
	}

	return id, nil
}
