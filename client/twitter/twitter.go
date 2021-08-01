package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jannikluhn/mevsky/client/mevsky"
	"github.com/pkg/errors"
)

const retryInterval = 5 * time.Second
const logInterval = 5 * time.Second

const consumerKey = ""
const consumerSecret = ""
const oauthToken = ""
const oauthSecret = ""

func main() {
	ctx := context.Background()

	connectionConfig := mevsky.ConnectionConfig{
		RpcUrl:             "wss://goerli.infura.io/ws/v3/cb47771bf3324acc895994de6752654b",
		ContractAddressHex: "0x284f6A78BB277716B6397fEd2Efe9D2FEfb8731e",
		SigningKeyHex:      "",
	}

	for {
		config := oauth1.NewConfig(consumerKey, consumerSecret)
		token := oauth1.NewToken(oauthToken, oauthSecret)
		httpClient := config.Client(oauth1.NoContext, token)
		client := twitter.NewClient(httpClient)

		conn, err := mevsky.NewConnection(ctx, connectionConfig)
		if err != nil {
			log.Println("Error:", err)
			time.Sleep(retryInterval)
			continue
		}

		err = eventLoop(ctx, conn, client)
		if err != nil {
			log.Println("Error:", err)
			time.Sleep(retryInterval)
			continue
		}
	}
}

func eventLoop(ctx context.Context, conn mevsky.Connection, client *twitter.Client) error {
	watchOpts := bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}
	turnedOnChannel := make(chan *mevsky.MevskyTurnedOn)
	turnedOffChannel := make(chan *mevsky.MevskyTurnedOff)

	log.Println("Subscribing to TurnedOn events")
	onSub, err := conn.Contract.MevskyFilterer.WatchTurnedOn(&watchOpts, turnedOnChannel)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to TurnedOn events")
	}
	log.Println("Subscribing to TurnedOff events")
	offSub, err := conn.Contract.MevskyFilterer.WatchTurnedOff(&watchOpts, turnedOffChannel)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to TurnedOff events")
	}

	logTicker := time.NewTicker(logInterval)
	for {
		select {
		case err := <-onSub.Err():
			return errors.Wrap(err, "event subscription error for TurnedOn events")
		case err := <-offSub.Err():
			return errors.Wrap(err, "event subscription error for TurnedOff events")
		case ev := <-turnedOnChannel:
			log.Printf("Received TurnedOn event in block %d triggered by tx %s", ev.Raw.BlockNumber, ev.Raw.TxHash)
			err := tweet(client, makeTurnOnMessage(ev))
			if err != nil {
				return err
			}
		case ev := <-turnedOffChannel:
			log.Printf("Received TurnedOff event in block %d triggered by tx %s", ev.Raw.BlockNumber, ev.Raw.TxHash)
			err := tweet(client, makeTurnOffMessage(ev))
			if err != nil {
				return err
			}
		case <-logTicker.C:
			log.Printf("Listening for events...")
		}
	}
}

func tweet(client *twitter.Client, message string) error {
	log.Println("Sending tweet")
	_, _, err := client.Statuses.Update(message, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to send tweet")
	}
	return nil
}

func makeTurnOnMessage(ev *mevsky.MevskyTurnedOn) string {
	return fmt.Sprintf(
		"✨ Account %s just switched me on! %s",
		ev.Sender.Hex(),
		ev.Raw.BlockNumber,
		makeBlockscoutUrl(ev.Raw.TxHash),
	)
}

func makeTurnOffMessage(ev *mevsky.MevskyTurnedOff) string {
	return fmt.Sprintf(
		"😞 Account %s just switched me off 💤 %s",
		ev.Sender.Hex(),
		ev.Raw.BlockNumber,
		makeBlockscoutUrl(ev.Raw.TxHash),
	)
}

func makeBlockscoutUrl(txHash common.Hash) string {
	return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tx/%s", txHash.Hex())
}
