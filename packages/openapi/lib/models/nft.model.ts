
import mongoose from 'mongoose';
import config from "../../config/index";
const db = mongoose.createConnection(config.mongoHost);
const { Schema } = mongoose;

const NftSchema = new Schema({
    id: Number,
    name: String,
    stars: String,
    whaleHolders: String,
    holders: String,
    mints: String,
    twentyfourhETHVolume: String,
    twentyfourhETHVolumePerchange: String,
    price: String,
    sevenDayShadowPrice: [String],
    listed: String,
    updated: { type: Date, default: Date.now }
})

const Nft = db.model('Nft', NftSchema);

export default Nft
