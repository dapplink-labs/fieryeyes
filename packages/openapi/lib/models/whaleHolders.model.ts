
import mongoose from 'mongoose';
// const db = mongoose.createConnection('mongodb://127.0.0.1:27017/nft');
import config from "../../config/index";
const db = mongoose.createConnection(config.mongoHost);
const { Schema } = mongoose;

const HolderSchema = new Schema({
    id: Number,
    address: String,
    tokenValue: {
        ccy: String,
        totalValue: String,
        percentChange: String
    },
    holdNfts: {
        totalValue: String,
        imgs: [String]
    },
    holdCollections: {
        totalValue: String,
        imgs: [String]
    },
    realizedPnl: String,
    label: String,
    updated: { type: Date, default: Date.now }
})

const Holder = db.model('Holder', HolderSchema);
export default Holder
