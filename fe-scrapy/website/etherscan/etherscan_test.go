package etherscan

import (
	"testing"

	"github.com/savour-labs/fieryeyes/fe-scrapy/models"
	"github.com/smartystreets/goconvey/convey"
)

func TestEtherscan(t *testing.T) {
	convey.Convey("TestScrapyEth", t, func() {
		models.NewMysqlClient()
		NewEtherscanClient("https://etherscan.io/labelcloud").ScrapyEth()
	})
}
