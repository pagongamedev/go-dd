package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/pagongamedev/go-dd/example/ex001portal/docs"
	"github.com/pagongamedev/go-dd/example/ex001portal/notexample"

	goddPortal "github.com/pagongamedev/go-dd/portal"
	goddGofiber "github.com/pagongamedev/go-dd/support/gofiber"
)

// Can test api by localhost:8081/hello/v1/hello
// Go DD Portal Automatic create localhost:xxxx/health : localhost:8081/health localhost:8082/health localhost:8083/health

func main() {

	appMain := appMain()
	db, _ := notexample.DummyDatabase()

	portal := goddPortal.New()
	// Manage Defer Interface.Close()
	portal.AppendInterfaceClose(db)
	// Create API On Port 8081
	portal.AppendApp(appMain, ":8081")
	// Create Swagger Document On Port 8082
	portal.AppendApp(appAPIDocument(), ":8082")
	// Create Prometheus Metrics On Port 8083
	portal.AppendApp(goddGofiber.AppMetricsPrometheus(appMain), ":8083")

	portal.StartServer()
}

func appMain() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	notexample.Router(app, "/hello/v1")
	return app
}

func appAPIDocument() *fiber.App {

	docs.SwaggerInfo.Title = "ex001portal"
	docs.SwaggerInfo.Version = "v0.0.1"
	docs.SwaggerInfo.Description = "..."
	docs.SwaggerInfo.Host = "localhost:8081"
	return goddGofiber.AppAPIDocument()
}
