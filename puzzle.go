package main

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Puzzle interface {
	Execute(rpcClient *ethclient.Client, chainId *big.Int, privateKey *ecdsa.PrivateKey, accountAddress common.Address) error
}
