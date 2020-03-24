package main

import (
	"github.com/duclm2609/micro-review-service/logger"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4"
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
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.GET("/review/:id", func(context echo.Context) error {
		time.Sleep(1400 * time.Millisecond) // Fake slow DB query
		logger.Ctx(context.Request().Context()).Info("load customer reviews from database")
		var res []Review
		res = append(res, Review{"1", "frank", "This is a fucking luxurious car", 5})
		res = append(res, Review{"2", "beckham", "I have one of this. And it's perfect.", 5})
		res = append(res, Review{"3", "rooney", "This car is masterpiece.", 5})
		logger.Ctx(context.Request().Context()).Info("<-- GET /review/:id OK 200")
		return context.JSON(http.StatusOK, res)
	})

	_ = e.Start(":8080")
}
