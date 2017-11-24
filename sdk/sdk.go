package sdk

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ipaynowORG/ipaynow_agent_go/protobuf/paramProto"
	"github.com/ipaynowORG/ipaynow_agent_go/protobuf/paymentProto"
	"github.com/ipaynowORG/ipaynow_agent_go/protobuf/reqParamProto"
	"github.com/ipaynowORG/ipaynow_agent_go/protobuf/transactionProto"
	"github.com/ipaynowORG/ipaynow_agent_go/util"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
)

type SdkParams struct {
	reqUrl     string
	key        string
	merchantId string
}

func MakeSdkParams(reqUrl string, key string, merchantId string) *SdkParams {
	var sdkParams = new(SdkParams)
	sdkParams.key = key
	sdkParams.reqUrl = reqUrl
	sdkParams.merchantId = merchantId
	return sdkParams
}

type QueryParams struct {
	appId   string
	transId string
}

func MakeQueryParams(appId string, transId string) *QueryParams {
	var queryParams = new(QueryParams)
	queryParams.appId = appId
	queryParams.transId = transId
	return queryParams
}

type QueryResult struct {
	AppId          string
	TransStatus    string
	DeviceType     string
	NowPayOrderNo  string
	ResponseMsg    string
	ResponseTime   string
	ResponseCode   string
	PayChannelType string
	MhtOrderNo     string
	Funcode        string
	MhtOrderAmt    string
	TransType      string
}

type AgentPayParams struct {
	mhtOrderAmt      uint64
	appId            string
	agentPayMemo     string
	mhtReqTime       string
	mhtOrderNo       string
	accType          string
	payeeName        string
	payeeCardNo      string
	payeeCardUnionNo string
}

func MakeAgentPayParams(mhtOrderAmt uint64, appId string, agentPayMemo string, mhtReqTime string, mhtOrderNo string, accType string, payeeName string, payeeCardNo string, payeeCardUnionNo string) *AgentPayParams {
	var agentPayParams = new(AgentPayParams)
	agentPayParams.appId = appId
	agentPayParams.mhtOrderAmt = mhtOrderAmt
	agentPayParams.agentPayMemo = agentPayMemo
	agentPayParams.mhtReqTime = mhtReqTime
	agentPayParams.mhtOrderNo = mhtOrderNo
	agentPayParams.accType = accType
	agentPayParams.payeeName = payeeName
	agentPayParams.payeeCardNo = payeeCardNo
	agentPayParams.payeeCardUnionNo = payeeCardUnionNo
	return agentPayParams
}

type AgentPayResult struct {
	AppId         string
	ResponseMsg   string
	ResponseCode  string
	MhtOrderNo    string
	NowPayOrderNo string
	TransStatus   string
}

type AgentReceiveParams struct {
	appId         string
	mhtOrderAmt   uint64
	mhtOrderNo    string
	mhtReqTime    string
	mhtUserId     string
	mhtUserCardId string
	cardNo        string
	cardOwner     string
	cardType      string
	cardIdenType  string
	cardIdenNo    string
	cardPhoneNo   string
	agentPayMemo  string
	accType       string
	notifyUrl     string
}

func MakeAgentReceiveParams(appId string, mhtOrderAmt uint64, mhtOrderNo string, mhtReqTime string, mhtUserId string, mhtUserCardId string, cardNo string, cardOwner string, cardType string, cardIdenType string, cardIdenNo string, cardPhoneNo string, agentPayMemo string, accType string, notifyUrl string) *AgentReceiveParams {
	var agentReceiveParams = new(AgentReceiveParams)
	agentReceiveParams.appId = appId
	agentReceiveParams.mhtOrderAmt = mhtOrderAmt
	agentReceiveParams.mhtOrderNo = mhtOrderNo
	agentReceiveParams.mhtReqTime = mhtReqTime
	agentReceiveParams.mhtUserId = mhtUserId
	agentReceiveParams.mhtUserCardId = mhtUserCardId
	agentReceiveParams.cardNo = cardNo
	agentReceiveParams.cardOwner = cardOwner
	agentReceiveParams.cardType = cardType
	agentReceiveParams.cardIdenType = cardIdenType
	agentReceiveParams.cardIdenNo = cardIdenNo
	agentReceiveParams.cardPhoneNo = cardPhoneNo
	agentReceiveParams.agentPayMemo = agentPayMemo
	agentReceiveParams.accType = accType
	agentReceiveParams.notifyUrl = notifyUrl
	return agentReceiveParams
}

