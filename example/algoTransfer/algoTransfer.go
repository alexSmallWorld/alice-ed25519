package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/sha512"
	"encoding/base32"
	"encoding/hex"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/future"
	transaction "github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/beego/beego/v2/core/logs"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		logs.Info("usage:%s 1 or 0, 1 means sign by self", os.Args[0])
		return
	}
	//const algodAddress = "https://node.algoexplorerapi.io"	//主网
	const algodAddress = "https://node.testnet.algoexplorerapi.io" //测试网
	const algodToken = ""

	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		logs.Info("Issue with creating algod client: %s", err)
		return
	}

	transfer(algodClient, os.Args[1] == "1")
}

func transfer(algodClient *algod.Client, needSign bool) {
	//设置发送地址，接收地址
	account := crypto.Account{}
	if needSign {
		pk, _ := hex.DecodeString("153c3f31ac01e2b41a043e6f2a751c8a44ebf1e9818f7f5ddd516e09f06a1118")
		account, _ = crypto.AccountFromPrivateKey(ed25519.NewKeyFromSeed(pk))
	} else {
		publicKey := readPublicKey()
		logs.Info("read publicKey:",publicKey)
		account.PublicKey, _ = hex.DecodeString(publicKey)
		copy(account.Address[:], account.PublicKey)	//algo直接使用公钥作为原始地址；
	}
	fromAddr := account.Address.String()	//真正使用的地址做了base32编码；这个编码可以反向解码出原始地址-》公钥；
	toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
	var amount uint64 = 100000
	closeToAddr := ""                                              //如果设定这个地址，就会将交易的UTXO剩余的资金全部发送到这个地址。不设置的话，就会返回给发送者地址；
	note := []byte("DevPortal - My First Transaction with Go SDK") //链上刻字

	logs.Info("My origin address: %s", hex.EncodeToString(account.Address[:]))
	logs.Info("My address: %s", fromAddr)
	logs.Info("My pk: %s", hex.EncodeToString(account.PrivateKey))
	logs.Info("My publickey: %s", hex.EncodeToString(account.PublicKey))

	//查询账户余额等信息
	accountInfo, err := algodClient.AccountInformation(fromAddr).Do(context.Background())
	if err != nil {
		logs.Info("Error getting account info: %s", err)
		return
	}
	var startingAmount uint64 = accountInfo.Amount //交易前发送者地址的余额；
	logs.Info("Account balance: %d microAlgos", accountInfo.Amount)

	//获取当前系统状态参数，比如gas price等；
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		logs.Info("Error getting suggested tx params: %s", err)
		return
	}
	//组装交易
	txn, err := transaction.MakePaymentTxn(fromAddr, toAddr, amount, note, closeToAddr, txParams)
	if err != nil {
		logs.Info("Error creating transaction: %s", err)
		return
	}

	//编码成等待签名的交易
	encodedTx := msgpack.Encode(txn)
	msgParts := [][]byte{[]byte("TX"), encodedTx}
	toBeSigned := bytes.Join(msgParts, nil)
	logs.Info("raw tx", hex.EncodeToString(toBeSigned))
	logs.Info("txn.valid block number", txn.FirstValid, txn.LastValid)
	writeRawTx2File(hex.EncodeToString(toBeSigned))

	//调用脚本拉起peer进行签名，然后拷贝签名结果到文件；
	execShell()

	//签名
	var signature []byte
	if needSign {
		signature = ed25519.Sign(account.PrivateKey, toBeSigned)
	} else {
		//var sign string
		//logs.Info("please input sign result:")
		//_, _ = fmt.Scanln(&sign)
		//signature, _ = hex.DecodeString(sign)
		sig := readSignResult()
		if sig == "" {
			logs.Info("readSignResult failed")
			return
		}
		signature, _ = hex.DecodeString(sig)
	}
	logs.Info("signature", hex.EncodeToString(signature))

	//组装成要广播的交易格式
	var s types.Signature
	n := copy(s[:], signature)
	if n != len(s) {
		return
	}
	stx := types.SignedTxn{
		Sig: s,
		Txn: txn,
	}
	if stx.Txn.Sender != account.Address {
		stx.AuthAddr = account.Address
		logs.Info("stx.AuthAddr", hex.EncodeToString(stx.AuthAddr[:]))
	}
	signedTxn := msgpack.Encode(stx)
	txidBytes := sha512.Sum512_256(toBeSigned)
	txID := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(txidBytes[:])
	//logs.Info("signedTxn", hex.EncodeToString(signedTxn))
	logs.Info("Signed txid: %s", txID)

	//txnJSON, err := json.MarshalIndent(stx.Txn, "", "\t")
	//if err != nil {
	//	fmt.Printf("Can not marshall txn data: %s\n", err)
	//}
	//fmt.Printf("Transaction information: %s\n", txnJSON)

	// 广播交易
	sendResponse, err := algodClient.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		logs.Info("failed to send transaction: %s", err)
		return
	}
	logs.Info("Submitted transaction %s", sendResponse)

	// 等待交易上链；
	confirmedTxn, err := future.WaitForConfirmation(algodClient, txID, 4, context.Background())
	if err != nil {
		logs.Info("Error waiting for confirmation on txID: %s", txID)
		return
	}
	logs.Info("Confirmed Transaction: %s in Round %d", txID, confirmedTxn.ConfirmedRound)

	//打印交易详情
	logs.Info("Amount sent: %d microAlgos", confirmedTxn.Transaction.Txn.Amount)
	logs.Info("Fee: %d microAlgos", confirmedTxn.Transaction.Txn.Fee)
	amountAndFee := uint64(confirmedTxn.Transaction.Txn.Amount + confirmedTxn.Transaction.Txn.Fee)
	logs.Info("Close to Amount: %d microAlgos", startingAmount-amountAndFee)
	logs.Info("Decoded note: %s", string(confirmedTxn.Transaction.Txn.Note))
}

func execShell() {
	command := `./copyAndSign.sh`
	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		logs.Info("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	logs.Info("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func writeRawTx2File(str string) {
	userFile := "algoTransfer/rawTx.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		logs.Info("writeRawTx2File failed", userFile, err)
		return
	}
	fout.WriteString("msg: \"" + str + "\"")
	logs.Info("writeRawTx2File success", userFile, str)
}

func readSignResult() string {
	userFile := "./signer/id-10001-ed25519-output.yaml"
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		logs.Info("readSignResult failed", userFile, err)
		return ""
	}
	br := bufio.NewReader(fin)
	line, err := br.ReadString('\n')
	if err != nil {
		logs.Info("readSignResult failed", userFile, err)
		return ""
	} else {
		logs.Info("read line:%v", line)
		return line
	}
}

func readPublicKey() string {
	userFile := "./dkg/publicKey-id-10001.txt"
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		logs.Info("readSignResult failed", userFile, err)
		return ""
	}
	br := bufio.NewReader(fin)
	line, err := br.ReadString('\n')
	if err != nil {
		logs.Info("readSignResult failed", userFile, err)
		return ""
	} else {
		logs.Info("read line:%v", line)
		return line
	}
}
