### 1. index api

request

```azure
curl --location --request GET '127.0.0.1:8888/api/v1/GetIndex'
```
response 

```azure
{
    "status": true,
    "code": 2000,
    "msg": "success",
    "data": {
        "support_chain": [
            {
                "chain_id": 1,
                "chain_name": "Ethereum",
                "chain_icon": "http://localhost/media/a.img",
                "api_url": "http://localhost:8080"
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
        ],
        "live_mint_list": [
            {
                "id": 1,
                "rank": 1,
                "image": "",
                "name": "SavourDao",
                "holder": 100,
                "whale_holder": 1000,
                "suggest_level": 4,
                "volume": 1000,
                "mint_percent": 0.95,
                "total_mint": 1000,
                "total_mint_percent": 0.95,
                "last_mint_time": "2023-03-12 17:000"
            }
        ],
        "whale_holder_list": [
            {
                "address": "0x000000000000000000",
                "total_value": "100000000010000000",
                "hold_nft_list": null,
                "hold_collection_list": null,
                "realize_pnl": "10",
                "label": "Cz wallet"
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

```azure
curl --location --request POST '127.0.0.1:8888/api/v1/GetHotCollectionList' \
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