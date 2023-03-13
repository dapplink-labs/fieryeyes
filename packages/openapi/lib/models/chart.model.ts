
import mongoose from 'mongoose';
import config from "../../config/index";
const db = mongoose.createConnection(config.mongoHost);
const { Schema } = mongoose;

const ChartSchema = new Schema({
    id: Number,
    nftValues: [{
        value: String,
        perChange: String
    }],
    collections: [{
        value: String,
        perChange: String
    }],
    whaleHolders: [{
        value: String,
        perChange: String
    }],
    nfts: [{
        value: String,
        perChange: String
    }],
    created: { type: Date, default: Date.now }
})

const Chart = db.model('Chart', ChartSchema);

export default Chart
