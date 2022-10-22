package etherscan

import (
	"context"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/savour-labs/fieryeyes/sav-scrapy/models"
	"github.com/weblazy/easy/utils/elog"
	"go.uber.org/zap"
)

func EtherscanRequest() ([]byte, error) {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	// On every a element which has href attribute call callback
	c.OnHTML("body", func(e *colly.HTMLElement) {
		if e.Request.URL.String() == "https://etherscan.io/labelcloud" {
			e.ForEach(".secondary-container a", func(i int, el *colly.HTMLElement) {
				link := el.Attr("href")
				fmt.Printf("Link found: %q -> %s\n", el.Text, link)
				e.Request.Visit(link)
			})
		} else {
			DealData(e)
		}

	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://etherscan.io/labelcloud")
	return []byte{}, nil
}

func DealData(e *colly.HTMLElement) {
	ctx := context.Background()
	pathArr := strings.Split(e.Request.URL.Path, "/")
	if len(pathArr) < 4 {
		elog.InfoCtx(ctx, "excpet path", zap.String("path", e.Request.URL.Path))
		return
	}
	params := make([]map[string]interface{}, 0)
	e.ForEach("tbody > tr", func(k1 int, e1 *colly.HTMLElement) {
		data := map[string]interface{}{
			"chain_name": "ETH",
			"addr_type":  pathArr[1],
			"holder":     pathArr[3],
		}
		switch pathArr[1] {
		case "txs", "blocks":
			return
		case "accounts":
			e1.ForEach("td", func(k2 int, e2 *colly.HTMLElement) {
				switch k2 {
				case 0:
					data["account_addr"] = e2.Text
				case 1:
					data["tag"] = e2.Text
				case 2:
					data["amount"] = e2.Text
				case 3:
					data["tx_count"] = e2.Text
				}
			})
			params = append(params, data)
		case "tokens":
			e1.ForEach("td", func(k2 int, e2 *colly.HTMLElement) {
				switch k2 {
				case 0:

				case 1:
					data["account_addr"] = e2.Text
				case 2:
					data["tag"] = e2.Text
				case 3:
					data["tx_count"] = "0"
				case 4:
					data["amount"] = e2.Text
				}
			})
			params = append(params, data)
		default:
			elog.InfoCtx(ctx, "excpet path", zap.String("path", e.Request.URL.Path))
		}

	})
	field := []string{"chain_name", "addr_type", "holder", "account_addr", "tag", "amount", "tx_count"}
	err := models.LabelHandler.BulkSave(nil, ctx, field, params)
	if err != nil {
		elog.InfoCtx(ctx, "LabelHandler.BulkSaveErr", zap.Error(err), zap.Any("params", params))
	}
}
