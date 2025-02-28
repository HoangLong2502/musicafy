package ginsong

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	songsbiz "example.com/musicafy_be/modules/songs/biz"
	songstorage "example.com/musicafy_be/modules/songs/storage"
	"github.com/gin-gonic/gin"
)

func SearchSuggestion(appContext appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appContext.GetMainDBConnection()
		store := songstorage.NewStore(db)
		biz := songsbiz.NewSearchSuggestionBiz(store)

		var arg songsbiz.SearchSuggestionReq
		if err := c.ShouldBind(&arg); err != nil {
			panic(err)
		}

		songs, err := biz.SearchSuggestion(appContext, songsbiz.SearchSuggestionReq{
			Search: c.Query("search"),
		})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(songs))
	}
}
