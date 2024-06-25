// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销加价购对外API
//
// 指定服务商可通过该接口报名加价购活动、查询某个区域内的加价购活动列表、锁定加价活动购资格以及解锁加价购活动资格。
//
// API version: 1.4.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package retailstore_test

import (
	"context"
	"log"

	"github.com/LoveChiba/wechatpay-go/core"
	"github.com/LoveChiba/wechatpay-go/core/option"
	"github.com/LoveChiba/wechatpay-go/services/retailstore"
	"github.com/LoveChiba/wechatpay-go/utils"
)

func ExampleActivityApiService_ApplyActivity() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := retailstore.ActivityApiService{Client: client}
	resp, result, err := svc.ApplyActivity(ctx,
		retailstore.ApplyActivityRequest{
			ActivityId:       core.String("123"),
			CallerMerchantId: core.String("1894101023"),
			ApplyInfos: []retailstore.ActApplyInfo{retailstore.ActApplyInfo{
				StoreInfo: &retailstore.StoreInfo{
					StoreId:              core.String("100"),
					AccountingMerchantId: core.String("2831255701"),
					MerchantId:           core.String("6281399112"),
				},
				GoodsOriginalPrice: core.Int64(100),
			}},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ApplyActivity err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleActivityApiService_ListActsByArea() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := retailstore.ActivityApiService{Client: client}
	resp, result, err := svc.ListActsByArea(ctx,
		retailstore.ListActsByAreaRequest{
			CityId: core.String("123"),
			Offset: core.Int64(0),
			Limit:  core.Int64(20),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListActsByArea err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