type AgentReceiveResult struct {
	AppId        string
	ResponseMsg  string
	ResponseCode string
	MhtOrderNo   string
}

type AgentPayRefundQueryParams struct {
	appId   string
	transId string
}

func MakeAgentPayRefundQueryParams(appId string, transId string) *AgentPayRefundQueryParams {
	var agentPayRefundQueryParams = new(AgentPayRefundQueryParams)
	agentPayRefundQueryParams.appId = appId
	agentPayRefundQueryParams.transId = transId
	return agentPayRefundQueryParams
}

type AgentPayRefundQueryResult struct {
	AppId                 string
	ResponseMsg           string
	ResponseCode          string
	MhtOrderNo            string
	NowPayOrderNo         string
	OriginalMhtOrderNo    string
	OriginalNowPayOrderNo string
	PayeeAccType          int
	PayeeName             string
	PayeeCardNo           string
	PayeeCardUnionNo      string
	RefundDate            string
	RefundCode            string
	RefundMsg             string
	TransStatus           string
}

type BatchQueryParams struct {
	appId      string
	transId    string
	nowPage    uint64
	pageSize   uint64
	refundDate string
}

func MakeBatchQueryParams(appId string, transId string, refundDate string, nowPage uint64, pageSize uint64) *BatchQueryParams {
	var batchQueryParams = new(BatchQueryParams)
	batchQueryParams.appId = appId
	batchQueryParams.transId = transId
	batchQueryParams.refundDate = refundDate
	batchQueryParams.nowPage = nowPage
	batchQueryParams.pageSize = pageSize
	return batchQueryParams
}

type BatchQueryResult struct {
	AppId                            string
	ResponseMsg                      string
	ResponseCode                     string
	MhtOrderNo                       string
	NowPage                          int
	PageSize                         int
	CountNums                        int
	CountPages                       int
	agentPayRefundQueryRespFieldList []AgentPayRefundQueryRespField
}

type AgentPayRefundQueryRespField struct {
	NowPayOrderNo         string
	OriginalMhtOrderNo    string
	OriginalNowPayOrderNo string
	PayeeAccType          string
	PayeeName             string
	PayeeCardNo           string
	PayeeCardUnionNo      string
	RefundDate            string
	RefundCode            string
	RefundMsg             string
	TransStatus           string
}

type BalanceQueryParams struct {
	appId   string
	transId string
}

func MakeBalanceQueryParams(appId string, transId string) *BalanceQueryParams {
	var balanceQueryParams = new(BalanceQueryParams)
	balanceQueryParams.appId = appId
	balanceQueryParams.transId = transId
	return balanceQueryParams
}

type BalanceQueryResp struct {
	ResponseMsg    string
	ResponseCode   string
	ResponseTime   string
	MhtOrderNo     string
	AccountBalance uint64
}

/**
 * 账户余额查询
 * @param balanceQueryParams,
 * 	balanceQueryParams.appId 商户应用唯一标识
 * 	balanceQueryParams.mhtOrderNo 请求流水号
 * @param sdkParams,
 * 	sdkParams.reqUrl 请求地址
 * 	sdkParams.key 请求流水号 商户生成的私钥(公钥发由现在支付)
 * 	sdkParams.merchantId 商户号,由现在支付分配
 * @return BalanceQueryRespDto
 *  BalanceQueryResp.appId,同输入
 *  BalanceQueryResp.mhtOrderNo,同输入
 *  BalanceQueryResp.accountBalance, 账户余额
 *  BalanceQueryResp.responseCode 响应码(见文档附录)
 *  BalanceQueryResp.responseMsg 响应信息
 */
