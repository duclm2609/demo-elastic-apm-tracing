package main

import (
	"github.com/duclm2609/micro-api-gateway/domain"
	"github.com/duclm2609/nplog"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4"
)

type Product struct {
	Id       string          `json:"id"`
	Name     string          `json:"name"`
	Category string          `json:"category"`
	Price    string          `json:"price"`
	Reviews  []domain.Review `json:"reviews"`
}

func main() {
	// Setup logger
	config := nplog.Options{
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
	logger, err := nplog.New(nplog.ZapLogger, config)
	if err != nil {
		log.Fatal("failed to initialize logger, exit")
	}
	// Add specific service field
	logger = logger.With(nplog.Fields{"service.name": "micro-api-gateway"})

	// Setup server
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.GET("/products/:id", func(context echo.Context) error {
		logger.For(context.Request().Context()).Infof("--> GET /products/:id")
		productId := context.Param("id")
		time.Sleep(3 * time.Millisecond)

		// Call inventory detail service
		productDetailCh, reviewsCh := make(chan domain.ProductDetail), make(chan []domain.Review)
		go func() {
			inventoryService := domain.NewInventoryService(logger)
			productDetailCh <- inventoryService.GetDetail(context.Request().Context(), productId)
		}()

		// Call review service
		go func() {
			reviewService := domain.NewReviewService(logger)
			reviewsCh <- reviewService.GetDetail(context.Request().Context(), productId)
		}()

		productDetail, reviews := <-productDetailCh, <-reviewsCh

		// Aggregate response
		var product Product
		product.Id = productId
		product.Name = productDetail.Name
		product.Category = productDetail.Category
		product.Price = productDetail.Price
		product.Reviews = reviews

		return context.JSON(http.StatusOK, product)
	})

	_ = e.Start(":8080")
}
