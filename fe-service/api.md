### 1. index api

request

```azure
curl --location --request GET '127.0.0.1:8888/api/v1/GetIndex'
```
response 

```json
{
  "status": true,
  "code": 2000,
  "msg": "success",
  "data": {
    "support_chain": [
      {
        "chain_id": 1,
        "chain_name": "Ethereum",
        "chain_icon": "https://etherscan.io/images/svg/brands/ethereum-original.svg",
        "api_url": "https://etherscan.io/images/svg/brands/ethereum-original.svg"
      }
    ],
    "head_data": {
      "total_nft_value": "10000",
      "total_nft_value_ratio": 0.95,
      "total_collections": "10000",
      "total_collections_ratio": 0.95,
      "total_whale": "10000",
      "total_whale_ratio": 0.95,
      "total_nft": "10000",
      "total_nft_ratio": 0.95
    },
    "hot_collection_list": [
      {
        "id": 1,
        "rank": 1,
        "image": "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
        "name": "SavourDao",
        "holder": 100,
        "whale_holder": 10,
        "suggest_level": 1,
        "volume": 1000,
        "floor_price": "10",
        "best_offer": "20",
        "shadow_score": "10"
      },
      {
        "id": 2,
        "rank": 2,
        "image": "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
        "name": "Chaineye",
        "holder": 100,
        "whale_holder": 10,
        "suggest_level": 1,
        "volume": 1000,
        "floor_price": "10",
        "best_offer": "20",
        "shadow_score": "10"
      },
      {
        "id": 3,
        "rank": 3,
        "image": "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
        "name": "问我学院",
        "holder": 100,
        "whale_holder": 10,
        "suggest_level": 1,
        "volume": 1000,
        "floor_price": "10",
        "best_offer": "20",
        "shadow_score": "10"
      }
    ],
    "live_mint_list": [
      {
        "id": 1,
        "rank": 1,
        "image": "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
        "name": "SavourDao",
        "holder": 100,
        "whale_holder": 10,
        "suggest_level": 1,
        "volume": 10,
        "mint_percent": 0.95,
        "total_mint": 10,
        "total_mint_percent": 0.95,
        "last_mint_time": "2022-12-05"
      },
      {
        "id": 2,
        "rank": 2,
        "image": "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
        "name": "Chaineye",
        "holder": 100,
        "whale_holder": 10,
        "suggest_level": 1,
        "volume": 10,
        "mint_percent": 0.95,
        "total_mint": 10,
        "total_mint_percent": 0.95,
        "last_mint_time": "2022-12-05"
      },
      {
        "id": 3,
        "rank": 3,
        "image": "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
        "name": "问我学院",
        "holder": 100,
        "whale_holder": 10,
        "suggest_level": 1,
        "volume": 10,
        "mint_percent": 0.95,
        "total_mint": 10,
        "total_mint_percent": 0.95,
        "last_mint_time": "2022-12-05"
      }
    ],
    "whale_holder_list": [
      {
        "address": "0x4675c7e5baafbffbca748158becba61ef3b0a263",
        "total_value": "100001000",
        "hold_nft_list": {
          "total_hold": 10,
          "images": [
            "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
            "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png"
          ]
        },
        "hold_collection_list": {
          "total_hold": 10,
          "images": [
            "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
            "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png"
          ]
        },
        "realize_pnl": "10",
        "label": "CZ"
      }
    ],
    "shadow_score": {
      "blue_chip": "95",
      "fluidity": "80",
      "reliability": "60",
      "community_active": "70",
      "heat": "50",
      "potential_income": "80"
    }
  }
}
```

### 1. collections api

request

```curl
curl --location --request POST 'http://193.203.215.185:8888/api/v1/GetHotCollectionList' \
--header 'Content-Type: application/json' \
--data-raw '{
    "page": 1,
    "page_size": 10,
    "order_by": 0
}'
```

response

```azure
{
    "status": true,
    "code": 2000,
    "msg": "success",
    "data": [
        {
            "id": 1,
            "rank": 1,
            "image": "",
            "name": "SavourDao",
            "holder": 100,
            "whale_holder": 1000,
            "suggest_level": 4,
            "volume": 11100,
            "floor_price": "2.5",
            "best_offer": "3",
            "shadow_score": "10"
        }
    ]
}
```