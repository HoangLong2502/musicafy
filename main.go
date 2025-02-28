package main

import (
	"errors"

	"example.com/musicafy_be/components/appctx"
	"example.com/musicafy_be/components/token"
	zingmp3 "example.com/musicafy_be/components/zing_mp3"
	"example.com/musicafy_be/middleware"
	"example.com/musicafy_be/router"
	"example.com/musicafy_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	dsn := "host=localhost user=root password=localhost1234 dbname=musicafy port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot open database")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	tokenMaker, err := token.NewPasetoMaker(config.TOKEN_SYMMETRIC_KEY)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create token maker")
	}

	zingmp3 := zingmp3.NewZingMp3Api(config.ZINGMP3_AC_URL, config.ZINGMP3_URL, "", "")

	appContext := appctx.NewAppContext(db, tokenMaker, zingmp3)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	v1 := r.Group("/v1")
	router.SetupRoute(appContext, v1)
	r.Run()
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}
	log.Info().Msg("db migrated successfully")
}
