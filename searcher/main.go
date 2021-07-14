package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jannikluhn/mevsky/searcher/mevsky"
	"github.com/pkg/errors"
)

const rpcUrl = "wss://goerli.infura.io/ws/v3/cb47771bf3324acc895994de6752654b"
const contractAddressHex = "0x57badf6DED5934E90Ce9dD40BD958C1Dc58Fcd97"
const privateKeyHex = ""
const retryInterval = 5 * time.Second
const logInterval = 5 * time.Second

type Connection struct {
	Client     *ethclient.Client
	Contract   *mevsky.Mevsky
	SigningKey *ecdsa.PrivateKey
	Auth       *bind.TransactOpts
}

func NewConnection(ctx context.Context) (Connection, error) {
	log.Println("Connecting to node at URL", rpcUrl)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return Connection{}, errors.Wrap(err, "failed to connect to Ethereum node")
	}

	log.Println("Querying chain ID")
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return Connection{}, errors.Wrap(err, "failed to query chain id")
	}

	log.Println("Creating contract instance at address", contractAddressHex)
	contractAddress := common.HexToAddress(contractAddressHex)
	contract, err := mevsky.NewMevsky(contractAddress, client)
	if err != nil {
		return Connection{}, errors.Wrap(err, "faild to create contract instance")
	}

	signingKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return Connection{}, errors.Wrap(err, "failed to parse private key")
	}
	auth, err := bind.NewKeyedTransactorWithChainID(signingKey, chainID)
	if err != nil {
		return Connection{}, errors.Wrapf(err, "failed to create transactor")
	}

	return Connection{
		Client:     client,
		Contract:   contract,
		SigningKey: signingKey,
		Auth:       auth,
	}, nil
}

func main() {
	ctx := context.Background()

	for {
		conn, err := NewConnection(ctx)
		if err != nil {
			log.Println("Error:", err)
			time.Sleep(retryInterval)
			continue
		}

		err = turnOff(ctx, conn)
		if err != nil {
			log.Println("Error:", err)
			time.Sleep(retryInterval)
			continue
		}

		err = eventLoop(ctx, conn)
		if err != nil {
			log.Println("Error:", err)
			time.Sleep(retryInterval)
			continue
		}
	}
}

func eventLoop(ctx context.Context, conn Connection) error {
	watchOpts := bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}
	turnedOnChannel := make(chan *mevsky.MevskyTurnedOn)

	log.Println("Subscribing to TurnedOn events")
	sub, err := conn.Contract.MevskyFilterer.WatchTurnedOn(&watchOpts, turnedOnChannel)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to TurnedOn events")
	}

	logTicker := time.NewTicker(logInterval)
	for {
		select {
		case err := <-sub.Err():
			return errors.Wrap(err, "event subscription error")
		case ev := <-turnedOnChannel:
			log.Printf("Received TurnedOn event in block %d triggered by tx %s", ev.Raw.BlockNumber, ev.Raw.TxHash)
			err := turnOff(ctx, conn)
			if err != nil {
				return err
			}
		case <-logTicker.C:
			log.Printf("Listening for events...")
		}
	}
}

func turnOff(ctx context.Context, conn Connection) error {
	log.Println("Check if contract is on")
	callOpts := bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	isOn, err := conn.Contract.On(&callOpts)
	if err != nil {
		return errors.Wrap(err, "failed to check if contract is on")
	}
	if !isOn {
		log.Println("Contract is off, no action is taken")
		return nil
	}

	log.Println("Turning contract off")
	tx, err := conn.Contract.MevskyTransactor.TurnOff(conn.Auth, crypto.PubkeyToAddress(conn.SigningKey.PublicKey))
	if err != nil {
		return errors.Wrap(err, "failed to create TurnOff transaction")
	}
	log.Printf("Turn off transaction with hash %s sent", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, conn.Client, tx)
	if err != nil {
		return errors.Wrapf(err, "error while waiting for receipt for transaction %s", tx.Hash().Hex())
	}
	if receipt.Status != 1 {
		return errors.Errorf("Transaction %s failed", tx.Hash().Hex())
	}
	log.Println("Transaction successful")

	return nil
}
