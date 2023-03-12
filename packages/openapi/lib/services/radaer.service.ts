import Shadow from "../models/rodaer.model";

export async function getShadow() {
    return await Shadow.findOne().sort({'_id': -1}).exec()
}