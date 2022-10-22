package oklink

import (
	"context"
	"os"

	"github.com/weblazy/easy/utils/elog"
	"github.com/weblazy/easy/utils/http/http_client"
	"github.com/weblazy/easy/utils/http/http_client/http_client_config"
	"go.uber.org/zap"
)

func OkLinkRequest(ctx context.Context, req interface{}) ([]byte, error) {
	cfg := http_client_config.DefaultConfig()
	client := http_client.NewHttpClient(cfg)
	request := client.Request.SetContext(ctx)
	key := os.Getenv("Ok_Access_Key")
	resp, err := request.SetHeader("Ok-Access-Key", key).Get("https://www.oklink.com/api/v5/explorer/address/rich-list?chainShortName=eth")
	if err != nil {
		elog.ErrorCtx(ctx, "OkLinkRequestErr", zap.Error(err))
		return nil, err
	}
	return resp.Body(), nil
}
