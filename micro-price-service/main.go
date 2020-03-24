package main

import (
	"github.com/duclm2609/micro-price-service/logger"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4"
	"net/http"
	"time"
)

// User
type Price struct {
	Price string `json:"price"`
}

func main() {
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.GET("/price/:id", func(context echo.Context) error {
		time.Sleep(1300 * time.Millisecond)
		res := &Price{Price: "7,369,000,000"}
		logger.Ctx(context.Request().Context()).Info("<-- GET /price/:id OK 200")
		return context.JSON(http.StatusOK, res)
	})

	_ = e.Start(":8080")
}