func BalanceQuery(sdkParams *SdkParams, balanceQueryParams *BalanceQueryParams) *BalanceQueryResp {
	queryDetailParams := &paymentProto.QueryDetailParams{}

	queryDetailParams.AppId = balanceQueryParams.appId
	queryDetailParams.TransId = balanceQueryParams.transId
	queryDetailParamsdata, _ := proto.Marshal(queryDetailParams)

	pInvoke := &paymentProto.PayInvoke{}
	pInvoke.Method = 1
	pInvoke.Params = queryDetailParamsdata
	pInvokedata, _ := proto.Marshal(pInvoke)

	content := &transactionProto.Content{}
	content.Nonce = sdkParams.merchantId + "_" + balanceQueryParams.transId
	content.Data = pInvokedata
	contentdata, _ := proto.Marshal(content)

	transaction := &transactionProto.Transaction{}
	transaction.From = sdkParams.merchantId
	transaction.To = "3" //
	transaction.Content = contentdata
	transactiondata, _ := proto.Marshal(transaction)

	signedTransaction := &transactionProto.SignedTransaction{}

	signedTransaction.Transaction = transactiondata
	signedTransaction.Crypto = 0

	pk, _ := crypto.HexToECDSA(sdkParams.key)
	signature, _ := crypto.Sign(crypto.Keccak256(transactiondata), pk)
	signedTransaction.Signature = signature

	//----
	param := &reqParamProto.Param{}
	param.Method = 6
	param.Sign = signedTransaction

	datapparam, _ := proto.Marshal(param)

	result, _ := util.Post(sdkParams.reqUrl, datapparam)
	//	fmt.Println(fmt.Sprintf("%x", result))

	// 进行解码
	resultParam := &paramProto.Param{}
	err := proto.Unmarshal(result, resultParam)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	balanceQueryResp := &BalanceQueryResp{}

	var err1 = json.Unmarshal([]byte(resultParam.Message), &balanceQueryResp)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		return balanceQueryResp
	}

	return nil
}

/**
 * 代付退票批量查询
 * @param batchQueryParams
 *  batchQueryParams.appId 商户应用唯一标识
 *  batchQueryParams.mhtOrderNo 商户请求流水号
 *  batchQueryParams.nowPage 商户请求页码
 *  batchQueryParams.pageSize 商户请求数据量
 *  batchQueryParams.refundDate 退票日期
 * @param sdkParams,
 * 	sdkParams.reqUrl 请求地址
 * 	sdkParams.key 请求流水号 商户生成的私钥(公钥发由现在支付)
 * 	sdkParams.merchantId 商户号,由现在支付分配
 * @return BatchQueryResult
 *  BatchQueryResult.responseCode 响应码(见文档附录)
 *  BatchQueryResult.responseMsg 响应信息
 *  BatchQueryResult.mhtOrderNo,同输入
 *  BatchQueryResult.appId,同输入
 *  BatchQueryResult.nowPage,同输入
 *  BatchQueryResult.pageSize,同输入
 *  BatchQueryResult.countNums,数据总数
 *  BatchQueryResult.countPages,总页数
 *  BatchQueryResult.agentPayRefundQueryRespFieldList 退票详情
 */
func AgentPayRefundBatchQuery(sdkParams *SdkParams, batchQueryParams *BatchQueryParams) *BatchQueryResult {

	queryAgentPayRefundParams := &paymentProto.QueryAgentPayRefundParams{}

	queryAgentPayRefundParams.AppId = batchQueryParams.appId
	queryAgentPayRefundParams.TransId = batchQueryParams.transId
	queryAgentPayRefundParams.NowPage = batchQueryParams.nowPage
	queryAgentPayRefundParams.PageSize = batchQueryParams.pageSize
	queryAgentPayRefundParams.RefundDate = batchQueryParams.refundDate
	queryAgentPayRefundParamsdata, _ := proto.Marshal(queryAgentPayRefundParams)

	pInvoke := &paymentProto.PayInvoke{}
	pInvoke.Method = 1
	pInvoke.Params = queryAgentPayRefundParamsdata
	pInvokedata, _ := proto.Marshal(pInvoke)

	content := &transactionProto.Content{}
	content.Nonce = sdkParams.merchantId + "_" + batchQueryParams.transId
	content.Data = pInvokedata
	contentdata, _ := proto.Marshal(content)

	transaction := &transactionProto.Transaction{}
	transaction.From = sdkParams.merchantId
	transaction.To = "3" //
	transaction.Content = contentdata

	signedTransaction := &transactionProto.SignedTransaction{}
	dataptransaction, errptransaction := proto.Marshal(transaction)
	if errptransaction != nil {
		log.Fatal("marshaling error: ", errptransaction)
	}
	signedTransaction.Transaction = dataptransaction
	signedTransaction.Crypto = 0

	pk, _ := crypto.HexToECDSA(sdkParams.key)
	signature, _ := crypto.Sign(crypto.Keccak256(dataptransaction), pk)
	signedTransaction.Signature = signature

	param := &reqParamProto.Param{}
	param.Method = 8
	param.Sign = signedTransaction

	datapparam, _ := proto.Marshal(param)

	result, _ := util.Post(sdkParams.reqUrl, datapparam)
	//	fmt.Println(fmt.Sprintf("%x", result))

	// 进行解码
	resultParam := &paramProto.Param{}
	err := proto.Unmarshal(result, resultParam)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	batchQueryResult := &BatchQueryResult{}
	var err1 = json.Unmarshal([]byte(resultParam.Message), &batchQueryResult)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		return batchQueryResult
	}

	return nil
}

