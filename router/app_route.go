package router

import (
	"example.com/musicafy_be/components/appctx"
	"example.com/musicafy_be/middleware"
	ginsong "example.com/musicafy_be/modules/songs/transport"
	ginuser "example.com/musicafy_be/modules/user/transport"
	"github.com/gin-gonic/gin"
)

func SetupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	user := v1.Group("/user")
	{
		user.POST("/register", ginuser.Register(appContext))
		user.POST("/login", ginuser.Login(appContext))
	}

	song := v1.Group("/song")
	{
		song.GET("/search-suggestions", middleware.RequiredAuth(appContext), ginsong.SearchSuggestion(appContext))
		song.GET("/streaming", ginsong.StreamingSong(appContext))
		song.GET("/info-detail", ginsong.DetailSong(appContext))
	}
}
