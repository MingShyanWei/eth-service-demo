
# 


## Module
- API Server
- RPC2DB
- DB Migration


## TODO Action
- API function
    - Transcation of blobk
    - Log of Transcation
- RPC to DB based on job-queue
- DB connection pool
- Error handle


## How to start

1. Env


| Key  | Description  | Example |
| - | - | - |
| DB_CONNECTION | MySQL connection | root:kzy0RV0lte@tcp(192.168.17.104:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local |


2. DB Migration

```
cd ./migration
go run migration.go
```

3. Eth RPC to DB

```
cd ./tools/rpc2db
go run rpc2db.go
```

4. API Service

```
go run main.go
```


## API

1. List last n blocks
[GET]
/v1/blocks?limit=n

2. Get the specified block detail with transcation hash
[GET]
/v1/blocks/:num

3. Get the specified transcation detail with event log
[GET]
/v1/transcation/:txHash



## Docker

1. Make api docker image
```
docker build -t eth-service-demo-api -f Docker/Dockerfile .
```


## Kubernetes

Deploymeny yaml in ./Kubernetes/*


## Demo

1. /v1/blocks?limit=n
```
curl --location --request GET 'https://eth-service-demo-api.stepnhub.com/v1/blocks?limit=10'
```

2. /v1/blocks/:num
```
curl --location --request GET 'https://eth-service-demo-api.stepnhub.com/v1/blocks/20597939'
```


3. /v1/transcation/:txHash
```
curl --location --request GET 'https://eth-service-demo-api.stepnhub.com/v1/transaction/0xdc33485c58067aae6a5704955e6a040d9a7fc81c7ee2a4bb208c7e342fd0426d'
```