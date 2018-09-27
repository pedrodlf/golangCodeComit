package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ethnetwork = os.Getenv("ETHNETWORK")
var keyFile = os.Getenv("KEYFILE")
var pass = os.Getenv("PASS")
var folderPath = os.Getenv("FOLDERPATH")

//AccountUser clave de acceso para la cuenta del usuario
type AccountUser struct {
	key string `json:"key"`
}

//OBTENEMOS  la conexion con la red de trabajo
func Unlock() (*ethclient.Client, error) {
	// Create an IPC based RPC connection to a remote node
	//conn, err := ethclient.Dial("/home/karalabe/.ethereum/testnet/geth.ipc")

	conn, err := ethclient.Dial(ethnetwork) //publica 52.31.219.254 - antes ws
	if err != nil {

		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}

	return conn, nil
}

//CreateAccount  nueva cuenta
func CreateAccount(id uint32) (string, string, string, error) {
	os.MkdirAll(folderPath, os.ModePerm)
	ks := keystore.NewKeyStore(
		folderPath,
		keystore.LightScryptN,
		keystore.LightScryptP)
	tfwAcount, err := ks.NewAccount(pass)
	if err != nil {
		log.Fatalf("Failed to creat TFWAccount : %v", err)
		return "", "", "", err
	}
	//adrs := string(tfwAcount.Address.Hex())
	adrs := string(tfwAcount.Address.String())
	log.Printf("tfwAcount.Address---> : %v", tfwAcount.Address)
	log.Printf("tfwAcount.Address---> : %v", adrs)
	log.Printf("tfwAcount.URL---> : %v", tfwAcount.URL)
	dat, err := ioutil.ReadFile(tfwAcount.URL.Path)
	check(err)

	log.Printf("Key---> : %v", string(dat))

	return string(dat), adrs, tfwAcount.URL.Path, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
