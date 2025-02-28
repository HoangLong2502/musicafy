package ginsong

import (
	"example.com/musicafy_be/components/appctx"
	"github.com/gin-gonic/gin"
)

func StreamingSong(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		// res := zingmp3.StreamFileSong("Z760BABZ")
		// if res == nil {
		// 	panic("không lấy được file streaming")
		// }

		// c.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}