/**
 * 单笔代付退票查询
 * @param agentPayRefundQueryParams
 * agentPayRefundQueryParams.appId 商户应用唯一标识
 * agentPayRefundQueryParams.mhtOrderNo 商户请求流水号
 * @param sdkParams,
 * 	sdkParams.reqUrl 请求地址
 * 	sdkParams.key 请求流水号 商户生成的私钥(公钥发由现在支付)
 * 	sdkParams.merchantId 商户号,由现在支付分配
 * @return AgentPayRefundQuery
 * AgentPayRefundQueryResult.responseCode 响应码(见文档附录)
 * AgentPayRefundQueryResult.responseMsg 响应信息
 * AgentPayRefundQueryResult.mhtOrderNo,同输入
 * AgentPayRefundQueryResult.appId,同输入
 * AgentPayRefundQueryResult.nowPayOrderNo , 现在支付流水号,退票流水号
 * AgentPayRefundQueryResult.originalMhtOrderNo ,
 * AgentPayRefundQueryResult.originalNowPayOrderNo , 原代付流水号
 * AgentPayRefundQueryResult.payeeAccType ,入账账户类型,01 对公  02 对私
 * AgentPayRefundQueryResult.payeeName, 入账账户户名
 * AgentPayRefundQueryResult.payeeCardNo, 入账账户账号
 * AgentPayRefundQueryResult.payeeCardUnionNo,入账账户联行号
 * AgentPayRefundQueryResult.refundDate,退票日期,yyyyMMdd
 * AgentPayRefundQueryResult.refundCode,退票码
 * AgentPayRefundQueryResult.refundMsg,退票原因
 * AgentPayRefundQueryResult.transStatus,交易状态。SUCCESSED,成功 。FAILED,失败。PROCESSING,处理中。
 */
func AgentPayRefundQuery(sdkParams *SdkParams, agentPayRefundQueryParams *AgentPayRefundQueryParams) *AgentPayRefundQueryResult {
	queryAgentPayRefundParams := &paymentProto.QueryAgentPayRefundParams{}

	queryAgentPayRefundParams.AppId = agentPayRefundQueryParams.appId
	queryAgentPayRefundParams.TransId = agentPayRefundQueryParams.transId
	queryAgentPayRefundParamsdata, _ := proto.Marshal(queryAgentPayRefundParams)

	pInvoke := &paymentProto.PayInvoke{}
	pInvoke.Method = 1
	pInvoke.Params = queryAgentPayRefundParamsdata
	pInvokedata, _ := proto.Marshal(pInvoke)

	content := &transactionProto.Content{}
	content.Nonce = sdkParams.merchantId + "_" + agentPayRefundQueryParams.transId
	content.Data = pInvokedata
	contentdata, _ := proto.Marshal(content)

	transaction := &transactionProto.Transaction{}
	transaction.From = sdkParams.merchantId
	transaction.To = "3" //
	transaction.Content = contentdata

	signedTransaction := &transactionProto.SignedTransaction{}
	dataptransaction, errptransaction := proto.Marshal(transaction)
	if errptransaction != nil {
		log.Fatal("marshaling error: ", errptransaction)
	}
	signedTransaction.Transaction = dataptransaction
	signedTransaction.Crypto = 0

	pk, _ := crypto.HexToECDSA(sdkParams.key)
	signature, _ := crypto.Sign(crypto.Keccak256(dataptransaction), pk)
	signedTransaction.Signature = signature

	param := &reqParamProto.Param{}
	param.Method = 7
	param.Sign = signedTransaction

	datapparam, _ := proto.Marshal(param)

	result, _ := util.Post(sdkParams.reqUrl, datapparam)
	//	fmt.Println(fmt.Sprintf("%x", result))

	// 进行解码
	resultParam := &paramProto.Param{}
	err := proto.Unmarshal(result, resultParam)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	agentPayRefundQueryResult := &AgentPayRefundQueryResult{}
	var err1 = json.Unmarshal([]byte(resultParam.Message), &agentPayRefundQueryResult)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		return agentPayRefundQueryResult
	}

	return nil
}

