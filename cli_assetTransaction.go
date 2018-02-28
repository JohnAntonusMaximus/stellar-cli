package main

func (cli *CLI) assetTransactionService(from, to, asset, amount string) {
	AssetTransaction(from, to, asset, amount)
}
