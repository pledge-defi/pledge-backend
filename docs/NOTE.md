The commands of creating go files from ABI and BIN files

    abigen --abi abi/v21/pledge_pool.abi --bin bin/v21/pledge_pool.bin --pkg tokengo --out
    pledgePoolToken.go

The go token files of the main network and the test network need to be generated

    pledge_pool_token           //get pool information 
    bsc_pledge_oracle_token     //get token price

/etc/systemd/system.conf 



abigen --abi abi/bridge.json --pkg bridge --out bridge.go
