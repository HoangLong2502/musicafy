package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	// UploadProvider() upload.UploadProvider
	// GetWorker() woker.TaskDistributor
	// GetTokenMaker() token.TokenMaker
}

type appCtx struct {
	db *gorm.DB
	// uploadProvider upload.UploadProvider
	// worker         woker.TaskDistributor
	// token          token.TokenMaker
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{
		db: db,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
