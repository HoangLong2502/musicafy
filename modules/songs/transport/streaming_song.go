package ginsong

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	songsbiz "example.com/musicafy_be/modules/songs/biz"
	songstorage "example.com/musicafy_be/modules/songs/storage"
	"github.com/gin-gonic/gin"
)

func StreamingSong(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		store := songstorage.NewStore(db)
		biz := songsbiz.NewStreamingSongBiz(store)
		res, err := biz.StreamingSong(appctx, c.Query("id"))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}
