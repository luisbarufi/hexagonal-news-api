package news_http

import (
	"hexagonal-news-api/adapter/output/model/response"
	"hexagonal-news-api/application/domain"
	"hexagonal-news-api/configuration/env"
	"hexagonal-news-api/configuration/rest_err"

	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type newsClient struct{}

func NewNewsClient() *newsClient {
	client = resty.New().SetBaseURL("https://newsapi.org/v2")
	return &newsClient{}
}

var (
	client *resty.Client
)

func (nc *newsClient) GetNewsPort(
	newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {

	newsResponse := &response.NewsClientResponse{}

	_, err := client.R().
		SetQueryParams(map[string]string{
			"q":      newsDomain.Subject,
			"from":   newsDomain.From,
			"apiKey": env.GetNewsTokenAPI(),
		}).
		SetResult(newsResponse).
		Get("/everything")

	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to call NewsAPI with params")
	}

	NewsResponseDomain := &domain.NewsDomain{}
	copier.Copy(NewsResponseDomain, newsResponse)

	return NewsResponseDomain, nil
}
