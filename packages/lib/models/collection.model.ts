
import mongoose from 'mongoose';
const db = mongoose.createConnection('mongodb://127.0.0.1:27017/nft');
const { Schema } = mongoose;

const CollectionSchema = new Schema({
    id: Number,
    collectionImg: String,
    name: String,
    chain: String,
    holders: String,
    whaleHolers: String,
    mint: String,
    mintDate: Date,
    twentyfourhrsValues: {
        count: String,
        percentChange: String
    },
    floorPrice: String,
    BestOffer: String,
    shadowScoreChange: [String],
    shadowScore: {
        price: String,
        whaleHold: String,
        volume: String,
        tnxs: String
    },
    totalPrice: {
        value: String,
        ccy: String,
        percentChange: String
    },
    totalHolders: {
        value: String,
        percentChange: String
    },
    totalWhaleHolders: {
        value: String,
        percentChange: String
    },
    totalTxns: {
        value: String,
        ccy: String,
        percentChange: String
    },
    trading: {
        price: [
            {   price: String,
                time: String,
                ccy: String
            }
        ],
        volume: [
            {
                time: String,
                value: String
            }
        ],
        list: [{
            time: String,
            value: String
        }],
        tnxs: [
            {
                time: String,
                value: String
            }
        ]
    },
    uniqueHolders: String,
    stars: String,
    uniqueHoldersChanges: String,
    desc: String,
    updated: { type: Date, default: Date.now } 
});

const Collection = db.model('Collection', CollectionSchema);

export default Collection