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

### 2. collections list api

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

```json
{
  "status": true,
  "code": 2000,
  "msg": "success",
  "data": [
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
  ]
}
```

### 3. collect detail

request 

``` curl
curl --location --request POST 'http://127.0.0.1:8888/api/v1/GetHotCollectionDetail' \
--header 'Content-Type: application/json' \
--data-raw '{
    "collection_id": 1,
    "page": 1,
    "page_size": 10
}'
```

response

```json
{
    "status": true,
    "code": 2000,
    "msg": "success",
    "data": {
        "id": 1,
        "name": "SavourDao",
        "image": "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
        "creator": "0x4675C7e5BaAFBFFbca748158bEcBA61ef3b0a263",
        "holder": 100,
        "chain": "Ethereum",
        "introduce": "Savour labs nft for user",
        "shadow_score": {
            "blue_chip": "95",
            "fluidity": "80",
            "reliability": "60",
            "community_active": "70",
            "heat": "50",
            "potential_income": "80"
        },
        "trading_list": [
            {
                "stat_time": "2023-03-15",
                "price": "300.00"
            },
            {
                "stat_time": "2023-03-16",
                "price": "400.00"
            },
            {
                "stat_time": "2023-03-16",
                "price": "400.00"
            },
            {
                "stat_time": "2023-03-15",
                "price": "300.00"
            }
        ],
        "volume_list": [
            {
                "stat_time": "2023-03-15",
                "volume": 1000
            },
            {
                "stat_time": "2023-03-16",
                "volume": 1300
            },
            {
                "stat_time": "2023-03-16",
                "volume": 1300
            },
            {
                "stat_time": "2023-03-15",
                "volume": 1000
            }
        ],
        "list_list": [
            {
                "stat_time": "2023-03-15",
                "price_dis": "150.00"
            },
            {
                "stat_time": "2023-03-16",
                "price_dis": "250.00"
            },
            {
                "stat_time": "2023-03-16",
                "price_dis": "250.00"
            },
            {
                "stat_time": "2023-03-15",
                "price_dis": "150.00"
            }
        ],
        "floor_price_list": [
            {
                "stat_time": "2023-03-15",
                "floor_price": "100.00",
                "best_offer": "300.00"
            },
            {
                "stat_time": "2023-03-16",
                "floor_price": "200.00",
                "best_offer": "500.00"
            },
            {
                "stat_time": "2023-03-16",
                "floor_price": "200.00",
                "best_offer": "500.00"
            },
            {
                "stat_time": "2023-03-15",
                "floor_price": "100.00",
                "best_offer": "300.00"
            }
        ],
        "whale_holder": [
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
        ]
    }
}
```

### 4. get live mint nft

request

```curl
curl --location --request POST 'http://193.203.215.185:8888/api/v1/GetLiveMintList' \
--data-raw ''
```

response

```json
{
    "status": true,
    "code": 2000,
    "msg": "get live mint success",
    "data": [
        {
            "id": 1,
            "rank": 1,
            "image": "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png",
            "name": "SavourDao",
            "holder": 100,
            "whale_holder": 10,
            "suggest_level": 1,
            "volume": 10,
            "mint_percent": 0.98,
            "total_mint": 10,
            "total_mint_percent": 0.98,
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
            "mint_percent": 0.98,
            "total_mint": 10,
            "total_mint_percent": 0.98,
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
            "mint_percent": 0.98,
            "total_mint": 10,
            "total_mint_percent": 0.98,
            "last_mint_time": "2022-12-05"
        }
    ]
}
```

### 5. get nft by collection

request

```curl
curl --location --request POST 'http://193.203.215.185:8888/api/v1/GetNftByCollectionId' \
--header 'Content-Type: application/json' \
--data-raw '{
    "collect_id": 1,
    "page": 1,
    "page_size": 10
}'
```

response

```json
{
    "status": true,
    "code": 2000,
    "msg": "get nft list success",
    "data": [
        {
            "id": 1,
            "image": "https://solana.nftscan.com/static/no-preview.11455274.png",
            "name": "savour labs seed nft",
            "chain": "Ethereum",
            "holder": 1000,
            "hold_label": "cz",
            "price": "1000.00",
            "usd_price": "7000.00"
        }
    ]
}
```

### 6. get nft detail

request

```curl
{
    "nft_id": 1,
    "type": 1,
    "page": 1,
    "page_size": 10
}
```

response

```json
{
    "status": true,
    "code": 2000,
    "msg": "get nft detail success",
    "data": {
        "id": 1,
        "image": "https://solana.nftscan.com/static/no-preview.11455274.png",
        "name": "savour labs seed nft",
        "chain": "Ethereum",
        "contract": "0xa9d1e08c7793af67e9d92fe308d5697fb81d3e43",
        "creator": "0x4675c7e5baafbffbca748158becba61ef3b0a263",
        "token_url": "http://savour.group",
        "toke_id": "#11",
        "describe": "savour labs logo",
        "mint_hash": "0x668dd633ecdd9ce7df373edef28f48455d1a712c6000f7dedb05b2509f577b93",
        "mint_time": "2022-10-11",
        "holder": 1000,
        "whale_holder": 100,
        "price": "1000.00",
        "usd_price": "7000.00",
        "total_txn": 100,
        "nft_txn": [
            {
                "from_addr": "0x4675c7e5baafbffbca748158becba61ef3b0a263",
                "to_addr": "0x4675c7e5baafbffbca748158becba61ef3b0a261",
                "type": 1,
                "price": "100",
                "trade_time": "2022-10-11"
            },
            {
                "from_addr": "0x4675c7e5baafbffbc1248158becba61ef3b0a263",
                "to_addr": "0x4675c7e5baafbffbca748158becba61ef3b0a123",
                "type": 1,
                "price": "100",
                "trade_time": "2022-10-11"
            }
        ]
    }
}
```