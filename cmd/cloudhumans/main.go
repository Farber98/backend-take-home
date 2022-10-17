package main

import (
	"cloudhumans/internal/config"
	"cloudhumans/internal/router"
	"log"
)

func main() {
	e := router.Init()
	log.Fatal(e.Start(config.Get().Context.Host + ":" + config.Get().Context.Port))
}
