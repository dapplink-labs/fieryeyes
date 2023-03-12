import Holder from "../models/whaleHolders.model";

export async function getHolders(skip, limit) {
    return await Holder.find().skip(skip).limit(limit).sort({'_id':-1}).exec()
}

export async function getHolder(id) {
    return await Holder.findById(id).exec();
}
