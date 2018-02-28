package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI is a command line interface for the bank
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("\n- Welcome To The Stellar CLI -\n")
	fmt.Println("  checkBalance -address PUBLIC_KEY\n")
	fmt.Println("  transaction -to PUBLIC_KEY -from PRIVATE KEY  -amount AMOUNT\n")
	fmt.Println("  trust -seed PRIVATE_KEY -issuer ISSUER_PUBLIC_KEY  -asset SYMBOL -limit AMOUNT\n")
	fmt.Println("  assetTransaction -from ISSUER_PRIVATE_KEY -to RECEIVER_PUBLIC_KEY -asset SYMBOL -amount AMOUNT\n")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 1 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run executes the CLI
func (cli *CLI) Run() {
	cli.validateArgs()

	createAccountCmd := flag.NewFlagSet("createAccount", flag.ExitOnError)
	fundAccountCmd := flag.NewFlagSet("fundAccount", flag.ExitOnError)
	checkBalanceCmd := flag.NewFlagSet("checkBalance", flag.ExitOnError)
	transactionCmd := flag.NewFlagSet("transaction", flag.ExitOnError)
	trustCmd := flag.NewFlagSet("trust", flag.ExitOnError)
	assetTransactionCmd := flag.NewFlagSet("assetTransaction", flag.ExitOnError)

	fundAddress := fundAccountCmd.String("address", "", "The address to be funded")
	balanceAddress := checkBalanceCmd.String("address", "", "The address to provide a balance for")
	transactionTo := transactionCmd.String("to", "", "Public address to send to")
	transactionFrom := transactionCmd.String("from", "", "Private key of the sender")
	transactionAmount := transactionCmd.String("amount", "", "The amount of lumens to send")
	trustFrom := trustCmd.String("seed", "", "Private key of initiating address")
	trustIssuer := trustCmd.String("issuer", "", "The public address of the issuing entity")
	trustAsset := trustCmd.String("asset", "", "Symbol of the asset to be accepted")
	trustLimit := trustCmd.String("limit", "", "The maximum amount to be held of the asset")
	assetTransactionFrom := assetTransactionCmd.String("from", "", "The private key of the issuer")
	assetTransactionTo := assetTransactionCmd.String("to", "", "The public address of the receiver")
	assetTransactionAsset := assetTransactionCmd.String("asset", "", "The symbol of the asset to send")
	assetTransactionAmount := assetTransactionCmd.String("amount", "", "The amount of the asset to send")

	switch os.Args[1] {
	case "createAccount":
		err := createAccountCmd.Parse(os.Args[1:])
		if err != nil {
			log.Panic(err)
		}
	case "fundAccount":
		err := fundAccountCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "checkBalance":
		err := checkBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "transaction":
		err := transactionCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "trust":
		err := trustCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "assetTransaction":
		err := assetTransactionCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)

	}

	if os.Args[1] == "createAccount" {
		cli.createAccountService()
	}

	if fundAccountCmd.Parsed() {
		cli.fundAccountService(*fundAddress)
	}

	if checkBalanceCmd.Parsed() {
		cli.checkBalanceService(*balanceAddress)
	}

	if transactionCmd.Parsed() {
		cli.transactionService(*transactionTo, *transactionFrom, *transactionAmount)
	}

	if trustCmd.Parsed() {
		cli.trustService(*trustFrom, *trustIssuer, *trustAsset, *trustLimit)
	}

	if assetTransactionCmd.Parsed() {
		cli.assetTransactionService(*assetTransactionFrom, *assetTransactionTo, *assetTransactionAsset, *assetTransactionAmount)
	}
}
