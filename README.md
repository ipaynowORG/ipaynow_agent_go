
# 代收付SDK使用说明 #

## 版本变更记录 ##


- 1.0.0 : 初稿

## 目录 ##

[1. 概述](#1)

&nbsp;&nbsp;&nbsp;&nbsp;[1.1 简介](#1.1)

&nbsp;&nbsp;&nbsp;&nbsp;[1.2 如何获取](#1.2)

&nbsp;&nbsp;&nbsp;&nbsp;[1.3 接口地址](#1.3)

[2. API](#2)

&nbsp;&nbsp;&nbsp;&nbsp;[2.1 代收付API](#2.1)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[单笔代收](#2.1.1)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[单笔代付](#2.1.2)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[账户余额查询](#2.1.3)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[交易查询](#2.1.4)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[单笔代付退票查询](#2.1.5)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[代付退票批量查询](#2.1.6)

&nbsp;&nbsp;&nbsp;&nbsp;[2.2 支付结果通知](#2.2)

[3. 生成公私钥方法](#3)


<h2 id='1'> 1. 概述 </h2>

<h4 id='1.1'> 1.1 简介 </h4>

- 代收: 代收是指现在支付商户与自己的客户签订代收协议，并委托现在支付对已经与自己签订代收协议的商户发起银行卡扣款，所扣款项进入现在支付商户的余额账户中。商户可以选择代扣资金是停留在商户余额账户中，还是T+1结算入商户的结算账户中。

- 代付: 代付是指现在支付商户根据自己的业务需要先将业务资金充值入自己在现在支付的余额账户中，然后发起指令，委托现在支付从自己的虚拟账户中将业务资金转账到指令中指定的银行卡中。业务资金将实时入账到指令指定的收款人银行卡。

<h4 id='1.2'> 1.2 如何获取 </h4>

[获取源码](https://github.com/ipaynowORG/ipaynow_agent_java)

<h4 id='1.3'> 1.3 接口地址 </h4>

    测试接口url: http://bc-test.ipaynow.cn/gateway

    正式接口url: https://bc-pay.ipaynow.cn/gateway



<h2 id='2'> 2. API </h2>

使用"go get github.com/ipaynowORG/ipaynow_agent_go/sdk" 命令clone 并 install ipaynow_agent模块

代码中 import (git "github.com/ipaynowORG/ipaynow_agent_go/sdk")使用

<h4 id='2.1'> 2.1 代收付API </h4>

<h5 id='2.1.1'></h4>

- 单笔代收

		/**
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
        func AgentReceive(sdkParams *SdkParams, agentReceiveParams *AgentReceiveParams) *AgentReceiveResult 

<h5 id='2.1.2'></h4>

- 单笔代付

		/**
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
        func AgentPay(sdkParams *SdkParams, agentPayParams *AgentPayParams) *AgentPayResult

<h5 id='2.1.3'></h4>

- 账户余额查询

		/**
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
        func BalanceQuery(sdkParams *SdkParams, balanceQueryParams *BalanceQueryParams) *BalanceQueryResp 

<h5 id='2.1.4'></h4>

- 交易查询

		/**
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
        func TransQuery(sdkParams *SdkParams, queryParams *QueryParams) *QueryResult

<h5 id='2.1.5'></h4>

- 单笔代付退票查询

		/**
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
        func AgentPayRefundQuery(sdkParams *SdkParams, agentPayRefundQueryParams *AgentPayRefundQueryParams) *AgentPayRefundQueryResult 

<h5 id='2.1.6'></h4>

- 代付退票批量查询

		/**
		 * 代付退票批量查询
		 * @param batchQueryParams
		 * batchQueryParams.appId 商户应用唯一标识
		 * batchQueryParams.mhtOrderNo 商户请求流水号
		 * batchQueryParams.nowPage 商户请求页码
		 * batchQueryParams.pageSize 商户请求数据量
		 * batchQueryParams.refundDate 退票日期
		 
		 * @param sdkParams,
		 * 	sdkParams.reqUrl 请求地址
		 * 	sdkParams.key 请求流水号 商户生成的私钥(公钥发由现在支付)
		 * 	sdkParams.merchantId 商户号,由现在支付分配
		 
		 * @return BatchQueryResult
		 * BatchQueryResult.responseCode 响应码(见文档附录)
		 * BatchQueryResult.responseMsg 响应信息
		 * BatchQueryResult.mhtOrderNo,同输入
		 * BatchQueryResult.appId,同输入
		 * BatchQueryResult.nowPage,同输入
		 * BatchQueryResult.pageSize,同输入
		 * BatchQueryResult.countNums,数据总数
		 * BatchQueryResult.countPages,总页数
		 * agentPayRefundQueryRespFieldList 退票详情
		 */
        func AgentPayRefundBatchQuery(sdkParams *SdkParams, batchQueryParams *BatchQueryParams) *BatchQueryResult 

<h4 id='2.2'>2.2 支付结果通知</h4>

- 通知方式采用httppost方式通知,接受demo如下

        func handler(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        if r.Method == "POST" {
            result, _:= ioutil.ReadAll(r.Body)
			r.Body.Close()
			//报文数据字符串
			fmt.Printf("%s\n", result)
        }
	...
	}

- 结果通知现在支付请求数据示例：

        {“data”:{“transType”:”PAY”,”chTransId”:”200003201703221910181130275”,”mhtOrderAmt”:”1”,”responseCode”:”A001”,”responseMsg”:”成功”,”responseTime”:”20170322191042”,
        ”transStatus”:”SUCCESSED”，“mhtOrderNo”:”20170322191017s”},
        ”sign_r”:”48f9a9554a3914e5d39489a3e626bd2d9b687455a1683631f5a64c5a0fed9af9”,sign_s:”152a310c0dac28c624783a9090ddd8717718c12629dbf0224c1a1170b03beaf4”,”sign_v”:”28”}

- 结果通知商户响应数据示例：

        {“success”:”Y”}

    若商户未成功响应现在支付发出的通知，现在支付会重新通知商户，规则如下：
    1.前5分钟，每分钟通知一次，累计5次
    2.5分钟后，每5分钟通知一次，累计5次

字段含义如下:

<table>
        <tr>
            <th>字段名称</th>
            <th>字段Key</th>
            <th>格式</th>
            <th>必填</th>
            <th>备注</th>
        </tr>
        <tr>
            <td>商户请求流水号</td>
            <td>mhtOrderNo</td>
            <td>String (1,40)</td>
            <td>Y</td>
            <td>同输入</td>
        </tr>
        <tr>
            <td>现在支付流水号</td>
            <td>chTransId</td>
            <td>String(1,60)</td>
            <td>N</td>
            <td></td>
        </tr>
        <tr>
            <td>交易类型</td>
            <td>transType</td>
            <td>String(1,60)</td>
            <td>N</td>
            <td>AGENT_PAY: 代付 。AGENT_RECEIVE：代收</td>
        </tr>
        <tr>
            <td>响应码</td>
            <td>responseCode</td>
            <td>String(4)</td>
            <td>Y</td>
            <td>见文档附录</td>
        </tr>
        <tr>
            <td>响应信息</td>
            <td>responseMsg</td>
            <td>String(1,100)</td>
            <td>Y</td>
            <td>相应信息</td>
        </tr>
        <tr>
            <td>响应时间</td>
            <td>responseTime</td>
            <td>String(14)</td>
            <td>Y</td>
            <td>yyyyMMddHHmmss</td>
        </tr>
        <tr>
            <td>交易金额</td>
            <td>mhtOrderAmt</td>
            <td>Number(22)</td>
            <td>N</td>
            <td>单位:分</td>
        </tr>
        <tr>
            <td>交易状态</td>
            <td>transStatus</td>
            <td>String(4)</td>
            <td>N</td>
            <td>SUCCESSED,成功。FAILED,失败。PROCESSING,处理中。</td>
        </tr>
    </table>

- 验证签名

        public static boolean verify(String data, String pub) {
                try {
                    ObjectMapper mapper = new ObjectMapper();
                    Map<String, Object> map = mapper.readValue(data, Map.class);
                    BigInteger r = new BigInteger(map.get("sign_r").toString(), 16);
                    BigInteger s = new BigInteger(map.get("sign_s").toString(), 16);
                    int v = Integer.parseInt(map.get("sign_v").toString());
                    String param = StringTools.postFormLinkReport((Map) map.get("data"));
        ECKey.ECDSASignature sig =
        ECKey.ECDSASignature.fromComponents(r.toByteArray(), s.toByteArray(),
                            (byte) v);
                    ECKey key = ECKey.fromPublicOnly(Hex.decode(pub));
                    return key.verify(HashUtil.sha3(param.getBytes("UTF-8")), sig);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                return false;
        }

<h2 id='3'> 3. 生成公私钥方法 </h2>

为了保证信息的安全性，公私钥信息为商户侧生成；生成完毕后，需将公钥信息给予现在支付侧；同时需记录现在支付给商户侧的公钥信息，以便双方进行数据解密。

    生成方法：
	
	"github.com/ethereum/go-ethereum/crypto"
	
    crypto.GenerateKey()