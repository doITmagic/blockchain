# exemples for using go-ethereum 
 Ethereum Development with Golang


## Local setup
### install Ganache testing purposes while developing dapps locally
- install it
```
npm install -g ganache-cli
```
- run ganache client, use the same mnemonic when starting ganache to generate the same sequence of public addresses:
``` 
ganache-cli -m "ticket version capable vocal country artwork industry hat senior breeze leg great cherry abuse inch round essence random laptop organ mesh resemble trade panic" 
```
- if ganache return error `digital envelope routines::unsupported` the run: 
```
export NODE_OPTIONS=--openssl-legacy-provider
 ```
 !!! You must start Ganache before running example !!!

 ## Each dir host an example:
 ### Client connection implementation: 
```
cd client
go run client.go
```
 ### Account balance implementation, example for: 
 - balance of an account 
 - account balance at the time of that block
 - pending account balance
 - generate new wallet 
 - generate hdwallet
 ```
cd account
go run account.go 
```
