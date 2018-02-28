package main

import (
	"fmt"
	"log"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func Transaction(to string, fromSeed string, amount string) {
	client := horizon.DefaultTestNetClient

	// Returns a *Transaction.Builder
	tx, err := build.Transaction(
		build.SourceAccount{fromSeed},
		build.AutoSequence{client},
		build.TestNetwork,
		build.Payment(
			build.Destination{to},
			build.NativeAmount{amount},
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Returns a TransactionEnvelopeBuilder
	txe, err := tx.Sign(fromSeed)
	if err != nil {
		log.Fatal(err)
	}

	// Returns xdr-then-base64 encoded string
	txeB64, err := txe.Base64()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Tx base64: %s", txeB64)
	resp, err := client.SubmitTransaction(txeB64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response: ", resp)
}

func AssetTransaction(_fromSeed string, _to string, asset string, amount string) {
	client := horizon.DefaultTestNetClient

	assetAmount := build.CreditAmount{asset, _fromSeed, amount}
	// Returns a *Transaction.Builder
	tx, err := build.Transaction(
		build.SourceAccount{_fromSeed},
		build.AutoSequence{client},
		build.TestNetwork,
		build.Payment(
			build.Destination{_to},
			assetAmount,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Returns a TransactionEnvelopeBuilder
	txe, err := tx.Sign(_fromSeed)
	if err != nil {
		log.Fatal(err)
	}

	// Returns xdr-then-base64 encoded string
	txeB64, err := txe.Base64()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Tx base64: %s", txeB64)
	resp, err := client.SubmitTransaction(txeB64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response: ", resp)
}

// TrustLine creates a line of trust from a receiver account to a particular asset by a specific anchor
func TrustLine(_fromSeed string, _issuerPublicAdr, asset string, limit string) {
	client := horizon.DefaultTestNetClient
	tx, err := build.Transaction(
		build.SourceAccount{_fromSeed},
		build.AutoSequence{client},
		build.TestNetwork,
		build.Trust(asset, _issuerPublicAdr, build.Limit(limit)),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Returns a TransactionEnvelopeBuilder
	txe, err := tx.Sign(_fromSeed)
	if err != nil {
		log.Fatal(err)
	}

	// Returns xdr-then-base64 encoded string
	txeB64, err := txe.Base64()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Tx base64: %s", txeB64)
	resp, err := client.SubmitTransaction(txeB64)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	fmt.Println("Response: ", resp)
}
