package userbiz

import (
	"context"
	"errors"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/token"
	usermodels "example.com/musicafy_be/modules/user/models"
	"example.com/musicafy_be/utils"
)

type LoginStore interface {
	FindUser(ctx context.Context, arg usermodels.User) (usermodels.User, error)
	CreateSession(ctx context.Context, arg usermodels.Session) (*string, error)
}

type loginBiz struct {
	store LoginStore
}

func NewLoginBiz(store LoginStore) *loginBiz {
	return &loginBiz{
		store: store,
	}
}

func (biz *loginBiz) LoginBiz(ctx context.Context, arg usermodels.LoginReq, token token.TokenMaker) (*usermodels.LoginRes, error) {
	user, err := biz.store.FindUser(ctx, usermodels.User{
		Username: arg.Username,
	})
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, common.ErrInvalidRequest(errors.New("account inactive"))
	}

	if err = utils.CheckPassword(arg.Password, user.HashedPassword); err != nil {
		return nil, usermodels.ErrUsernameOrPasswordInvalid
	}

	accressToken, accressTokenPayload, _ := token.CreateToken(user)
	refreshToken, refreshTokenPayload, _ := token.CreateToken(user)

	session, err := biz.store.CreateSession(ctx, usermodels.Session{
		ID:           refreshTokenPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		// UserAgent:    metadata.UserAgent,
		ClientIp:  arg.ClientIp,
		IsBlocked: false,
		ExpiresAt: refreshTokenPayload.ExpireAt,
	})
	if err != nil {
		return nil, common.NewCustomError(err, "failed to create session: %e", "Lỗi tạo session đăng nhập", "SESSION_CREATE")
		// return nil, status.Errorf(codes.Internal, "failed to create session: %e", err)
	}

	res := usermodels.LoginRes{
		User:                  user,
		SessionId:             *session,
		AccessToken:           accressToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accressTokenPayload.ExpireAt,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpireAt,
	}

	return &res, nil
}
