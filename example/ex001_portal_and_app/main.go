package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/pagongamedev/go-dd/example/ex001_portal_and_app/docs"
	"github.com/pagongamedev/go-dd/example/ex001_portal_and_app/notexample"

	godd "github.com/pagongamedev/go-dd"
	goddPortal "github.com/pagongamedev/go-dd/portal"
	goddGofiber "github.com/pagongamedev/go-dd/support/gofiber"
)

// Can test api by localhost:8081/hello/v1/hello
// Go DD Portal Automatic create localhost:xxxx/health : localhost:8081/health localhost:8082/health localhost:8083/health

// Example ex001_portal_and_app
func main() {

	appMain := appMain()
	db, err := notexample.DummyDatabase()
	godd.MustError(err)

	portal := goddPortal.New()
	// Manage Defer Interface.Close()
	portal.AppendDeferClose(godd.DeferClose{Name: "database", I: db.(godd.InterfaceClose)})
	// Create API On Port 8081
	portal.AppendApp(appMain, ":8081")
	// Create Swagger Document On Port 8082
	portal.AppendApp(appAPIDocument(), ":8082")
	// Create Prometheus Metrics On Port 8083
	portal.AppendApp(goddGofiber.AppMetricsPrometheus(appMain), ":8083")

	portal.StartServer()
}

func appAPIDocument() *fiber.App {
	//docs generate by swaggo : swag init
	docs.SwaggerInfo.Title = "ex001portal"
	docs.SwaggerInfo.Version = "v0.0.1"
	docs.SwaggerInfo.Description = "..."
	docs.SwaggerInfo.Host = "localhost:8081"
	return goddGofiber.AppAPIDocument()
}

func appMain() *fiber.App {
	app := fiber.New()
	// Use App Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	notexample.Router(app, "/hello/v1")
	return app
}
