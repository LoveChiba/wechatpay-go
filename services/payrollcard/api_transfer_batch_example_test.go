// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微工卡接口文档
//
// 服务商通过本API文档提供的接口，查询商户和微工卡的授权关系、生成预授权的token口令、核身预下单、核身结果的查询等。
//
// API version: 1.5.2

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package payrollcard_test

import (
	"context"
	"log"

	"github.com/LoveChiba/wechatpay-go/core"
	"github.com/LoveChiba/wechatpay-go/core/option"
	"github.com/LoveChiba/wechatpay-go/services/payrollcard"
	"github.com/LoveChiba/wechatpay-go/utils"
)

func ExampleTransferBatchApiService_CreateTransferBatch() {
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

	svc := payrollcard.TransferBatchApiService{Client: client}
	resp, result, err := svc.CreateTransferBatch(ctx,
		payrollcard.CreateTransferBatchRequest{
			SubMchid:          core.String("1900000109"),
			SubAppid:          core.String("wxf636efh567hg4356"),
			AuthorizationType: payrollcard.AUTHTYPE_INFORMATION_AUTHORIZATION_TYPE.Ptr(),
			OutBatchNo:        core.String("plfk2020042013"),
			BatchName:         core.String("2019年1月深圳分部报销单"),
			BatchRemark:       core.String("2019年1月深圳分部报销单"),
			TotalAmount:       core.Int64(4000000),
			TotalNum:          core.Int64(200),
			TransferDetailList: []payrollcard.TransferDetailInput{payrollcard.TransferDetailInput{
				OutDetailNo:    core.String("x23zy545Bd5436"),
				TransferAmount: core.Int64(200000),
				TransferRemark: core.String("2020年4月报销"),
				Openid:         core.String("o-MYE42l80oelYMDE34nYD456Xoy"),
				UserName:       core.String("757b340b45ebef5467rter35gf464344v3542sdf4t6re4tb4f54ty45t4yyry45"),
			}},
			SpAppid:         core.String("wxf636efh567hg4388"),
			EmploymentType:  payrollcard.EMPLOYMENTTYPE_LONG_TERM_EMPLOYMENT.Ptr(),
			EmploymentScene: payrollcard.EMPLOYMENTSCENE_LOGISTICS.Ptr(),
			BusinessType:    payrollcard.BUSINESSTYPE_UNDEFINE.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateTransferBatch err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
