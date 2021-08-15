package restapi

import (
	"github.com/jdpadillaac/go-waitgroups-example/internal"
	"github.com/labstack/echo/v4"
	"log"
)

func Start(c *internal.AppConfig) {
	e := echo.New()

	log.Printf("app ready and running in %s port", c.AppPort)
	e.Logger.Fatal(e.Start(":" + c.AppPort))
}
