package main

import (
	"log"
	"os"
	"time"

	"github.com/hsxflowers/cat-api/config"
	"github.com/hsxflowers/cat-api/internal/http/router"
)

const TIMEOUT = 30 * time.Second

// @title			cat-api
// @contact.name	hsxflowers
// @version			1.0
// @BasePath		/v1
func main() {
	var err error
	config.Envs, err = config.LoadEnvVars()

	if err != nil {
		log.Fatalln("Failed loading env", err)
	}

	h := router.Handlers(config.Envs)

	err = h.Start(":8090")
	if err != nil {
		log.Fatal("Error running API: ", err)
		os.Exit(1)
	}
}
