package oklink

import (
	"context"

	"github.com/savour-labs/fieryeyes/sav-scrapy/models"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"github.com/weblazy/easy/utils/elog"
	"go.uber.org/zap"
)

func (c *OklinkClient) ScrapyEth() {
	ctx := context.Background()
	resp, err := c.OkLinkRequest(ctx, nil)
	if err != nil {
		return
	}
	jsonObj := gjson.ParseBytes(resp)
	resultList := jsonObj.Get("data").Array()
	params := make([]map[string]interface{}, 0)
	for _, data := range resultList {
		amount, err := decimal.NewFromString(data.Get("amount").String())
		if err != nil {
			elog.ErrorCtx(ctx, "amountErr", zap.Error(err))
		}

		obj := map[string]interface{}{
			"chain_name":   "ETH",
			"account_addr": data.Get("address").String(),
			"amount":       amount,
		}
		params = append(params, obj)
	}
	field := []string{"chain_name", "account_addr", "amount"}
	err = models.ChainAccountHandler.BulkSave(nil, ctx, field, params)
	if err != nil {
		elog.ErrorCtx(ctx, "ChainAccountHandler.BulkSaveErr", zap.Error(err))
	}
}
