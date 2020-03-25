package main

import (
	"github.com/duclm2609/nplog"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4"
	"log"
	"net/http"
	"time"
)

// User
type Price struct {
	Price string `json:"price"`
}

func main() {
	// Setup logger
	config := nplog.NpLoggerOption{
		EnableConsole:  true,
		EnableFile:     true,
		FileJSONFormat: true,
		Filename:       "/var/logs/micro-price.log",
		FileMaxSize:    512,
		FileMaxBackups: 1,
		FileMaxAge:     1,
		FileCompress:   true,
		FileLevel:      nplog.Info,
	}
	logger, err := nplog.NewNpLogger(nplog.ZapLogger, config)
	if err != nil {
		log.Fatal("failed to initialize logger, exit")
	}
	// Add specific service field
	logger = logger.With(nplog.Fields{"service": "micro-price"})

	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.GET("/price/:id", func(context echo.Context) error {
		time.Sleep(1300 * time.Millisecond)
		res := &Price{Price: "7,369,000,000"}
		logger.For(context.Request().Context()).Infof("<-- GET /price/:id OK 200")
		return context.JSON(http.StatusOK, res)
	})

	_ = e.Start(":8080")
}
