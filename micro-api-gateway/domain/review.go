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

type Review struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating"`
}

type ReviewService struct {
	logger     nplog.Logger
	httpClient *http.Client
}

func NewReviewService(logger nplog.Logger) ReviewService {
	return ReviewService{logger: logger, httpClient: apmhttp.WrapClient(http.DefaultClient)}
}

func (r ReviewService) GetDetail(ctx context.Context, productId string) []Review {
	r.logger.For(ctx).Infof("call Reviews service for products' review")
	reviewRes, err := ctxhttp.Get(ctx, r.httpClient, "http://micro-review:8080/review/"+productId)
	if err != nil {
		r.logger.For(ctx).Fatalf("failed to fetch product's review:")
	}
	defer reviewRes.Body.Close()
	var reviews []Review
	reviewContent, _ := ioutil.ReadAll(reviewRes.Body)
	_ = json.Unmarshal(reviewContent, &reviews)
	return reviews
}
