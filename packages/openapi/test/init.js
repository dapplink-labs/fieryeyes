import Collection from "../lib/models/collection.model";
import Chart from "../lib/models/chart.model";
import Nft from "../lib/models/nft.model";
import Shadow from "../lib/models/rodaer.model";
import Holder from "../lib/models/whaleHolders.model";


async function run() {
  const collectionData = {}
  const chartData = {}
  const nftData = {}
  const ShadowData = {}
  const holderData = {}

  const collection = new Collection()
  const chart = new Chart()
  const nft = new Nft()
  const shadow = new Shadow()
  const holder = new Holder()

  await collection.save(collectionData);
  await chart.save(chartData)
  await nft.save(nftData)
  await shadow.save(ShadowData)
  await holder.save(holderData)
}

run();


