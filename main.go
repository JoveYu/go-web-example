package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"test.com/example/apps/base"
	"test.com/example/apps/user"
	"test.com/example/config"
	_ "test.com/example/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @host example.test.com
func main() {
	// parse config
	c, err := config.Parse()
	if err != nil {
		panic(err)
	}

	r := gin.New()

	// add logger
	zerolog.DurationFieldUnit = time.Millisecond // ms
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stdout,
				NoColor: false,
			},
		)
	} else {
		gin.SetMode(gin.ReleaseMode)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	r.Use(logger.SetLogger())
	r.Use(gin.Recovery())

	// add session
	store := cookie.NewStore([]byte("secret")) // use redis in production
	r.Use(sessions.Sessions("session", store))

	// add swagger
	swaggerUrl := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	// add url
	r.GET("/ping", base.PingView)
	user.SetRouter(r) // /user

	log.Info().
		Int("port", c.Port).
		Msg("server starting")
	err = r.Run(fmt.Sprintf(":%d", c.Port))
	if err != nil {
		panic(err)
	}
}
