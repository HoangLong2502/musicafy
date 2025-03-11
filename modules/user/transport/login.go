package ginuser

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	userbiz "example.com/musicafy_be/modules/user/biz"
	userstorage "example.com/musicafy_be/modules/user/storage"
	"github.com/gin-gonic/gin"
)

func Login(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		store := userstorage.NewStore(db)
		token := appctx.GetTokenMaker()
		biz := userbiz.NewLoginBiz(store)

		var arg userbiz.LoginReq
		if err := c.ShouldBind(&arg); err != nil {
			panic(err)
		}
		clientIp := c.ClientIP()
		arg.ClientIp = &clientIp
		user, err := biz.LoginBiz(c.Request.Context(), arg, token)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
