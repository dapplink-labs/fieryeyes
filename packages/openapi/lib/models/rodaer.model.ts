
import mongoose from 'mongoose';
import config from "../../config/index";
const db = mongoose.createConnection(config.mongoHost);
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
