package ginsong

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	songsbiz "example.com/musicafy_be/modules/songs/biz"
	songstorage "example.com/musicafy_be/modules/songs/storage"
	"github.com/gin-gonic/gin"
)

func DetailSong(appContext appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appContext.GetMainDBConnection()
		store := songstorage.NewStore(db)
		biz := songsbiz.NewDetailSongBiz(store)

		arg := songsbiz.DetailSongReq{
			ID: c.Query("id"),
		}
		song, err := biz.DetailSong(appContext, arg)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(song))
	}
}
