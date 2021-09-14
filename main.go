package main

import (
	"flag"
	"fmt"
	"google-analytics-101/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

var port = flag.String("port", ":5000", "Port to listen on")

func main() {
	app := SetupApp()

	logrus.Fatal(app.Listen(*port))
}

func SetupApp() *fiber.App {
	utils.InitLogrus()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(500).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	// build.SetupVersion(app)

	// database.Init()

	app.Use(logger.New())
	app.Use(recover.New())

	// router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	return app
}