/**
 * 单笔代收
 * 	@param agentReceiveParams
 *  agentReceiveParams.appId : 商户应用唯一标识
 *  agentReceiveParams.agentPayMemo : 资金用途,代收原因，显示在银行流水中
 *  agentReceiveParams.mhtUserId :商户客户唯一标识,持卡人在商户这边的唯一标识，用于获取持卡人的多张卡信息
 *  agentReceiveParams.mhtUserCardId :商户白名单ID,商户自己记录的持卡人卡信息ID，因为一个人可以由多张卡
 *  agentReceiveParams.cardOwner :持卡人姓名
 *  agentReceiveParams.cardType :银行卡类别,0001借记卡
 *  agentReceiveParams.cardNo :持卡人银行卡号
 *  agentReceiveParams.cardIdenType :持卡人证件类型,0001身份证
 *  agentReceiveParams.cardIdenNo :持卡人证件号
 *  agentReceiveParams.cardPhoneNo :持卡人预留手机号
 *  agentReceiveParams.accType :账户类型。0对私 1对公
 *  agentReceiveParams.notifyUrl :商户通知地址
 *  agentReceiveParams.mhtReqTime :商户请求时间
 *  agentReceiveParams.mhtOrderNo :商户请求流水号
 *  agentReceiveParams.mhtOrderAmt :代付金额
 * @param sdkParams,
 * 	sdkParams.reqUrl 请求地址
 * 	sdkParams.key 请求流水号 商户生成的私钥(公钥发由现在支付)
 * 	sdkParams.merchantId 商户号,由现在支付分配
 * 	@return AgentReceiveResult
 *  AgentReceiveResult.AppId,同输入
 *  AgentReceiveResult.mhtOrderNo,同输入
 *  AgentReceiveResult.responseCode 响应码(见文档附录)
 *  AgentReceiveResult.responseMsg 响应信息
 */
