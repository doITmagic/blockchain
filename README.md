# blockchain
 Ethereum Development with Golang


## Local setup
### install Ganache testing purposes while developing dapps locally
- install it
```
npm install -g ganache-cli
```
- run ganache client, use the same mnemonic when starting ganache to generate the same sequence of public addresses:
``` 
ganache-cli -m "current affair tone subway about embrace skill away kick hedgehog off orchard" 
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
 ```
cd account
go run account.go 
```
