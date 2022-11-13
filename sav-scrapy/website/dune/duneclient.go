package dune

import (
	"context"

	"github.com/weblazy/easy/utils/elog"
	"github.com/weblazy/easy/utils/http/http_client"
	"github.com/weblazy/easy/utils/http/http_client/http_client_config"
	"go.uber.org/zap"
)

type DueClient struct {
	Url string
}

func NewDueClient(url string) *DueClient {
	return &DueClient{
		Url: url, //"https://core-hsr.dune.com/v1/graphql"
	}
}

func (c *DueClient) DueRequest(ctx context.Context, req interface{}) ([]byte, error) {
	cfg := http_client_config.DefaultConfig()
	client := http_client.NewHttpClient(cfg)
	request := client.Request.SetContext(ctx)

	resp, err := request.SetHeader("x-hasura-api-key", "").SetBody(req).Post(c.Url)
	if err != nil {
		elog.ErrorCtx(ctx, "DueRequestErr", zap.Error(err))
		return nil, err
	}
	return resp.Body(), nil
}
