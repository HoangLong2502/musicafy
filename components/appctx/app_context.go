package appctx

import (
	"example.com/musicafy_be/components/mail"
	"example.com/musicafy_be/components/token"
	zingmp3 "example.com/musicafy_be/components/zing_mp3"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	// UploadProvider() upload.UploadProvider
	// GetWorker() woker.TaskDistributor
	GetTokenMaker() token.TokenMaker
	GetZingmp3Api() zingmp3.ZingMp3API
	GetMailer() mail.Mailer
}

type appCtx struct {
	db *gorm.DB
	// uploadProvider upload.UploadProvider
	// worker         woker.TaskDistributor
	token      token.TokenMaker
	zingmp3Api zingmp3.ZingMp3API
	mailer     mail.Mailer
}

func NewAppContext(db *gorm.DB, token token.TokenMaker, zingmp3Api zingmp3.ZingMp3API, mailer mail.Mailer) *appCtx {
	return &appCtx{
		db:         db,
		token:      token,
		zingmp3Api: zingmp3Api,
		mailer:     mailer,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB     { return ctx.db }
func (ctx *appCtx) GetTokenMaker() token.TokenMaker   { return ctx.token }
func (ctx *appCtx) GetZingmp3Api() zingmp3.ZingMp3API { return ctx.zingmp3Api }
func (ctx *appCtx) GetMailer() mail.Mailer            { return ctx.mailer }
