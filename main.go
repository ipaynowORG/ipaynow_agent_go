package main

import (
	"fmt"

	"github.com/ipaynowORG/ipaynow_agent_go/sdk"
)

func main() {
	sdkParams := sdk.MakeSdkParams("https://bc-test.ipaynow.cn/gateway",
		"013f81ac3ee1101b620031c00eac22ab53334c083c09fc191e05c29c9f0d26ad",
		"000100000000010000000000000001")

	//	queryResult := sdk.TransQuery(sdkParams, sdk.MakeQueryParams("1459846530407363", "CS20171106174649"))
	//	if queryResult != nil {
	//		fmt.Println(queryResult)
	//	}

	//	agentPayResult := sdk.AgentPay(sdkParams, sdk.MakeAgentPayParams(1, "1459846530407363", "test", "20170307100312", "CS20171123165910", "0", "张三", "6226200105111457", ""))
	//	if agentPayResult != nil {
	//		fmt.Println(agentPayResult)
	//	}

	//	agentReceiveResult := sdk.AgentReceive(sdkParams, sdk.MakeAgentReceiveParams("1459846530407363", 201, "CS20171123165910", "20171123165910", "112121", "21231", "6226888888888888", "张三", "0001", "0001", "14039999999999999", "13323233214", "test", "0", "http://posp.ipaynow.cn:10808/mobilewaptest/notify"))
	//	if agentReceiveResult != nil {
	//		fmt.Println(agentReceiveResult)
	//	}

	//	agentPayRefundQueryResult := sdk.AgentPayRefundQuery(sdkParams, sdk.MakeAgentPayRefundQueryParams("1459846530407363", "gzh20170414145229wiY2pCpjlk8pwuKK"))
	//	if agentPayRefundQueryResult != nil {
	//		fmt.Println(agentPayRefundQueryResult)
	//	}

	//	batchQueryResult := sdk.AgentPayRefundBatchQuery(sdkParams, sdk.MakeBatchQueryParams("1459846530407363", "CS20171123165910", "20170414", 1, 100))
	//	if batchQueryResult != nil {
	//		fmt.Println(batchQueryResult)
	//	}

	//	balanceQueryResp := sdk.BalanceQuery(sdkParams, sdk.MakeBalanceQueryParams("1459846530407363", "test11112"))
	//	if balanceQueryResp != nil {
	//		fmt.Println(balanceQueryResp)
	//	}
}
