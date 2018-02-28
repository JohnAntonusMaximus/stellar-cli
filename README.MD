# Stellar-CLI

Stellar-CLI is a command line tool built in Go that interfaces with the Stellar TestNet

# Usage
Just clone the repo, then can use the CLI with './stellar' and pass any commands listed below:
```sh
$./stellar createAccount 
$./stellar checkBalance -address PUBLIC_KEY
$./stellar transaction -to PUBLIC_KEY -from PRIVATE_KEY  -amount AMOUNT
$./stellar trust -seed PRIVATE_KEY -issuer ISSUER_PUBLIC_KEY  -asset SYMBOL -limit AMOUNT
$./stellar assetTransaction -from ISSUER_PRIVATE_KEY -to RECEIVER_PUBLIC_KEY -asset SYMBOL -amount AMOUNT
```

# Work In Progress

  - Will continue to add more functionality to this tool as I build it. If you're a Go developer, feel free to contribute as part of the Stellar community.

# Connect On LinkedIn

[LinkedIn]

# Connect On Medium

[Medium]

License
----

MIT

[LinkedIn]: <https://www.linkedin.com/in/johnradosta>
[Medium]: <https://medium.com/@jradosta>