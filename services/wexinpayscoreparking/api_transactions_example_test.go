// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付分停车服务
//
// 微信支付分停车服务 扣费API
//
// API version: 1.2.1

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package wexinpayscoreparking_test

import (
	"context"
	"log"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/wexinpayscoreparking"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func ExampleTransactionsApiService_CreateTransaction() {
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

	svc := wexinpayscoreparking.TransactionsApiService{Client: client}
	resp, result, err := svc.CreateTransaction(ctx,
		wexinpayscoreparking.CreateTransactionRequest{
			Appid:         core.String("wxcbda96de0b165486"),
			SubAppid:      core.String("wxcbda96de0b165486"),
			SubMchid:      core.String("1900000109"),
			Description:   core.String("停车场扣费"),
			Attach:        core.String("深圳分店"),
			OutTradeNo:    core.String("20150806125346"),
			TradeScene:    core.String("PARKING"),
			GoodsTag:      core.String("WXG"),
			NotifyUrl:     core.String("https://yoursite.com/wxpay.html"),
			ProfitSharing: core.String("Y"),
			Amount: &wexinpayscoreparking.OrderAmount{
				Total:    core.Int64(888),
				Currency: core.String("CNY"),
			},
			ParkingInfo: &wexinpayscoreparking.ParkingTradeScene{
				ParkingId:        core.String("5K8264ILTKCH16CQ250"),
				PlateNumber:      core.String("粤B888888"),
				PlateColor:       wexinpayscoreparking.PLATECOLOR_BLUE.Ptr(),
				StartTime:        core.Time(time.Now()),
				EndTime:          core.Time(time.Now()),
				ParkingName:      core.String("欢乐海岸停车场"),
				ChargingDuration: core.Int64(3600),
				DeviceId:         core.String("12313"),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateTransaction err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransactionsApiService_QueryTransaction() {
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

	svc := wexinpayscoreparking.TransactionsApiService{Client: client}
	resp, result, err := svc.QueryTransaction(ctx,
		wexinpayscoreparking.QueryTransactionRequest{
			SubMchid:   core.String("1900000109"),
			OutTradeNo: core.String("20150806125346"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryTransaction err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
