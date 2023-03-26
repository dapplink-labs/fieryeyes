package oklink

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

type IOkClient interface {
	GetWhaleAddress() ([]whale.AddressLabel, error)
}

type OkClientConfig struct {
	ClientUrl   string
	OkAccessKey string
}

type OkClient struct {
	cfg    *OkClientConfig
	ctx    context.Context
	client *resty.Client
}

func NewOkClient(cfg *OkClientConfig) *OkClient {
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
	return &OkClient{
		cfg: cfg,
		ctx: context.Background(),
	}
}

func (c *OkClient) GetWhaleAddress() ([]whale.AddressLabel, error) {
	response, err := c.client.R().SetHeader("Ok-Access-Key", c.cfg.OkAccessKey).Get("/api/v5/explorer/address/rich-list?chainShortName=eth")
	if err != nil {
		return nil, fmt.Errorf("cannot fetch whale address: %w", err)
	}
	jsonObj := gjson.ParseBytes(response.Body())
	resultList := jsonObj.Get("data").Array()
	var addrLabelList []whale.AddressLabel
	for _, data := range resultList {
		amount, err := decimal.NewFromString(data.Get("amount").String())
		if err != nil {
			log.Error("amount error", "err", err)
			return nil, err
		}
		obj := whale.AddressLabel{
			ChainName:   "ETH",
			AccountAddr: data.Get("address").String(),
			Amount:      amount.String(),
		}
		addrLabelList = append(addrLabelList, obj)
	}
	return addrLabelList, err
}
