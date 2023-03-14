'use strict';

const { Controller } = require('egg');
const OpenseaScraper = require("opensea-scraper");

class OpenseaController extends Controller {
    async saveRank() {
        console.log("saveRank start")

        const { ctx } = this;
        const ret = await ctx.service.opensea.saveCollection();
        console.log(ret.affectedRows)

        ctx.body = "return:"+ret
    }
}

module.exports = OpenseaController;
