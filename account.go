package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
)

// CreateAccount creates a new account
func CreateAccount() {
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAccount Keys Created - This does NOT create an account on the network. Please fund your account to create it.\n")
	fmt.Printf("\nPrivate: %s\nPublic: %s\n\n", pair.Seed(), pair.Address())
}

// BalanceCheck checks the balance of an account
func BalanceCheck(address string) {
	client := horizon.DefaultTestNetClient
	account, err := client.LoadAccount(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n## Account Balance - %s ##\n\n", address)
	for _, balance := range account.Balances {
		fmt.Printf("Asset: %s\nBalance: %s\n\n", balance.Asset.Code, balance.Balance)
	}
}

// FundAccount funds an account
func FundAccount(address string) {
	resp, err := http.Get("https://horizon-testnet.stellar.org/friendbot?addr=" + address)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
