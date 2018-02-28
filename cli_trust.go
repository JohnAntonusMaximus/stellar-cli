package main

func (cli *CLI) trustService(from, issuer, asset, limit string) {
	TrustLine(from, issuer, asset, limit)
}
