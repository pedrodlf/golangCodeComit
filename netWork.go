package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ethnetwork = os.Getenv("ETHNETWORK")
var keyFile = os.Getenv("KEYFILE")
var pass = os.Getenv("PASS")
var folderPath = os.Getenv("FOLDERPATH")
var scAddress = os.Getenv("SCADDRESS")

//AddPoints crate a new meeting
func AddPoints(user common.Address, meeting uint32, points uint32, action uint32) (int64, error) {

	log.Print("En AddBreackPointToChain ")
	log.Print(ethnetwork)
	log.Print(scAddress)
	log.Print(keyFile)
	log.Print(pass)
	conn, err := Unlock()
	scConn, err := getSCConnector(conn)

	auth, err := getAuth(scConn, keyFile, pass)
	if err != nil {

		return 0, err
	}
	log.Print("auth vale")
	log.Print(auth)
	result, err := addPointsToUserAndMeeting(scConn, auth, conn, meeting, user, points, action)
	if err != nil {

		return 0, err
	}
	return result, nil
}
func MakePayment(user common.Address, meeting uint32, points uint32, action uint32) (int64, error) {

	log.Print("En AddBreackPointToChain ")
	log.Print(ethnetwork)
	log.Print(scAddress)
	log.Print(keyFile)
	log.Print(pass)
	conn, err := Unlock()
	scConn, err := getSCConnector(conn)

	auth, err := getAuth(scConn, keyFile, pass)
	if err != nil {

		return 0, err
	}
	log.Print("auth vale")
	log.Print(auth)
	result, err := MakePaymentOnChain(scConn, auth, conn, meeting, user, points, action)
	if err != nil {

		return 0, err
	}
	return result, nil
}

func MakePaymentOnChain(scConn *Token, auth *bind.TransactOpts, conn *ethclient.Client, meeting uint32, user common.Address, points uint32, acction uint32) (int64, error) {

	transactor, err := scConn.PayPrize(auth, meeting, user, acction, points)
	log.Printf("ScConn-------------->: %v", scConn)
	log.Printf("auth---------------->: %v", auth)
	log.Printf("Ethclient.Client---->: %v", conn)
	log.Printf("MeetingID----------->: %v", meeting)
	log.Printf("User common.Address->: %v", user)
	log.Printf("Points-------------->: %v", points)

	if err != nil {
		log.Print("error en AddUserAndProduct")
		log.Fatalf("Failed to AddUserAndProduct: %v", err)
		return 0, err
	}
	log.Print("AddUserAndProduct")
	log.Print("transactor vale")
	log.Print(transactor)
	log.Print("transactor CheckNonce")
	log.Print(transactor.CheckNonce())
	log.Print("transactor data")
	log.Print(transactor.Data())
	log.Print("transactor value")
	log.Print(transactor.Value())
	log.Print("COST")
	log.Print(transactor.Cost())
	log.Print("hash")
	log.Print(transactor.Hash().Hex())
	ctx := context.Background()
	//ctx := auth.Context
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	for {
		receipt, err := conn.TransactionReceipt(ctx, transactor.Hash())

		if receipt != nil {

			log.Print(receipt)
			if receipt.Logs != nil {
				log.Print("receip logs != nil")
				lgs := receipt.Logs
				n := len(lgs)
				log.Print("TxHash")
				m := receipt.TxHash.Hex()
				log.Print(m)
				log.Print("logs size")
				log.Print(n)

			} else {
				log.Print("receipt logs not recovered")
			}
			log.Print("Receipt Status")
			log.Print(receipt.Status)
			log.Print("SC receipt")
			var scresult int64
			for k, v := range receipt.Logs {

				for l, w := range v.Topics {
					log.Print("receipt Topics")
					fmt.Printf("%#v -> %s\n", l, w.Hex())

					fmt.Printf("%#v -> %s\n", l, w.Big())
					d, _ := strconv.ParseInt(w.Hex(), 0, 64)
					fmt.Println(d)

				}
				a := v.Topics[len(v.Topics)-1]
				scresult, _ = strconv.ParseInt(a.Hex(), 0, 64)
				log.Printf("RESULTADO: %d\n", scresult)
				fmt.Printf("%#v -> %#v\n", k, v.BlockNumber)
			}
			log.Print("result")
			log.Print(scresult)
			return scresult, nil
		}
		if err != nil {
			log.Print("error al analizar la transaccion")
			log.Printf("Failed geting receipt: %v", err)
			//return err
		} else {
			log.Print("Transaction not yet mined")
		}

		select {
		case <-ctx.Done():
			log.Print("Transaccion Realizada")
			return 0, nil

		case <-queryTicker.C:
			log.Println("Pendiente")
		}

	}

}

