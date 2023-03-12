
import mongoose from 'mongoose';
const db = mongoose.createConnection('mongodb://127.0.0.1:27017/nft');
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
