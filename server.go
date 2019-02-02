package main

import (
	"fmt"
	"net/http"

	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"

	"github.com/andrideng/tax-calculator/app"
	"github.com/andrideng/tax-calculator/errors"
	routing "github.com/go-ozzo/ozzo-routing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Sirupsen/logrus"
)

func main() {
	// - load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// - load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// - create the logger
	logger := logrus.New()

	// - connect to the database
	db, err := dbx.MustOpen("mysql", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	// - wire up API routing
	http.Handle("/", buildRouter(logger, db))

	// - start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()
	str := "Welcome to Tax Calculator Version: " + app.Version
	router.To("GET,HEAD", "/", func(c *routing.Context) error {
		c.Abort() // skip all other middlewares/ahdnlers
		return c.Write(str)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	rg := router.Group("/api")

	// - default route
	rg.Get("", func(c *routing.Context) error {
		return c.Write(str)
	})
	rg.Get("/ping", func(c *routing.Context) error {
		return c.Write("PONG")
	})

	return router
}
