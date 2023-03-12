import Collection from "../models/collection.model";

export async function getCollections(skip, limit, mint) {
    if (mint) {
        return await Collection.find().skip(skip).limit(limit).sort({'_id': -1, 'mintDate': -1}).exec()
    }
    return await Collection.find().skip(skip).limit(limit).sort({'_id': -1}).exec()
}

export async function getCollection(id) {
    return await Collection.findById(id).exec();
}
