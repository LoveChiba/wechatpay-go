// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付营销系统开放API
//
// 新增立减金api
//
// API version: 3.4.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package cashcoupons

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/LoveChiba/wechatpay-go/core"
	"github.com/LoveChiba/wechatpay-go/core/consts"
	"github.com/LoveChiba/wechatpay-go/services"
)

type CouponApiService services.Service

// ListCouponsByFilter 根据过滤条件查询用户的券
//
// 根据过滤条件查询用户的券（商户号角色、批次、状态）
//
// 接口频率：500qps
//
// 前置条件：已发代金券
//
// 注意事项：
//
// 1. 创建方只能查自制批次的券，发券商户只能查询自发批次的券
func (a *CouponApiService) ListCouponsByFilter(ctx context.Context, req ListCouponsByFilterRequest) (resp *CouponCollection, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.Openid == nil {
		return nil, nil, fmt.Errorf("field `Openid` is required and must be specified in ListCouponsByFilterRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/users/{openid}/coupons"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"openid"+"}", neturl.PathEscape(core.ParameterToString(*req.Openid, "")), -1)

	// Make sure All Required Params are properly set
	if req.Appid == nil {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in ListCouponsByFilterRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("appid", core.ParameterToString(*req.Appid, ""))
	if req.StockId != nil {
		localVarQueryParams.Add("stock_id", core.ParameterToString(*req.StockId, ""))
	}
	if req.Status != nil {
		localVarQueryParams.Add("status", core.ParameterToString(*req.Status, ""))
	}
	if req.CreatorMchid != nil {
		localVarQueryParams.Add("creator_mchid", core.ParameterToString(*req.CreatorMchid, ""))
	}
	if req.SenderMchid != nil {
		localVarQueryParams.Add("sender_mchid", core.ParameterToString(*req.SenderMchid, ""))
	}
	if req.AvailableMchid != nil {
		localVarQueryParams.Add("available_mchid", core.ParameterToString(*req.AvailableMchid, ""))
	}
	if req.Offset != nil {
		localVarQueryParams.Add("offset", core.ParameterToString(*req.Offset, ""))
	}
	if req.Limit != nil {
		localVarQueryParams.Add("limit", core.ParameterToString(*req.Limit, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract CouponCollection from Http Response
	resp = new(CouponCollection)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryCoupon 查询代金券详情
//
// 查询用户的券详情
//
// 接口频率：500qps
//
// 前置条件：用户的券已发放
//
// 注意：
//
// 1. 该接口支持批次创建商户号与批次发放商户调用
func (a *CouponApiService) QueryCoupon(ctx context.Context, req QueryCouponRequest) (resp *Coupon, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.CouponId == nil {
		return nil, nil, fmt.Errorf("field `CouponId` is required and must be specified in QueryCouponRequest")
	}
	if req.Openid == nil {
		return nil, nil, fmt.Errorf("field `Openid` is required and must be specified in QueryCouponRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/users/{openid}/coupons/{coupon_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"coupon_id"+"}", neturl.PathEscape(core.ParameterToString(*req.CouponId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"openid"+"}", neturl.PathEscape(core.ParameterToString(*req.Openid, "")), -1)

	// Make sure All Required Params are properly set
	if req.Appid == nil {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in QueryCouponRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("appid", core.ParameterToString(*req.Appid, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Coupon from Http Response
	resp = new(Coupon)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// SendCoupon 发放指定批次的代金券
//
// 发放指定批次的代金券
//
// 接口频率：500qps
//
// 前置条件：已创建并激活代金券批次
//
// 是否支持幂等：是
//
// 注意：
//
// 1. 商户可在H5活动页面、商户小程序、商户APP等自有场景内调用该接口完成发券，商户默认只允许发放本商户号（调用发券接口的商户号）创建的代金券，如需发放其他商户商户创建的代金券，请参考常见问题Q1。
// 2. 跨商户发券时，请求参数中除了stock\_id和stock\_creator\_mchid为创建方提供的数据，其他的所有调用数据都由发放方提供。
func (a *CouponApiService) SendCoupon(ctx context.Context, req SendCouponRequest) (resp *SendCouponResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.Openid == nil {
		return nil, nil, fmt.Errorf("field `Openid` is required and must be specified in SendCouponRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/users/{openid}/coupons"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"openid"+"}", neturl.PathEscape(core.ParameterToString(*req.Openid, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &SendCouponBody{
		StockId:           req.StockId,
		OutRequestNo:      req.OutRequestNo,
		Appid:             req.Appid,
		StockCreatorMchid: req.StockCreatorMchid,
		CouponValue:       req.CouponValue,
		CouponMinimum:     req.CouponMinimum,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract SendCouponResponse from Http Response
	resp = new(SendCouponResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
