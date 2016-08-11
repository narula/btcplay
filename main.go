package main

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"log"
)

func main() {
	chainParams := &chaincfg.TestNet3Params

	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), []byte("nehanarula123456"))
	_ = privKey
	addr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubKey.SerializeCompressed()), chainParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v %v %v\n", pubKey, addr, addr.ScriptAddress())
	// addr: mpYJ4Uj4GRDxt9xNk1Z4FtZBZcYRbhQDQk
	// txid: e9eb1dc1961a1dfcb0ed16a43efed331dd19e985bf74a5ca16ef8a6c24f09d92

	// TODO
	// X Make a bitcoin address
	// X Fund it
	// - Create a signed transaction spending from it
	// - Send that transaction to a testnet node

	mytx := wire.NewMsgTx()
	_ = mytx
	//	ti := wire.NewTxIn(nil)
	//	mytx.AddTxIn(ti)

}

func network_crap() {
	// Use Tadge's testnet server
	host := "lit3.co:18333"
	_ = host

	//	conn, err := net.Dial("tcp", host)
	//	if err != nil {
	//		panic(err)
	//	}

	//	n, err := wire.WriteMessageN(conn, msg, 0, chainParams.Net)
	//	if err != nil {
	//		panic(err)
	//	}
}
