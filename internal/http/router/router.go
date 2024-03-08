package router

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"


	"github.com/hsxflowers/cat-api/cat"
	catDatabase "github.com/hsxflowers/cat-api/cat/db"
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
	var err error

    db, err := sql.Open("postgres", "postgres://hsxflowers:N9tiQ81qNuzKP8Axv0Ae8aXZU8Pg6APf@dpg-cnl6bn7sc6pc73cc28vg-a.oregon-postgres.render.com/dbcat?sslmode=require")
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }

	catDb = catDatabase.NewSQLStore(db)

	log.Debug("")

	catRepository := cat.NewCatRepository(catDb)
	catService := cat.NewCatService(catRepository)
	catHandler := catHandler.NewCatHandler(ctx, catService)

	e.GET("/swagger*", echoSwagger.WrapHandler)

	cat := e.Group("cat")

	cat.GET("/:tag", catHandler.Get)
	cat.POST("", catHandler.Create)

	return e
}