func AgentReceive(sdkParams *SdkParams, agentReceiveParams *AgentReceiveParams) *AgentReceiveResult {

	mhtField := "{\"payChannelType\":\"14\",\"appId\":\"" + agentReceiveParams.appId + "\",\"cardOwner\":\"" + agentReceiveParams.cardOwner + "\",\"mhtUserId\":\"" + agentReceiveParams.mhtUserId + "\",\"cardIdenNo\":\"" + agentReceiveParams.cardIdenNo + "\",\"cardPhoneNo\":\"" + agentReceiveParams.cardPhoneNo + "\",\"cardType\":\"" + agentReceiveParams.cardType + "\",\"cardIdenType\":\"" + agentReceiveParams.cardIdenType + "\",\"cardNo\":\"" + agentReceiveParams.cardNo + "\",\"accType\":\"" + agentReceiveParams.accType + "\",\"mhtReqTime\":\"" + agentReceiveParams.mhtReqTime + "\",\"agentPayMemo\":\"" + agentReceiveParams.agentPayMemo + "\",\"deviceType\":\"12\",\"notifyUrl\":\"" + agentReceiveParams.notifyUrl + "\",\"mhtUserCardId\":\"" + agentReceiveParams.mhtUserCardId + "\"}"
	payRequestParams := &paymentProto.PayRequestParams{}
	payRequestParams.TransId = agentReceiveParams.mhtOrderNo
	payRequestParams.ChannelType = "14"
	payRequestParams.Amount = agentReceiveParams.mhtOrderAmt
	payRequestParams.Router = "00010000000003"
	payRequestParams.MhtField = mhtField
	payRequestParams.Mode = 0
	datapayRequestParams, _ := proto.Marshal(payRequestParams)

	pInvoke := &paymentProto.PayInvoke{}
	pInvoke.Method = 0
	pInvoke.Params = datapayRequestParams
	datapInvoke, _ := proto.Marshal(pInvoke)

	content := &transactionProto.Content{}
	content.Nonce = sdkParams.merchantId + "_" + agentReceiveParams.mhtOrderNo + "_14"
	content.Data = datapInvoke
	datacontent, _ := proto.Marshal(content)

	transaction := &transactionProto.Transaction{}
	transaction.From = sdkParams.merchantId
	transaction.To = "3" //
	transaction.Content = datacontent
	datatransaction, _ := proto.Marshal(transaction)

	signedTransaction := &transactionProto.SignedTransaction{}
	signedTransaction.Transaction = datatransaction
	signedTransaction.Crypto = 0

	pk, _ := crypto.HexToECDSA(sdkParams.key)
	signature, _ := crypto.Sign(crypto.Keccak256(datatransaction), pk)
	signedTransaction.Signature = signature

	param := &reqParamProto.Param{}
	param.Method = 4
	param.Sign = signedTransaction
	dataparam, _ := proto.Marshal(param)

	result, _ := util.Post(sdkParams.reqUrl, dataparam)

	// 进行解码
	resultParam := &paramProto.Param{}
	err := proto.Unmarshal(result, resultParam)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	agentReceiveResult := &AgentReceiveResult{}
	var err1 = json.Unmarshal([]byte(resultParam.Message), &agentReceiveResult)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		return agentReceiveResult
	}
	return nil
}

/**
 * 单笔代付
 * @param agentPayParams,
 * agentPayParams.appId:应用编号
 * agentPayParams.accType: 入账账户类型。0 对私 1对公
 * agentPayParams.payeeName : 入账账户户名
 * agentPayParams.payeeCardNo :入账账户账号
 * agentPayParams.payeeCardUnionNo : 入账账户联行号
 * agentPayParams.agentPayMemo: 资金用途
 * notifyUrl : 商户通知地址
 * agentPayParams.mhtReqTime : 商户请求时间
 * agentPayParams.mhtOrderNo : 商户订单号
 * agentPayParams.mhtOrderAmt : 代付金额
 * @param sdkParams,
 * 	sdkParams.reqUrl 请求地址
 * 	sdkParams.key 请求流水号 商户生成的私钥(公钥发由现在支付)
 * 	sdkParams.merchantId 商户号,由现在支付分配
 * @return AgentPay
 *  AgentPayResult.appId,同输入
 *  AgentPayResult.mhtOrderNo,同输入
 *  AgentPayResult.nowPayOrderNo, 现在支付流水号
 *  AgentPayResult.responseCode 响应码(见文档附录)
 *  AgentPayResult.responseMsg 响应信息
 *  AgentPayResult.transStatus ,交易状态。SUCCESSED,成功 。FAILED,失败。PROCESSING,处理中。
 */
