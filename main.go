package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var (
	FlagPuzzle = flag.String("puzzle", "", "specify the puzzle to run")

	Puzzles = map[string]Puzzle{
		"fallback":   &FallbackPuzzle{},
		"fallout":    &FalloutPuzzle{},
		"coinflip":   &CoinflipPuzzle{},
		"delegation": &DelegationPuzzle{},
		"vault":      &VaultPuzzle{},
	}
)

func init() {
	godotenv.Load()

	flag.Parse()
}

func Transactor(chainId *big.Int) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		return nil, err
	}

	return bind.NewKeyedTransactorWithChainID(privateKey, chainId)
}

func main() {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		log.Fatalln(err)
	}

	rpcClient, err := ethclient.Dial("https://sepolia.infura.io/v3/c6b15721b1044ab7a30d3b911f535e47")

	if err != nil {
		log.Fatalln(err)
	}

	chainId, err := rpcClient.ChainID(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	Puzzles[*FlagPuzzle].Execute(rpcClient, chainId, privateKey, crypto.PubkeyToAddress(privateKey.PublicKey))
}
