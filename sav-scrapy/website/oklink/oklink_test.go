package oklink

import (
	"os"
	"testing"

	"github.com/savour-labs/fieryeyes/sav-scrapy/models"
	"github.com/smartystreets/goconvey/convey"
)

func TestScrapyEth(t *testing.T) {
	convey.Convey("TestScrapyEth", t, func() {
		models.NewMysqlClient()
		NewOklinkClient("https://www.oklink.com/api/v5/explorer/address/rich-list?chainShortName=eth", os.Getenv("Ok_Access_Key")).ScrapyEth()
	})
}
