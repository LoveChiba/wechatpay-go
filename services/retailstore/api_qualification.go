// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销加价购对外API
//
// 指定服务商可通过该接口报名加价购活动、查询某个区域内的加价购活动列表、锁定加价活动购资格以及解锁加价购活动资格。
//
// API version: 1.4.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package retailstore

import (
	"context"
	nethttp "net/http"
	neturl "net/url"

	"github.com/LoveChiba/wechatpay-go/core"
	"github.com/LoveChiba/wechatpay-go/core/consts"
	"github.com/LoveChiba/wechatpay-go/services"
)

type QualificationApiService services.Service

// LockQualification 锁定品牌加价购活动资格
//
// 接口介绍：指定服务商仅能通过该接口锁定品牌加价购活动资格。品牌加价购活动资格被成功锁定后不能被其它订单使用，且成功锁定品牌加价购活动资格是微信支付系统对门店和服务商发放补贴资金的前置条件，否则微信支付系统不对门店和服务商发放补贴资金。品牌加价购活动资格被锁定后，如未被解锁，则不能被其它订单使用。
// 使用对象：指定服务商。
func (a *QualificationApiService) LockQualification(ctx context.Context, req LockQualificationRequest) (resp *LockQualificationResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/goods-subsidy-activity/qualification/lock"
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

	// Extract LockQualificationResponse from Http Response
	resp = new(LockQualificationResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// UnlockQualification 解锁品牌加价购活动资格
//
// 接口介绍：指定服务商仅能通过该接口解锁品牌加价购活动资格。商户调用微信支付下单接口失败或者在其它交易流程被阻断的场景下，如果此时品牌加价购活动资格已被当前订单锁定，需调用该接口解锁被锁定的品牌加价购活动资格，否则，如果被锁定的品牌加价购活动资格未被成功解锁，该品牌加价购活动资格无法被使用。
// 使用对象：指定服务商。
func (a *QualificationApiService) UnlockQualification(ctx context.Context, req UnlockQualificationRequest) (resp *UnlockQualificationResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/goods-subsidy-activity/qualification/unlock"
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

	// Extract UnlockQualificationResponse from Http Response
	resp = new(UnlockQualificationResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
