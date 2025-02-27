package ginsong

import (
	"example.com/musicafy_be/components/appctx"
	zingmp3 "example.com/musicafy_be/components/zing_mp3"
	"github.com/gin-gonic/gin"
)

func ListSong(appContext appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// e := "/song/get-song-info"
		// SuggestionSong("tháng năm", 10)
		zingmp3.SuggestionSong("tháng năm", 10)
	}
}
