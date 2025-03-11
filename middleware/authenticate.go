package middleware

import (
	"strings"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong authen header",
		"lỗi xác thực",
		"ErrWrongAuthHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(appctx appctx.AppContext) func(*gin.Context) {
	tokenProvider := appctx.GetTokenMaker()
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.VerifyToken(token)
		if err != nil {
			panic(err)
		}

		print(payload)
	}
}
