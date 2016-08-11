package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"log"
)

// TODO
// X Make a bitcoin address
// X Fund it
// - Create a signed transaction spending from it
// - Send that transaction to a testnet node

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
	// pay me script: 76a91462fa0c3c58c7d618994808e7789d7912cc8e40e388ac
	txid := "e9eb1dc1961a1dfcb0ed16a43efed331dd19e985bf74a5ca16ef8a6c24f09d92"

	h, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		panic(err)
	}

	outpoint := wire.NewOutPoint(h, 0)
	ti := wire.NewTxIn(outpoint, nil)
	mytx := wire.NewMsgTx()
	mytx.AddTxIn(ti)
	b1, err := hex.DecodeString("76a91462fa0c3c58c7d618994808e7789d7912cc8e40e388ac")

	if err != nil {
		panic(err)
	}

	to1 := wire.NewTxOut(1520000000, b1)
	b2, err := hex.DecodeString("001443aac20a116e09ea4f7914be1c55e4c17aa600b7")

	if err != nil {
		panic(err)
	}

	to2 := wire.NewTxOut(100000000, b2)
	mytx.AddTxOut(to1)
	mytx.AddTxOut(to2)
	sig, err := txscript.SignatureScript(mytx, 0, b1, txscript.SigHashAll, privKey, true)
	if err != nil {
		panic(err)
	}
	mytx.TxIn[0].SignatureScript = sig
	var buf bytes.Buffer
	err = mytx.Serialize(&buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x \n\n%#v\n", buf.Bytes(), mytx)
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
