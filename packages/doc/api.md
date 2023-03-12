### API
#### apikey
```
header里加上 apikey
apikey: FgfW27uQdqXThfmVe-6y4Cq238e4X9-
```

#### collections
post /collections
```
[{
    id: 1
    collectionImg: './././'
    name: xxx
    chain: eth
    holders: 1000
    whaleHolers: 1000
    mint: 10000
    24hrsValues:{
        count
        percentChange
    }
    floorPrice: 200
    BestOffer: 1000
    shadowScoreChange: [1,2,3,4,5,6,7],
    shadowScore: {
        price
        whaleHold
        volume
        tnxs
    }
    totalPrice: {
        value: 1000
        ccy: ETH
        percentChange: 0.0248
    }
    totalHolders {
        value: 1000
        percentChange: 0.0248
    }
    totalWhaleHolders {
        value: 1000
        percentChange: 0.0248
    }
    totalTxns{
        value: 1.29
        ccy: eth
        percentChange: 0.0248
    }
    trading {
        price: [
            {
                price 80
                time 180
                ccy: eth
            }
        ]
        volume: [
            {
                time: 16900000
                value: 400
            }
        ]
        list: [{
            time: 16900000
            value: 800
        }]
        tnxs: [
            {
                time: 16900000,
                value: 100
            }
        ]
    }
    uniqueHolders: 1200
    stars: 1000
    uniqueHoldersChanges: 0.3
    desc: "xxx"
}]
```
#### collection
get /collections/:id
```
{
    id: 1
    collectionImg: './././'
    name: xxx
    chain: eth
    holders: 1000
    whaleHolers: 1000
    mint: 10000
    24hrsValues:{
        count
        percentChange
    }
    floorPrice: 200
    BestOffer: 1000
    shadowScoreChange: [1,2,3,4,5,6,7],
    shadowScore: {
        price
        whaleHold
        volume
        tnxs
    }
    totalPrice: {
        value: 1000
        ccy: ETH
        percentChange: 0.0248
    }
    totalHolders {
        value: 1000
        percentChange: 0.0248
    }
    totalWhaleHolders {
        value: 1000
        percentChange: 0.0248
    }
    totalTxns{
        value: 1.29
        ccy: eth
        percentChange: 0.0248
    }
    trading {
        price: [
            {
                price 80
                time 180
                ccy: eth
            }
        ]
        volume: [
            {
                time: 16900000
                value: 400
            }
        ]
        list: [{
            time: 16900000
            value: 800
        }]
        tnxs: [
            {
                time: 16900000,
                value: 100
            }
        ]
    }
    uniqueHolders: 1200
    stars: 1000
    uniqueHoldersChanges: 0.3
    desc: "xxx"
}
```

#### whale Holders top 100
post /whaleHolders
```
[{
    id
    address: '0x00'
    tokenValue{
        ccy: 'ETH',
        totalValue: 10000000000.00,
        percentChange: 0.0248
    }
    holdNfts: {
        totalValue: 1223,
        imgs: [
            '/s/d',
            '',
            '',
            ''
        ]
    }
    holdCollections: {
        totalValue: 232,
        imgs: [
            '/s/d',
            '',
            '',
            ''
        ]
    }
    realizedPnl: 2333,
    label: cz wallet
}]
```
#### hotNfts
post /nfts
```
[{
    id: 1
    name: 'xxxxxx'
    stars: 200
    whaleHolders: 200
    holders: 1000
    mints: 1000
    24hETHVolume: 1000
    24hETHVolumePerchange: 0.24
    price: 123333
    7dShadowPrice: [1,2,3,4,5,6]
    listed: 12222

}]
```

#### twitter Explore mock
post /twitter/explore
```
['tama', 'pink','dps']
```