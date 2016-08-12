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
// X Create a signed transaction spending from it
// X Manually send that transaction to a testnet node
// - Do the handshake wire protocol

func main() {
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), []byte("nehanarula123456"))
	chainParams := &chaincfg.TestNet3Params
	addr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubKey.SerializeCompressed()), chainParams)
	_ = addr
	if err != nil {
		log.Fatal(err)
	}
	// addr: mpYJ4Uj4GRDxt9xNk1Z4FtZBZcYRbhQDQk

	// got bitcoin from testnet faucet
	txid := "e9eb1dc1961a1dfcb0ed16a43efed331dd19e985bf74a5ca16ef8a6c24f09d92"
	pay_me := "76a91462fa0c3c58c7d618994808e7789d7912cc8e40e388ac"
	pay_tadge := "001443aac20a116e09ea4f7914be1c55e4c17aa600b7"

	// create transaction spending it
	h, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		panic(err)
	}
	outpoint := wire.NewOutPoint(h, 0)
	ti := wire.NewTxIn(outpoint, nil)
	mytx := wire.NewMsgTx()
	mytx.AddTxIn(ti)

	b1, err := hex.DecodeString(pay_me)
	if err != nil {
		panic(err)
	}
	to1 := wire.NewTxOut(1520000000, b1)
	mytx.AddTxOut(to1)
	b2, err := hex.DecodeString(pay_tadge)
	if err != nil {
		panic(err)
	}
	to2 := wire.NewTxOut(100000000, b2)
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
	fmt.Printf("%x\n", buf.Bytes())
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
