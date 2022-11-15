package dune

import (
	"context"

	"github.com/savour-labs/fieryeyes/sav-scrapy/models"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"github.com/weblazy/easy/utils/elog"
	"go.uber.org/zap"
)

func (c *DueClient) ScrapyEth() {
	ctx := context.Background()
	req := map[string]interface{}{
		"operationName": "FindResultDataByResult",
		"variables": map[string]interface{}{
			"result_id": "82ecc188-a294-4eba-8030-2b06d3d5c812",
			"error_id":  "9088a1aa-1c4c-416a-9cbb-a5ef6ae7cc81",
		},
		"query": "query FindResultDataByResult($result_id: uuid!, $error_id: uuid!) {\n  query_results(where: {id: {_eq: $result_id}}) {\n    id\n    job_id\n    runtime\n    generated_at\n    columns\n    __typename\n  }\n  query_errors(where: {id: {_eq: $error_id}}) {\n    id\n    job_id\n    runtime\n    message\n    metadata\n    type\n    generated_at\n    __typename\n  }\n  get_result_by_result_id(args: {want_result_id: $result_id}) {\n    data\n    __typename\n  }\n}\n",
	}
	resp, err := c.DueRequest(ctx, req)
	if err != nil {
		return
	}
	jsonObj := gjson.ParseBytes(resp)
	resultList := jsonObj.Get("data.get_result_by_result_id").Array()
	params := make([]map[string]interface{}, 0)
	for _, v1 := range resultList {
		data := v1.Get("data")
		price, err := decimal.NewFromString(data.Get("持有代币价格").String())
		if err != nil {
			elog.ErrorCtx(ctx, "priceErr", zap.Error(err))
		}
		amount, err := decimal.NewFromString(data.Get("持有代币余额").String())
		if err != nil {
			elog.ErrorCtx(ctx, "amountErr", zap.Error(err))
		}
		amountUsd, err := decimal.NewFromString(data.Get("持有usd总金额").String())
		if err != nil {
			elog.ErrorCtx(ctx, "amountUsdErr", zap.Error(err))
		}

		obj := map[string]interface{}{
			"chain_name":    "ETH",
			"coin_name":     data.Get("持有代币名称").String(),
			"contract_addr": data.Get("持有代币合约地址").String(),
			"account_addr":  data.Get("钱包地址").String(),
			"holder":        data.Get("钱包持有者").String(),
			"price":         price,
			"amount":        amount,
			"amount_usd":    amountUsd,
		}
		params = append(params, obj)
	}
	field := []string{"chain_name", "coin_name", "contract_addr", "account_addr", "holder", "price", "amount", "amount_usd"}
	err = models.ContractAccountHandler.BulkSave(nil, ctx, field, params)
	if err != nil {
		elog.ErrorCtx(ctx, "ContractAccountHandler.BulkSaveErr", zap.Error(err))
	}
}
