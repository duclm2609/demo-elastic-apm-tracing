package domain

import (
	"context"
	"encoding/json"
	"github.com/duclm2609/nplog"
	"go.elastic.co/apm/module/apmhttp"
	"golang.org/x/net/context/ctxhttp"
	"io/ioutil"
	"net/http"
)

type ProductDetail struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    string `json:"price"`
}

type InventoryService struct {
	logger     nplog.NpLogger
	httpClient *http.Client
}

func NewInventoryService(logger nplog.NpLogger) InventoryService {
	return InventoryService{logger: logger, httpClient: apmhttp.WrapClient(http.DefaultClient)}
}

func (i InventoryService) GetDetail(ctx context.Context, productId string) ProductDetail {
	i.logger.For(ctx).Infof("getting product detail information")
	resp, err := ctxhttp.Get(ctx, i.httpClient, "http://micro-inventory:8080/"+productId)
	if err != nil {
		i.logger.For(ctx).Fatalf("failed to fetch product detail: err=" + err.Error())
	}
	defer resp.Body.Close()
	productDetail := ProductDetail{}
	content, err := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(content, &productDetail)
	return productDetail
}
