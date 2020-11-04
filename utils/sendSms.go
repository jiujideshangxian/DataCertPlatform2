package utils

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type SmsCode struct{
	Code string `json:"code"`
}
type  SmsResutl struct{

}

const SMS_TLP_REGISTER = ""//注册业务的短信模板
const SMS_TLP_  = ""//用户登入短信模板
const SMS_TLP_KYC  = ""//实名认证的短信模板

/*
	dy：dynamic动态的
	函数用来发送一条短信
	phone:电话，接受验证码的号码
	code:发送验证码数字
	templateType
 */
func SendSms(phone string, code string,templateType string)(*SmsResutl,error){
	config:=beego.AppConfig
	//获取配置文件的accessKey
	accessKey:=config.String("sms_access_key")
	//secret
	client,err:=dysmsapi.NewClientWithAccessKey()
	if err!=nil{
		return nil,err
	}
	//batch:批量，批次
	request:=dysmsapi.CreateSendSmRequest()
	request.PhoneNumbers = phone
	request.SiginName = "线上餐厅"
	request.TemplateCode =templateType

	SmsCode := SmsCode{
		Code :code,
	}

	smsbytes,_=json.Marshal(SmsCode)
	request.TemplateParam = string(smsbytes)

	response,err:=client.SendSms(request)
		if err!=nil{
			return nil,err
		}
		//Biz : business,商务，业务。
	smsResutl:=SmsResutl{

	}
}