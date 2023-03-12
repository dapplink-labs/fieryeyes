
import mongoose from 'mongoose';
const db = mongoose.createConnection('mongodb://127.0.0.1:27017/nft');
const { Schema } = mongoose;

const ShadowSchema = new Schema({
    id: Number,
    fluidity: String,
    blueChipDegree: String,
    potentialIncome: String,
    heat: String,
    activeCommunity: String,
    Reliability: String,
    updated: { type: Date, default: Date.now } 
})

const Shadow = db.model('Shadow', ShadowSchema);

export default Shadow
