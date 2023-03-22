package etherscan

import (
	"context"
	"emperror.dev/errors"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website/whale"
	"strings"
)

type EtherClientConfig struct {
	ClientUrl string
}

type EtherClient struct {
	cfg *EtherClientConfig
	ctx context.Context
}

func NewEtherClient(cfg *EtherClientConfig) *EtherClient {
	return &EtherClient{
		cfg: cfg,
		ctx: context.Background(),
	}
}

func (ec *EtherClient) CollEtherLabelAddress() (*colly.HTMLElement, error) {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	var htmlElement *colly.HTMLElement
	c.OnHTML("body", func(e *colly.HTMLElement) {
		if e.Request.URL.String() == ec.cfg.ClientUrl {
			e.ForEach(".secondary-container a", func(i int, el *colly.HTMLElement) {
				link := el.Attr("href")
				log.Info("Link found: %q -> %s\n", "el.Text", el.Text, "link", link)
				e.Request.Visit(link)
			})
		} else {
			htmlElement = e
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Info("visiting", "url", r.URL.String())
	})
	c.Visit(ec.cfg.ClientUrl)
	return htmlElement, nil
}

func (ec *EtherClient) DealData(element *colly.HTMLElement) ([]whale.AddressLabel, error) {
	pathArr := strings.Split(element.Request.URL.Path, "/")
	if len(pathArr) < 4 {
		log.Info("exec path", "path", element.Request.URL.Path)
		return nil, errors.New("exec path error")
	}
	var addrLabelList []whale.AddressLabel
	element.ForEach("tbody > tr", func(k1 int, e1 *colly.HTMLElement) {
		data := whale.AddressLabel{
			ChainName:   "ETH",
			AddressType: pathArr[1],
			Holder:      pathArr[3],
		}
		switch pathArr[1] {
		case "txs", "blocks":
			return
		case "accounts":
			e1.ForEach("td", func(k2 int, e2 *colly.HTMLElement) {
				switch k2 {
				case 0:
					data.AccountAddr = e2.Text
				case 1:
					data.Holder = e2.Text
				case 2:
					data.Amount = e2.Text
				case 3:
					data.TxCount = e2.Text
				}
			})
			addrLabelList = append(addrLabelList, data)
		case "tokens":
			// contract address
			e1.ForEach("td", func(k2 int, e2 *colly.HTMLElement) {
				switch k2 {
				case 0:
					// id no value
				case 1:
					data.AccountAddr = e2.Text
				case 2:
					data.Holder = e2.Text
				case 3:
					data.TxCount = "0"
				case 4:
					data.Amount = e2.Text
				}
			})
			addrLabelList = append(addrLabelList, data)
		default:
			log.Info("exec path", "path", element.Request.URL.Path)
		}
	})
	return addrLabelList, nil
}
