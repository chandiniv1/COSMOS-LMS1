# LEAVE MANAGEMENT SYSTEM USING COSMOS-SDK

This module is built based on the cosmos-sdk

### REQUIREMENTS
* Golang
* GRPC
* Cobra-Cli
* Docker
* Protobuffers
* Cosmos-SDK
* Local Testnet Setup
* ***
### AVAILABLE TRANSACTIONS IN THIS MODULE
* Register Admin
* Add Student
* Apply Leave
* Accept Leave
* ***
### AVAILABLE QUERIES IN THIS MODULE
* Get Admin
* Get Student
* Get Students
* Get Leave Request
* Get Leave Requests
* Get Status By ID
* Get Leave Approves 
*  ***
### SET UP
#### INSTALL GOLANG
* Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local,     creating a fresh Go tree in /usr/local/go: 
  ```
  $ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
  ```
* Add /usr/local/go/bin to the PATH environment variable:
  ```
  $ export PATH=$PATH:/usr/local/go/bin
  ```
* Verify the go version:
  ```
  $ go version
  ```
#### INSTALL GRPC & PROTOCOL BUFFER
```
   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
   
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
#### UPDATE THE PATH
```
   export PATH="$PATH:$(go env GOPATH)/bin"
```
#### INSTALL COBRA-CLI
```
   go install github.com/spf13/cobra-cli@latest
```
#### INSTALL COSMOS-SDK
```
   go get github.com/cosmos/cosmos-sdk
```
#### INSTALL DOCKER
   Reference [how to install docker engine](https://docs.docker.com/engine/install/ubuntu/)
   
#### SET UP LOCAL TESTNET
*  Setup Cosmos SDK and create binaries
   ```
   go get github.com/cosmos/cosmos-sdk

   cd $GOPATH/src/github.com/cosmos/cosmos-sdk

   make build-simd #this creates a blockchain binary for testing

   cd build

   ./simd -h #lists all the commands available
   ```
*  Create some keys (blockchain accounts)
   ```
   ./simd keys add validator-key
   
   ./simd keys add mykey1
   
   ./simd keys add mykey2
   
   #list all keys
   ./simd keys list
   ```
*  Setup a local testnet

   * Initialize the testnet with chain-id and some validator name
   ```
   ./simd init --chain-id testnet myvalidator
   ```
   
   * Add some genesis accounts and tokens
   ```
   ./simd add-genesis-account validator-key 1000000000stake
   
   ./simd add-genesis-account $(./simd keys show mykey1 -a) 10000000000stake
   ```
   
   * Create gentx
   ```
   ./simd gentx validator-key --chain-id testnet
   ```
   
   * Collect gentxs
   ```
   ./simd collect-gentxs
   ```
   
   * Start the node
   ```
   ./simd start
   ```
   
***
### CLIENT

User can interact with ```lms``` module using ```cli```

### TRANSACTIONS

We can see the available transactions using:
```
   ./lmsd tx lms --help
```
#### REGISTER ADMIN

This command is used to register the admin to the store
```
   ./lmsd tx lms register-admin [address] [name]
```
##### EXAMPLE
```
   ./lmsd tx lms register-admin cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5 chandini --from validator-key --chain-id testnet
```
***
#### Add Student

This command is used to add the student to the store
```
   ./lmsd tx lms add-student [address] [name] [id]
```
##### EXAMPLE
```
   ./lmsd tx lms add-student cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukmycqn priya 00317 --from validator-key --chain-id testnet
```
***
#### Apply Leave

This command is used to apply the leave by the student
```
   ./lmsd tx lms apply-leave [address] [reason] [leaveid] [from] [to]
```
##### EXAMPLE
```
   ./lmsd tx lms apply-leave cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukmycqn cold 00317 2023-03-10 2023-03-11  --from validator-key --chain-id testnet
```
***
#### ACCEPT LEAVE

This command is used to either accept or reject the leave
```
   ./lmsd tx lms accept-leave [leaveid] [status]
```
##### EXAMPLE
```
   ./lmsd tx lms accept-leave 00317 2 --from validator-key --chain-id testnet
```
***
### QUERIES

We can see the available queries using :
```
   ./lmsd query lms --help
```
#### GET STUDENT

This query is used to get the student by address
```
  ./lmsd query lms get-student [flag]
 ```
##### EXAMPLE
```
  ./lmsd query lms get-student cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukmycqn
```
##### OUTPUT
```
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   address:cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukmycqn
   name:priya
   id:00317
```
***
#### Get ADMIN

This query is used to get the admin by using address

```
  ./lmsd query lms get-admin [flags]
```
##### EXAMPLE
```
  ./lmsd query lms get-admin cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
```
##### OUTPUT
```
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   name:chandini
```
***
#### GET LEAVE REQUEST

This query is used to get the leave request of the student by taking leaveID
```
  ./lmsd query lms get-leave-request [flag]
 ```
##### EXAMPLE
```
  ./lmsd query lms get-leave-request 00317
```
##### OUTPUT
```
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   address:cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukmycqn
   reason:cold
   leaveID:00317
   from: "2023-03-19T00:00:00Z"
   to: "2023-03-24T00:00:00Z"
   status:STATUS_PENDING
```
***
#### GET STATUS

This query is used to get the leave status of the student by taking admin address and leaveID
```
  ./lmsd query lms get-status [flag] [flag]
 ```
##### EXAMPLE
```
  ./lmsd query lms get-status cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5 00317
```
##### OUTPUT
```
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   leaveID:00317
   status:STATUS_ACCEPTED
```
***
#### GET STUDENTS

This query is used to get the students
```
  ./lmsd query lms get-students
 ```
##### OUTPUT
```
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   address:cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukmycqn
   name:priya
   id:00317
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   address:cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukqwert
   name:anu
   id:00313
```
***
#### GET LEAVE REQUESTS

This query is used to get the leave requests of the students
```
  ./lmsd query lms get-leave-requests
 ```
##### OUTPUT
```
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   address:cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukmycqn
   reason:cold
   leaveID:00317
   from: "2023-03-19T00:00:00Z"
   to: "2023-03-24T00:00:00Z"
   status:STATUS_ACCEPTED
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   address:cosmos1et0unt0p2ya5934mwyg9cfxmjcjvf4hukqwert
   reason:cough
   leaveID:00313
   from: "2023-03-20T00:00:00Z"
   to: "2023-03-22T00:00:00Z"
   status:STATUS_REJECTED
  
```
***
#### GET LEAVE APPROVES

This query is used to get the leave status of all the students
```
  ./lmsd query lms get-leave-approves
 ```
##### OUTPUT
```
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   leaveID:00317
   status:STATUS_ACCEPTED
  -admin:cosmos1uh5rv23rsrgn9lcn5f503cl4qdxh0jkqv8q3s5
   leaveID:00313
   status:STATUS_REJECTED
```

