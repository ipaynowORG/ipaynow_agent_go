package main

import (
	"fmt"
	"log"

	"github.com/ipaynowORG/ipaynow_agent_go/protobuf/paymentProto"
	"github.com/ipaynowORG/ipaynow_agent_go/protobuf/reqParamProto"
	"github.com/ipaynowORG/ipaynow_agent_go/protobuf/transactionProto"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
)

func main() {

	queryDetailParams := &paymentProto.QueryDetailParams{}

	queryDetailParams.AppId = "1459846530407363"
	queryDetailParams.TransId = "CS20171106174649"
	queryDetailParams.Router = "00010000000003"

	//-----
	pInvoke := &paymentProto.PayInvoke{}
	pInvoke.Method = 1

	// 进行编码
	data, err := proto.Marshal(queryDetailParams)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	pInvoke.Params = data

	//-----

	content := &transactionProto.Content{}
	content.Nonce = "000100000000010000000000000001" + "_" + "CS20171106174649"
	datapInvoke, errpInvoke := proto.Marshal(pInvoke)
	if errpInvoke != nil {
		log.Fatal("marshaling error: ", errpInvoke)
	}
	content.Data = datapInvoke
	//-----

	transaction := &transactionProto.Transaction{}
	transaction.From = "000100000000010000000000000001"
	transaction.To = "3" //???????????
	datapcontent, errpcontent := proto.Marshal(content)
	if errpcontent != nil {
		log.Fatal("marshaling error: ", errpcontent)
	}

	transaction.Content = datapcontent

	//------

	signedTransaction := &transactionProto.SignedTransaction{}
	dataptransaction, errptransaction := proto.Marshal(transaction)
	if errptransaction != nil {
		log.Fatal("marshaling error: ", errptransaction)
	}
	signedTransaction.Transaction = dataptransaction
	signedTransaction.Crypto = 0

	pk, _ := crypto.HexToECDSA("013f81ac3ee1101b620031c00eac22ab53334c083c09fc191e05c29c9f0d26ad")
	signature, _ := crypto.Sign(crypto.Keccak256(dataptransaction), pk)
	signedTransaction.Signature = signature

	//----
	param := &reqParamProto.Param{}
	param.Method = 5
	param.Sign = signedTransaction

	datapparam, _ := proto.Marshal(param)
	fmt.Println(fmt.Sprintf("%x", datapparam))

}
