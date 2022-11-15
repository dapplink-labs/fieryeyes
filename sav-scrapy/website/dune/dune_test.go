package dune

import (
	"testing"

	"github.com/savour-labs/fieryeyes/sav-scrapy/models"
	"github.com/smartystreets/goconvey/convey"
)

func TestScrapyEth(t *testing.T) {
	convey.Convey("TestScrapyEth", t, func() {
		models.NewMysqlClient()
		NewDueClient("https://core-hsr.dune.com/v1/graphql").ScrapyEth()
	})
}
