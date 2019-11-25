package sms

import (
	"encoding/json"
	"fmt"
	"gin-project-template/pkg/util"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type info struct {
	PhoneNumberSet   []string `json:"PhoneNumberSet"`
	TemplateID       string   `json:"TemplateID"`
	SmsSdkAppid      string   `json:"SmsSdkAppid"`
	Sign             string   `json:"Sign"`
	TemplateParamSet []string `json:"TemplateParamSet"`
}

func Send(phone []string, templateID string, params []string) {
	appConf, err := util.FetchAppConf()
	if err != nil {
		panic(err)
	}
	tencentConf := appConf.Tencent
	credential := common.NewCredential(
		tencentConf.SecretId,
		tencentConf.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "", cpf)

	request := sms.NewSendSmsRequest()

	smsConf := appConf.Sms
	i := info{
		PhoneNumberSet:   phone,
		TemplateID:       templateID,
		TemplateParamSet: params,
		Sign:             smsConf.Sign,
		SmsSdkAppid:      smsConf.SdkAppid,
	}
	p, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	err = request.FromJsonString(string(p))
	if err != nil {
		panic(err)
	}
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