func AgentPay(sdkParams *SdkParams, agentPayParams *AgentPayParams) *AgentPayResult {
	mhtField := "{\"accType\":\"" + agentPayParams.accType + "\",\"payeeName\":\"" + agentPayParams.payeeName + "\",\"payeeCardUnionNo\":\"\",\"payChannelType\":\"15\",\"appId\":\"" + agentPayParams.appId + "\",\"mhtReqTime\":\"" + agentPayParams.mhtReqTime + "\",\"agentPayMemo\":\"" + agentPayParams.agentPayMemo + "\",\"deviceType\":\"13\",\"payeeCardNo\":\"" + agentPayParams.payeeCardNo + "\"}"

	payRequestParams := &paymentProto.PayRequestParams{}
	payRequestParams.TransId = agentPayParams.mhtOrderNo
	payRequestParams.ChannelType = "15"
	payRequestParams.Amount = agentPayParams.mhtOrderAmt
	payRequestParams.Router = "00010000000003"
	payRequestParams.MhtField = mhtField
	payRequestParams.Mode = 0
	datapayRequestParams, _ := proto.Marshal(payRequestParams)

	pInvoke := &paymentProto.PayInvoke{}
	pInvoke.Method = 0
	pInvoke.Params = datapayRequestParams
	datapInvoke, _ := proto.Marshal(pInvoke)

	content := &transactionProto.Content{}
	content.Nonce = sdkParams.merchantId + "_" + agentPayParams.mhtOrderNo + "_15"
	content.Data = datapInvoke
	datacontent, _ := proto.Marshal(content)

	transaction := &transactionProto.Transaction{}
	transaction.From = sdkParams.merchantId
	transaction.To = "3" //
	transaction.Content = datacontent
	datatransaction, _ := proto.Marshal(transaction)

	signedTransaction := &transactionProto.SignedTransaction{}
	signedTransaction.Transaction = datatransaction
	signedTransaction.Crypto = 0

	pk, _ := crypto.HexToECDSA(sdkParams.key)
	signature, _ := crypto.Sign(crypto.Keccak256(datatransaction), pk)
	signedTransaction.Signature = signature

	param := &reqParamProto.Param{}
	param.Method = 4
	param.Sign = signedTransaction
	dataparam, _ := proto.Marshal(param)

	result, _ := util.Post(sdkParams.reqUrl, dataparam)

	// 进行解码
	resultParam := &paramProto.Param{}
	err := proto.Unmarshal(result, resultParam)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	agentPayResult := &AgentPayResult{}
	var err1 = json.Unmarshal([]byte(resultParam.Message), &agentPayResult)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		return agentPayResult
	}

	return nil

}

/**
         * 交易查询
         * @param queryParams,
		 * queryParams.appId:商户应用唯一标识
		 * queryParams.mhtOrderNo: 商户请求流水号
 		 * @param sdkParams,
   		 * 	sdkParams.reqUrl 请求地址
 	 	 * 	sdkParams.key 请求流水号 商户生成的私钥(公钥发由现在支付)
 		 * 	sdkParams.merchantId 商户号,由现在支付分配
         * @return QueryResult,
         *  QueryResult.appId,同输入
         *  QueryResult.mhtOrderNo,同输入
         *  QueryResult.nowPayOrderNo, 现在支付流水号
         *  QueryResult.transType,交易类型： AGENT_PAY: 代付。    AGENT_RECEIVE：代收
         *  QueryResult.responseCode 响应码(见文档附录)
		 *  QueryResult.responseTime 应答时间
         *  QueryResult.responseMsg 响应信息
         *  QueryResult.mhtOrderAmt ,交易金额
         *  QueryResult.transStatus ,交易状态。SUCCESSED,成功 。FAILED,失败。PROCESSING,处理中。
         *
*/
func TransQuery(sdkParams *SdkParams, queryParams *QueryParams) *QueryResult {

	queryDetailParams := &paymentProto.QueryDetailParams{}

	queryDetailParams.AppId = queryParams.appId
	queryDetailParams.TransId = queryParams.transId
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
	content.Nonce = sdkParams.merchantId + "_" + "CS20171106174649"
	datapInvoke, errpInvoke := proto.Marshal(pInvoke)
	if errpInvoke != nil {
		log.Fatal("marshaling error: ", errpInvoke)
	}
	content.Data = datapInvoke
	//-----

	transaction := &transactionProto.Transaction{}
	transaction.From = sdkParams.merchantId
	transaction.To = "3" //
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

	pk, _ := crypto.HexToECDSA(sdkParams.key)
	signature, _ := crypto.Sign(crypto.Keccak256(dataptransaction), pk)
	signedTransaction.Signature = signature

	//----
	param := &reqParamProto.Param{}
	param.Method = 5
	param.Sign = signedTransaction

	datapparam, _ := proto.Marshal(param)

	result, _ := util.Post(sdkParams.reqUrl, datapparam)
	//	fmt.Println(fmt.Sprintf("%x", result))

	// 进行解码
	resultParam := &paramProto.Param{}
	err = proto.Unmarshal(result, resultParam)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	queryResult := &QueryResult{}
	var err1 = json.Unmarshal([]byte(resultParam.Message), &queryResult)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		return queryResult
	}

	return nil
}
