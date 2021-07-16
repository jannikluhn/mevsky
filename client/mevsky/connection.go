package mevsky

import (
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type ConnectionConfig struct {
	RpcUrl             string
	ContractAddressHex string
	SigningKeyHex      string
}

type Connection struct {
	Client     *ethclient.Client
	Contract   *Mevsky
	SigningKey *ecdsa.PrivateKey
	Auth       *bind.TransactOpts
}

func NewConnection(ctx context.Context, config ConnectionConfig) (Connection, error) {
	log.Println("Connecting to node at URL", config.RpcUrl)
	client, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		return Connection{}, errors.Wrap(err, "failed to connect to Ethereum node")
	}

	log.Println("Querying chain ID")
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return Connection{}, errors.Wrap(err, "failed to query chain id")
	}

	log.Println("Creating contract instance at address", config.ContractAddressHex)
	contractAddress := common.HexToAddress(config.ContractAddressHex)
	contract, err := NewMevsky(contractAddress, client)
	if err != nil {
		return Connection{}, errors.Wrap(err, "faild to create contract instance")
	}

	var auth *bind.TransactOpts
	var signingKey *ecdsa.PrivateKey
	if config.SigningKeyHex != "" {
		signingKey, err := crypto.HexToECDSA(config.SigningKeyHex)
		if err != nil {
			return Connection{}, errors.Wrap(err, "failed to parse signing key")
		}
		auth, err = bind.NewKeyedTransactorWithChainID(signingKey, chainID)
		if err != nil {
			return Connection{}, errors.Wrapf(err, "failed to create transactor")
		}
	} else {
		auth = nil
		signingKey = nil
	}

	return Connection{
		Client:     client,
		Contract:   contract,
		SigningKey: signingKey,
		Auth:       auth,
	}, nil
}
