// Copyright 2021 Tencent Inc. All rights reserved.
//
// H5支付
//
// H5支付API
//
// API version: 1.2.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package h5_test

import (
	"context"
	"log"
	"time"

	"github.com/LoveChiba/wechatpay-go/core"
	"github.com/LoveChiba/wechatpay-go/core/option"
	"github.com/LoveChiba/wechatpay-go/services/partnerpayments/h5"
	"github.com/LoveChiba/wechatpay-go/utils"
)

func ExampleH5ApiService_CloseOrder() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := h5.H5ApiService{Client: client}
	result, err := svc.CloseOrder(ctx,
		h5.CloseOrderRequest{
			OutTradeNo: core.String("OutTradeNo_example"),
			SpMchid:    core.String("1230000109"),
			SubMchid:   core.String("1230000109"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CloseOrder err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d", result.Response.StatusCode)
	}
}

func ExampleH5ApiService_Prepay() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := h5.H5ApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		h5.PrepayRequest{
			SpAppid:       core.String("wxd678efh567hg6787"),
			SpMchid:       core.String("1230000109"),
			SubAppid:      core.String("wxd678efh567hg6787"),
			SubMchid:      core.String("1230000109"),
			Description:   core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:    core.String("1217752501201407033233368018"),
			TimeExpire:    core.Time(time.Now()),
			Attach:        core.String("自定义数据说明"),
			NotifyUrl:     core.String("https://www.weixin.qq.com/wxpay/pay.php"),
			GoodsTag:      core.String("WXG"),
			LimitPay:      []string{"LimitPay_example"},
			SupportFapiao: core.Bool(false),
			Amount: &h5.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(100),
			},
			Detail: &h5.Detail{
				CostPrice: core.Int64(608800),
				GoodsDetail: []h5.GoodsDetail{h5.GoodsDetail{
					GoodsName:        core.String("iPhoneX 256G"),
					MerchantGoodsId:  core.String("ABC"),
					Quantity:         core.Int64(1),
					UnitPrice:        core.Int64(828800),
					WechatpayGoodsId: core.String("1001"),
				}},
				InvoiceId: core.String("wx123"),
			},
			SceneInfo: &h5.SceneInfo{
				DeviceId: core.String("013467007045764"),
				H5Info: &h5.H5Info{
					AppName:     core.String("王者荣耀"),
					AppUrl:      core.String("https://pay.qq.com"),
					BundleId:    core.String("com.tencent.wzryiOS"),
					PackageName: core.String("com.tencent.tmgp.sgame"),
					Type:        core.String("iOS"),
				},
				PayerClientIp: core.String("14.23.150.211"),
				StoreInfo: &h5.StoreInfo{
					Address:  core.String("广东省深圳市南山区科技中一道10000号"),
					AreaCode: core.String("440305"),
					Id:       core.String("0001"),
					Name:     core.String("腾讯大厦分店"),
				},
			},
			SettleInfo: &h5.SettleInfo{
				ProfitSharing: core.Bool(false),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleH5ApiService_QueryOrderById() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := h5.H5ApiService{Client: client}
	resp, result, err := svc.QueryOrderById(ctx,
		h5.QueryOrderByIdRequest{
			TransactionId: core.String("TransactionId_example"),
			SpMchid:       core.String("SpMchid_example"),
			SubMchid:      core.String("SubMchid_example"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryOrderById err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleH5ApiService_QueryOrderByOutTradeNo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := h5.H5ApiService{Client: client}
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		h5.QueryOrderByOutTradeNoRequest{
			OutTradeNo: core.String("OutTradeNo_example"),
			SpMchid:    core.String("SpMchid_example"),
			SubMchid:   core.String("SubMchid_example"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryOrderByOutTradeNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
