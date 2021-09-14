package main

import (
	"flag"
	"fmt"
	"google-analytics-101/utils"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/analytics/v3"

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

	way2()
	// apiKey()
	// apiKey2Func()

	return app
}

// // apiKey shows how to use an API key to authenticate.
// func apiKeyFunc() error {
// 	client, err := pubsub.NewClient(context.Background(), "sale-page-325603", option.WithAPIKey(apiKey))
// 	if err != nil {
// 		return fmt.Errorf("pubsub.NewClient: %v", err)
// 	}
// 	defer client.Close()
// 	// Use the authenticated client.
// 	_ = client

// 	return nil
// }

// func apiKey2Func() {
// 	analyticsService, err := analytics.NewService(context.Background(), option.WithAPIKey(apiKey))
// 	if err != nil {
// 		logrus.Error(err)
// 	}

// 	x, err := analyticsService.Management.Accounts.List().Do()
// 	p(err)

// 	// x := analyticsService.Management.Accounts.List().Header()
// 	logrus.Infof("%v", x)
// }

func way2() {
	// key, _ := ioutil.ReadFile("key.json")
	key, _ := ioutil.ReadFile("sale-page-325603-7e23a25820d0.json")

	jwtConf, err := google.JWTConfigFromJSON(
		key,
		analytics.AnalyticsReadonlyScope,
	)
	p(err)

	httpClient := jwtConf.Client(oauth2.NoContext)
	svc, err := analytics.New(httpClient)
	p(err)

	accountResponse, err := svc.Management.Accounts.List().Do()
	p(err)

	var accountId string

	fmt.Println("Found the following accounts:")
	for i, acc := range accountResponse.Items {

		if i == 0 {
			accountId = acc.Id
		}

		fmt.Println(acc.Id, acc.Name)
	}

	webProps, err := svc.Management.Webproperties.List(accountId).Do()
	p(err)

	var wpId string

	fmt.Println("\nFound the following properties:")
	for i, wp := range webProps.Items {

		if i == 0 {
			wpId = wp.Id
		}

		fmt.Println(wp.Id, wp.Name)
	}

	profiles, err := svc.Management.Profiles.List(accountId, wpId).Do()
	p(err)

	var viewId string

	fmt.Println("\nFound the following profiles:")
	for i, p := range profiles.Items {

		if i == 0 {
			viewId = "ga:" + p.Id
		}

		fmt.Println(p.Id, p.Name)
	}

	fmt.Println("\nTime to retrieve the real time users of this profile")

	metrics := "rt:activeUsers"
	rt, err := svc.Data.Realtime.Get(viewId, metrics).Do()
	p(err)

	fmt.Println(rt.Rows)
}

func p(err error) {
	if err != nil {
		panic(err)
	}
}
