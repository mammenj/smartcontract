package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/mammenj/smartcontract/api" // generated smart contract bindings in go

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// address of etherum env
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		e.Logger.Printf("Error in getting client: %v", err)
		panic(err)
	}

	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(client, "c801acad90aeed53d91027767d0df2826052c8f624bee9de90adb15f3051ea58", 0)

	//deploying smart contract
	address, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		e.Logger.Printf("Error in deploging: %v", err)
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx
	fmt.Println("instance->", instance)
	fmt.Println("tx->", tx.Hash().Hex())

	//creating api object to intract with smart contract function
	conn, err := api.NewApi(common.HexToAddress(address.Hex()), client)
	if err != nil {
		e.Logger.Printf("Error in creating smart contract instance: %v", err)
		panic(err)
	}



	e.GET("/balance", func(c echo.Context) error {
		reply, err := conn.Balance(&bind.CallOpts{})
		if err != nil {
			e.Logger.Printf("Error in getting balance: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, reply)
	})
	e.GET("/cbalance", func(c echo.Context) error {
		reply, err := conn.ContractBalance(&bind.CallOpts{})
		if err != nil {
			e.Logger.Printf("Error in getting balance of smart contract: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/admin", func(c echo.Context) error {
		reply, err := conn.Admin(&bind.CallOpts{})
		if err != nil {
			e.Logger.Printf("Error in getting Admin of smart contract: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, reply)
	})
	e.POST("/deposit/:amount", func(c echo.Context) error {
		amount := c.Param("amount")
		amt, _ := strconv.Atoi(amount)

		//gets address of account by which amount to be deposite
		var v map[string]interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			e.Logger.Printf("Error in JSON Decode in deposit: %v", err)
			panic(err)
		}

		//creating auth object for above account
		auth := getAccountAuth(client, v["accountPrivateKey"].(string), int64(amt))
		reply, err := conn.Deposit(auth)
		if err != nil {
			e.Logger.Printf("Error in deposit transaction details: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, reply)
	})
	e.POST("/withdrawl/:amount", func(c echo.Context) error {
		amount := c.Param("amount")
		amt, _ := strconv.Atoi(amount)

		var v map[string]interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			e.Logger.Printf("Error in withdraw json decode: %v", err)
			panic(err)
		}

		auth := getAccountAuth(client, v["accountPrivateKey"].(string), 0)
		// auth.Nonce.Add(auth.Nonce, big.NewInt(int64(1))) //it is use to create next nounce of account if it has to make another transaction

		reply, err := conn.Withdrawl(auth, big.NewInt(int64(amt)))
		if err != nil {
			e.Logger.Printf("Error in withdrawing: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, reply)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// function to create auth for any account from its private key
func getAccountAuth(client *ethclient.Client, privateKeyAddress string, msg_value int64) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		log.Printf("Error in private key: %v", err)
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Error in getting nonce: %v", err)
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("Error in getting chainID: %v", err)
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Error in getting transaction: %v", err)
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(msg_value) // in wei
	auth.GasLimit = uint64(3000000)    // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
