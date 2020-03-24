package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/duclm2609/micro-api-gateway/logger"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4"
	"go.elastic.co/apm/module/apmhttp"
	"golang.org/x/net/context/ctxhttp"
)

type ProductDetail struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    string `json:"price"`
}

// User
type Review struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating"`
}

type Product struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Price    string   `json:"price"`
	Reviews  []Review `json:"reviews"`
}

func main() {
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.GET("/products/:id", func(context echo.Context) error {
		logger.Ctx(context.Request().Context()).Sugar().Info("--> GET /products/:id")
		productId := context.Param("id")
		time.Sleep(3 * time.Millisecond)
		client := apmhttp.WrapClient(http.DefaultClient)

		// Call inventory detail service
		productDetailCh, reviewsCh := make(chan ProductDetail), make(chan []Review)


		go callInventory(context, client, productId, productDetailCh)

		// Call review service
		go callReview(context, client, productId, reviewsCh)

		productDetail, reviews := <- productDetailCh, <-reviewsCh

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

func callReview(context echo.Context, client *http.Client, productId string, ch chan<- []Review) {
	logger.Ctx(context.Request().Context()).Sugar().Info("call Reviews service for products' review")
	reviewRes, err := ctxhttp.Get(context.Request().Context(), client, "http://micro-review:8080/review/"+productId)
	if err != nil {
		log.Fatal("failed to fetch product's review:")
	}
	defer reviewRes.Body.Close()
	var reviews []Review
	reviewContent, _ := ioutil.ReadAll(reviewRes.Body)
	_ = json.Unmarshal(reviewContent, &reviews)
	ch <- reviews
}

func callInventory(context echo.Context, client *http.Client, productId string, ch chan<- ProductDetail) {
	logger.Ctx(context.Request().Context()).Sugar().Info("call Inventory service for product details")
	resp, err := ctxhttp.Get(context.Request().Context(), client, "http://micro-inventory:8080/"+productId)
	if err != nil {
		log.Fatal("failed to fetch product detail: err=" + err.Error())
	}
	defer resp.Body.Close()
	productDetail := ProductDetail{}
	content, err := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(content, &productDetail)
	ch <- productDetail
}
