package main

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jannikluhn/mevsky/client/mevsky"
	"github.com/pkg/errors"
)

const retryInterval = 5 * time.Second
const logInterval = 5 * time.Second

func main() {
	ctx := context.Background()

	connectionConfig := mevsky.ConnectionConfig{
		RpcUrl:             "wss://rpc.xdaichain.com/wss",
		ContractAddressHex: "0xeCbb670751ED019b021A062E30dfdf22Bc3fe8d9",
		SigningKeyHex:      "",
	}

	for {
		conn, err := mevsky.NewConnection(ctx, connectionConfig)
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

func eventLoop(ctx context.Context, conn mevsky.Connection) error {
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

func turnOff(ctx context.Context, conn mevsky.Connection) error {
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
