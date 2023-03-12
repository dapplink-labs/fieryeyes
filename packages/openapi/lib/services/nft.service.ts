import Nft from "../models/nft.model";

export async function getNfts(skip, limit) {
    return await Nft.find().skip(skip).limit(limit).sort({'_id':-1}).exec()
}

export async function getNft(id) {
    return await Nft.findById(id).exec();
}
