package oklink

import (
	"context"

	"github.com/weblazy/easy/utils/elog"
	"github.com/weblazy/easy/utils/http/http_client"
	"github.com/weblazy/easy/utils/http/http_client/http_client_config"
	"go.uber.org/zap"
)

type OklinkClient struct {
	Url         string
	OkAccessKey string
}

func NewOklinkClient(url, okAccessKey string) *OklinkClient {
	return &OklinkClient{
		Url:         url, //"https://www.oklink.com/api/v5/explorer/address/rich-list?chainShortName=eth"
		OkAccessKey: okAccessKey,
	}
}

func (c *OklinkClient) OkLinkRequest(ctx context.Context, req interface{}) ([]byte, error) {
	cfg := http_client_config.DefaultConfig()
	client := http_client.NewHttpClient(cfg)
	request := client.Request.SetContext(ctx)
	resp, err := request.SetHeader("Ok-Access-Key", c.OkAccessKey).Get(c.Url)
	if err != nil {
		elog.ErrorCtx(ctx, "OkLinkRequestErr", zap.Error(err))
		return nil, err
	}
	return resp.Body(), nil
}
