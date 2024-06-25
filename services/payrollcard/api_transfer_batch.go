// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微工卡接口文档
//
// 服务商通过本API文档提供的接口，查询商户和微工卡的授权关系、生成预授权的token口令、核身预下单、核身结果的查询等。
//
// API version: 1.5.2

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package payrollcard

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"

	"github.com/LoveChiba/wechatpay-go/core"
	"github.com/LoveChiba/wechatpay-go/core/consts"
	"github.com/LoveChiba/wechatpay-go/services"
)

type TransferBatchApiService services.Service

// CreateTransferBatch 发起批量转账
//
// ## 发起批量转账
//
// 服务商可以通过该接口，批量向用户零钱进行转账操作。请求消息中应包含特约商户号、特约商户授权的AppID、授权类型、商家批次单号、转账名称、转账总金额、转账总笔数、转账OpenID、收款用户姓名、服务商AppID、转账用途等信息
// 1、当特约商户授权类型为INFORMATION\_AUTHORIZATION\_TYPE（特约商户信息授权），需要填特约商户的公众号AppID，特约商户公众号AppID的用户的OpenID
// 2、当特约商户授权类型为FUND\_AUTHORIZATION\_TYPE（特约商户资金授权），需要填服务商的公众号AppID，服务商公众号AppID的用户的OpenID
//
// 接口限频：
// 单个服务商（发起批量转账接口）50QPS，如果超过频率限制，会报错FREQUENCY\_LIMITED，请降低频率请求。
//
// 注意事项：
//
// * 因服务商自身系统设置存在问题导致的资金损失，由服务商自行承担。
// * 批量转账一旦发起后，不允许撤销。
// * 转账批次单和明细单中涉及金额的字段单位为“分”。
// * 微信支付视任何不同“发起的服务商商户号+商家批次单号（out\_batch\_no）”的请求为一个全新的批次。在未查询到明确的转账批次单处理结果之前，请勿修改商家批次单号重新提交！如有发生，服务商应当自行承担因此产生的所有损失和责任。
// * 当返回错误时，请不要更换商家批次单号，一定要使用原商家批次单号重试，否则可能造成重复转账等资金风险。
// * 如果遇到回包返回新的错误码，请务必不要换单重试，请联系客服确认转账情况。如果有新的错误码，会更新到此API文档中。
// * 错误码描述字段message只供人工定位问题时做参考，系统实现时请不要依赖这个字段来做自动化处理。
// * 请服务商在自身的系统中合理设置转账频次并做好并发控制，防范错付风险。
func (a *TransferBatchApiService) CreateTransferBatch(ctx context.Context, req CreateTransferBatchRequest) (resp *TransferBatchEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// 对请求中敏感字段进行加密
	encReq := req.Clone()
	encryptCertificate, err := a.Client.EncryptRequest(ctx, encReq)
	if err != nil {
		return nil, nil, fmt.Errorf("encrypt request failed: %v", err)
	}

	if encryptCertificate != "" {
		localVarHeaderParams.Set(consts.WechatPaySerial, encryptCertificate)
	}
	req = *encReq

	localVarPath := consts.WechatPayAPIServer + "/v3/payroll-card/transfer-batches"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract TransferBatchEntity from Http Response
	resp = new(TransferBatchEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
