# Intract with Smart contract using Golang

## Setup
1. Install Ganache -> [Download](https://trufflesuite.com/ganache/)

2. Install  Solidity -> [Docs](https://docs.soliditylang.org/en/v0.8.2/installing-solidity.html)
```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

```
solc --optimize --abi ./contracts/MySmartContract.sol -o build
solc --optimize --bin ./contracts/MySmartContract.sol -o build
abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/MySmartContract.go

```

3. Install  Geth -> [Docs](https://geth.ethereum.org/docs/install-and-build/installing-geth)
```
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum
```

## To deply contract 
```
solc --optimize --abi ./contracts/MySmartContract.sol -o build
solc --optimize --bin ./contracts/MySmartContract.sol -o build
abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/MySmartContract.go
```

## To intract with contract run 
```
go run cmd\main.go
```

## To Check and test use postman 
Import `blockChain.postman_collection.json` 
 
## Or curl commands

1. to check balance :- `curl --location --request GET 'http://localhost:1323/balance'`
2. to check admin address :- `curl --location --request GET 'http://localhost:1323/admin'`
3. to deposite 50 abount of value from account's private key in address :- 

`curl --location --request POST 'http://localhost:1323/deposit/1000000000000000000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accountPrivateKey":"c801acad90aeed53d91027767d0df2826052c8f624bee9de90adb15f3051ea58"
}'`

4. to withdrawl 10 :- 
`curl --location --request POST 'http://localhost:1323/withdrawl/1000000000000000000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accountPrivateKey":"c801acad90aeed53d91027767d0df2826052c8f624bee9de90adb15f3051ea58"
}'`
# smartcontract
