package dune

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-resty/resty/v2"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website/whale"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

type IDuneClient interface {
	GetWhaleAddress() ([]whale.AddressLabel, error)
}

type DuneClientConfig struct {
	ClientUrl string
	ResultId  string
	ErrorId   string
}

type DuneClient struct {
	cfg    *DuneClientConfig
	ctx    context.Context
	client *resty.Client
}

func NewDuneClient(cfg *DuneClientConfig) *DuneClient {
	client := resty.New()
	client.SetHostURL(cfg.ClientUrl)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			url := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, url, website.ErrHTTPError)
		}
		return nil
	})
	return &DuneClient{
		cfg:    cfg,
		ctx:    context.Background(),
		client: client,
	}
}

func (c *DuneClient) GetWhaleAddress() ([]whale.AddressLabel, error) {
	req := map[string]interface{}{
		"operationName": "FindResultDataByResult",
		"variables": map[string]interface{}{
			"result_id": c.cfg.ResultId,
			"error_id":  c.cfg.ErrorId,
		},
		"query": "query FindResultDataByResult($result_id: uuid!, $error_id: uuid!) {\n  query_results(where: {id: {_eq: $result_id}}) {\n    id\n    job_id\n    runtime\n    generated_at\n    columns\n    __typename\n  }\n  query_errors(where: {id: {_eq: $error_id}}) {\n    id\n    job_id\n    runtime\n    message\n    metadata\n    type\n    generated_at\n    __typename\n  }\n  get_result_by_result_id(args: {want_result_id: $result_id}) {\n    data\n    __typename\n  }\n}\n",
	}
	response, err := c.client.R().SetHeader("x-hasura-api-key", "").SetBody(req).Post("v1/graphql")
	if err != nil {
		return nil, fmt.Errorf("cannot fetch enqueue: %w", err)
	}
	jsonObj := gjson.ParseBytes(response.Body())
	resultList := jsonObj.Get("data.get_result_by_result_id").Array()
	var duneResultList []whale.AddressLabel
	for _, v1 := range resultList {
		data := v1.Get("data")
		price, err := decimal.NewFromString(data.Get("持有代币价格").String())
		if err != nil {
			log.Error("Parse price error", "err", err)
			return nil, err
		}
		amount, err := decimal.NewFromString(data.Get("持有代币余额").String())
		if err != nil {
			log.Error("Parse amount error", "err", err)
			return nil, err
		}
		amountUsd, err := decimal.NewFromString(data.Get("持有usd总金额").String())
		if err != nil {
			log.Error("Parse amount usd error", "err", err)
			return nil, err
		}
		dune := whale.AddressLabel{
			ChainName:    "ETH",
			CoinName:     data.Get("持有代币名称").String(),
			ContractAddr: data.Get("持有代币合约地址").String(),
			AccountAddr:  data.Get("钱包地址").String(),
			Holder:       data.Get("钱包持有者").String(),
			Price:        price.String(),
			Amount:       amount.String(),
			AmountUsd:    amountUsd.String(),
			TxCount:      "0",
			AddressType:  "eoa",
		}
		duneResultList = append(duneResultList, dune)
	}
	return duneResultList, err
}
