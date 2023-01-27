# Open api for nft service

### 1. GetMainTokenPrice

request example
```json
{
    "token_id": "111",
    "page": 1,
    "page_size": 100
}
```

response
```json
{
    "status": true,
    "code": 2000,
    "msg": "get main token price success",
    "data": [
        {
            "main_token_name": "BTC",
            "usd_price": "12000",
            "cny_price": "80018",
            "date_time": "2020-11-11"
        },
        {
            "main_token_name": "ETH",
            "usd_price": "1400",
            "cny_price": "12000",
            "date_time": "2020-11-11"
        }
    ]
}
```

curl 
```
curl --location --request GET '127.0.0.1:8888/openapi/v1/GetMainTokenPrice' \
--header 'Content-Type: application/json' \
--data-raw '{
    "token_id": "111",
    "page": 1,
    "page_size": 100
}'
```

### 2. GetAddressInfo

request example

```json
{
    "address_id": 1,
    "daily_page": 1,
    "daily_page_size": 100
}
```

response 

```json
{
    "status": true,
    "code": 2000,
    "msg": "get address info success",
    "data": {
        "id": 1,
        "address": "0x4E5B2e1dc63F6b91cb6Cd759936495434C7e972F",
        "label": "FixedFloat",
        "is_giant_whale": 1,
        "balance": "1000000000000000000",
        "token_value": "100000000000",
        "nft_value": "10000000",
        "address_daily_list": [
            {
                "address_id": 1,
                "balance": "10000",
                "token_value": "1",
                "nft_value": "1",
                "date_time": "2022-01-10"
            },
            {
                "address_id": 1,
                "balance": "10000000000",
                "token_value": "10000000000",
                "nft_value": "10000000000",
                "date_time": "2022-01-11"
            }
        ]
    }
}
```

curl
```
curl --location --request POST '127.0.0.1:8888/openapi/v1/getAddressInfo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "address_id": 1,
    "daily_page": 1,
    "daily_page_size": 100
}'
```

### 3. getNftCollectionsInfo

request example

```json
{
  "token_address": "0xD9E24CF4328df7CC34966C058CBAFeAe744beC1c",
  "daily_page": 1,
  "daily_page_size": 100
}
```

response

```json
{
  "status": true,
  "code": 2000,
  "msg": "get address info success",
  "data": {
    "name": "savourlabs",
    "address": "0xD9E24CF4328df7CC34966C058CBAFeAe744beC1c",
    "introduce": "savourlabs nft",
    "total_holder": 1000,
    "average_holder": 0,
    "total_giant_whale_holder": 10000,
    "average_giant_whale_holder": 10000,
    "total_txn": 10000,
    "average_txn": 10000,
    "average_price": "100",
    "total_price": "1000",
    "suggest_level": 3,
    "collection_daily": [
      {
        "total_holder": 1000,
        "average_holder": 0,
        "total_giant_whale_holder": 1000,
        "average_giant_whale_holder": 10000,
        "total_txn": 1000,
        "average_txn": 10000,
        "average_price": "1000",
        "total_price": "1000",
        "date_time": "2022-01-10"
      }
    ]
  }
}
```

curl
```
curl --location --request POST '127.0.0.1:8888/openapi/v1/getNftCollectionsInfo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "token_address": "0xD9E24CF4328df7CC34966C058CBAFeAe744beC1c",
    "daily_page": 1,
    "daily_page_size": 100
}'
```

### 4. GetNftInfo

request example
```json
{
    "token_id": "111",
    "page": 1,
    "page_size": 100
}
```

response

```json
{
    "status": true,
    "code": 2000,
    "msg": "get nft info success",
    "data": {
        "id": 1,
        "address": "0x00000000000000000000",
        "token_id": "111",
        "token_url": "https://github.com/savour-labs/fieryeyes",
        "total_txn": 10,
        "total_holder": 110,
        "total_giant_whale_holder": 120,
        "latest_price": "100",
        "suggest_level": 2,
        "nft_daily": [
            {
                "nft_id": 1,
                "total_txn": 10,
                "total_holder": 10,
                "total_giant_whale_holder": 10,
                "latest_price": "1111",
                "date_time": "2022-01-22"
            }
        ],
        "current_holder": [
            {
                "address_id": 2,
                "address": "0x4E5B2e1dc63F6b91cb6Cd759936495434C7e9723",
                "label": "Savourlabs"
            }
        ],
        "historical_holder": [
            {
                "address_id": 1,
                "address": "0x4E5B2e1dc63F6b91cb6Cd759936495434C7e972F",
                "label": "FixedFloat"
            }
        ]
    }
}
```

curl 

```
curl --location --request POST '127.0.0.1:8888/openapi/v1/getNftInfo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "token_id": "111",
    "page": 1,
    "page_size": 100
}'
```