// obtenemos el Authorization

func addPointsToUserAndMeeting(scConn *Token, auth *bind.TransactOpts, conn *ethclient.Client, meeting uint32, user common.Address, points uint32, acction uint32) (int64, error) {

	transactor, err := scConn.AddPoints(auth, meeting, user, acction, points)
	log.Printf("ScConn-------------->: %v", scConn)
	log.Printf("auth---------------->: %v", auth)
	log.Printf("Ethclient.Client---->: %v", conn)
	log.Printf("MeetingID----------->: %v", meeting)
	log.Printf("User common.Address->: %v", user)
	log.Printf("Points-------------->: %v", points)

	if err != nil {
		log.Print("error en AddUserAndProduct")
		log.Fatalf("Failed to AddUserAndProduct: %v", err)
		return 0, err
	}
	log.Print("AddUserAndProduct")
	log.Print("transactor vale")
	log.Print(transactor)
	log.Print("transactor CheckNonce")
	log.Print(transactor.CheckNonce())
	log.Print("transactor data")
	log.Print(transactor.Data())
	log.Print("transactor value")
	log.Print(transactor.Value())
	log.Print("COST")
	log.Print(transactor.Cost())
	log.Print("hash")
	log.Print(transactor.Hash().Hex())
	ctx := context.Background()
	//ctx := auth.Context
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	for {
		receipt, err := conn.TransactionReceipt(ctx, transactor.Hash())

		if receipt != nil {

			log.Print(receipt)
			if receipt.Logs != nil {
				log.Print("receip logs != nil")
				lgs := receipt.Logs
				n := len(lgs)
				log.Print("TxHash")
				m := receipt.TxHash.Hex()
				log.Print(m)
				log.Print("logs size")
				log.Print(n)

			} else {
				log.Print("receipt logs not recovered")
			}
			log.Print("Receipt Status")
			log.Print(receipt.Status)
			log.Print("SC receipt")
			var scresult int64
			for k, v := range receipt.Logs {

				for l, w := range v.Topics {
					log.Print("receipt Topics")
					fmt.Printf("%#v -> %s\n", l, w.Hex())

					fmt.Printf("%#v -> %s\n", l, w.Big())
					d, _ := strconv.ParseInt(w.Hex(), 0, 64)
					fmt.Println(d)

				}
				a := v.Topics[len(v.Topics)-1]
				scresult, _ = strconv.ParseInt(a.Hex(), 0, 64)
				log.Printf("RESULTADO: %d\n", scresult)
				fmt.Printf("%#v -> %#v\n", k, v.BlockNumber)
			}
			log.Print("result")
			log.Print(scresult)
			return scresult, nil
		}
		if err != nil {
			log.Print("error al analizar la transaccion")
			log.Printf("Failed geting receipt: %v", err)
			//return err
		} else {
			log.Print("Transaction not yet mined")
		}

		select {
		case <-ctx.Done():
			log.Print("Transaccion Realizada")
			return 0, nil

		case <-queryTicker.C:
			log.Println("Pendiente")
		}

	}

}

//Unlock OBTENEMOS  la conexion con la red de trabajo
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

func getSCConnector(conn *ethclient.Client) (*Token, error) {
	log.Print("NewToken")
	token, err := NewToken(common.HexToAddress(scAddress), conn)
	log.Print("token value")
	log.Print(token)

	if err != nil {
		log.Print("error en NewToken")
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
		return nil, err
	}
	return token, nil
}

// obtenemos el Authorization
func getAuth(scConn *Token, wallet string, pass string) (*bind.TransactOpts, error) {

	ctx := context.Background()

	log.Print("En el auth")
	auth, err := bind.NewTransactor(strings.NewReader(wallet), pass)
	auth.GasLimit = 300000
	auth.Context = ctx
	log.Print("Address")
	log.Print(auth.From.Hex())
	log.Print("Signer")
	log.Print(auth.Signer)
	log.Print("Nonce (nil = use pending state)")
	log.Print(auth.Nonce)
	if err != nil {
		log.Print("error en auth")
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, err
	}
	return auth, nil
}
