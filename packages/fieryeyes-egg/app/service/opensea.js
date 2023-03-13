const Service = require('egg').Service;
const OpenseaScraper = require("opensea-scraper");

class OpenseaService extends Service {
    async saveCollection() {
        console.log("saveCollection start")
        const ranking = await OpenseaScraper.rankings("total")

        const date = new Date();
        const dateStr = date.getFullYear()+"-"+date.getMonth()+"-"+date.getDate();

        let res = []
        for(let i = 0; i < ranking.length; i++) {
            res[i] = new Object()
            res[i].date = dateStr
            res[i].name = ranking[i].name
            res[i].slug = ranking[i].slug
            res[i].logo = ranking[i].logo


            res[i].floor_price_amount = ranking[i]?.floorPrice?.amount ?? "0"
            res[i].floor_price_currency = ranking[i]?.floorPrice?.currency ?? ""
        }


        const result = await this.app.mysql.insert('nft_collection', res)
        console.log("affectedRows:", result.affectedRows)
        return result
    }
}

module.exports = OpenseaService;