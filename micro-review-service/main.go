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
type Review struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating"`
}

func main() {
	// Setup logger
	config := nplog.NpLoggerOption{
		EnableConsole:  true,
		EnableFile:     true,
		FileJSONFormat: true,
		Filename:       "/var/logs/micro-api-gateway.log",
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
	logger = logger.With(nplog.Fields{"service": "micro-api-gateway"})

	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.GET("/review/:id", func(context echo.Context) error {
		time.Sleep(1400 * time.Millisecond) // Fake slow DB query
		logger.For(context.Request().Context()).Infof("load customer reviews from database")
		var res []Review
		res = append(res, Review{"1", "frank", "This is a fucking luxurious car", 5})
		res = append(res, Review{"2", "beckham", "I have one of this. And it's perfect.", 5})
		res = append(res, Review{"3", "rooney", "This car is masterpiece.", 5})
		logger.For(context.Request().Context()).Infof("<-- GET /review/:id OK 200")
		return context.JSON(http.StatusOK, res)
	})

	_ = e.Start(":8080")
}
