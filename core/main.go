package main

import (
	etherwallet "core/contracts"
	// "io/ioutil"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main(){

   err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	infura_url := os.Getenv("INFURA_URL")
	contract_address_from_env := os.Getenv("CONTRACT_ADDRESS")
	project_id := os.Getenv("PROJECT_ID")
	pass_phrase := os.Getenv("PASSWORD_PHRASE")
	file_name := os.Getenv("FILE_NAME")

	fmt.Println("INF",infura_url)
	fmt.Println("CON",contract_address_from_env)
	fmt.Println("pro",project_id)
	fmt.Println("PAS",pass_phrase)
	fmt.Println("FILE",file_name)

	if infura_url == "" || contract_address_from_env == "" || project_id == "" || pass_phrase == "" || file_name == "" {
		log.Fatalln("Env not set")
	}

	client,err:= ethclient.Dial(infura_url+project_id)
	if err!=nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contract_address := common.HexToAddress(contract_address_from_env)
	instance,err:= etherwallet.NewEtherwallet(contract_address,client)
	if err!=nil {
		log.Fatalf("Failed to load contract: %v",err)
	}
	keyfile:= "key-store/"+file_name


	file,err:= os.Open(keyfile)
	if err!=nil {
		log.Fatalf("Failed to open the key file : %v",err)
	}
	defer file.Close()
	
	auth,err:= bind.NewTransactorWithChainID(file,pass_phrase,big.NewInt(11155111))
	if err!=nil {
		log.Fatalf("Failed to create authorized transactor: %v",err)
	}

	auth.Value = big.NewInt(1e15)
	tx,err:= instance.Deposit(auth)

	if err!= nil {
		log.Fatalln("Failed to deposit:",err)
	}

	fmt.Printf("Deposit transaction sent: %s\n", tx.Hash().Hex())


	tx,err = instance.Withdraw(auth)

	if err!=nil {
		log.Fatalf("Failed to withdraw funds: %v",err)
	}

	fmt.Println("Transaction succeed: ",tx.Hash().Hex())
}