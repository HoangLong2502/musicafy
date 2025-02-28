package router

import (
	"example.com/musicafy_be/components/appctx"
	ginsong "example.com/musicafy_be/modules/songs/transport"
	ginuser "example.com/musicafy_be/modules/user/transport"
	"github.com/gin-gonic/gin"
)

func SetupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	user := v1.Group("/user")
	{
		user.POST("/register", ginuser.Register(appContext))
	}

	song := v1.Group("/song")
	{
		song.GET("/search-suggestions", ginsong.SearchSuggestion(appContext))
		song.GET("/streaming", ginsong.StreamingSong(appContext))
		song.GET("/info-detail", ginsong.SearchSuggestion(appContext))
	}
}
