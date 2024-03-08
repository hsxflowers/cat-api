package router

import (
	"context"

	"github.com/hsxflowers/cat-api/cat"
	"github.com/hsxflowers/cat-api/cat/domain"
	"github.com/hsxflowers/cat-api/config"
	catHandler "github.com/hsxflowers/cat-api/internal/http/cat" // cat Handler
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Handlers(envs *config.Environments) *echo.Echo {
	e := echo.New()
	ctx := context.Background()

	var catDb domain.CatDatabase
	// var err error
	// var cassandraSession *cassandra.CassandraConnection

	// cassandraSession, err = cassandra.NewCassandraDatabase(envs, log)

	// if err != nil {
	// 	log.Errorf("Error connecting to Cat CassandraDB: %v", err)
	// 	panic(err)
	// }

	// catDb = catCassandra.NewCassandraStore(cassandraSession, envs, log)

	log.Debug("")

	catRepository := cat.NewCatRepository(catDb)
	catService := cat.NewCatService(catRepository)
	catHandler := catHandler.NewCatHandler(ctx, catService)

	e.GET("/swagger*", echoSwagger.WrapHandler)

	cat := e.Group("cat")

	cat.GET("/:cat_id", catHandler.Get)
	cat.POST("", catHandler.Create)

	return e
}
