package utils

import "regexp"

// 常用正则校验格式
var (
	// 微信号（6至20位，以字母开头，可包含字母、数字、减号和下划线，如：wechat01）
	REGEXP_WE_CHAT = `^[a-zA-Z][-_a-zA-Z0-9]{5,19}$`
	// QQ号（如：123456）
	REGEXP_QQ_NUMBER = `^[1-9][0-9]{4,10}$`
	// 金额（正数，可有最多两位小数，如：8.99）
	REGEXP_MONEY = `(?:^[1-9]([0-9]+)?(?:\.[0-9]{1,2})?$)|(?:^(?:0)$)|(?:^[0-9]\.[0-9](?:[0-9])?$)`
	// 用户名校验（4到16位，可包含字母、数字和下划线，如：test01）
	REGEXP_USERNAME = `^[a-zA-Z0-9_]{4,16}$`
	// 邮箱格式
	REGEXP_EMAIL = `^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
	// 密码8位字母与数字组合
	REGEXP_PASSWORD = `^[a-zA-Z0-9]{8}$`
	// 检查日期格式（可判断闰年，格式：YYYY-MM-DD，如：2021-12-01）
	REGEXP_DATE = `^(([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))-02-29)$`
	// 时间（12小时，格式：hh:mm:ss）
	REGEXP_TIME_12 = `^(?:1[0-2]|0?[1-9]):[0-5]\d:[0-5]\d$`
	// 时间（24小时，格式：HH:mm:ss）
	REGEXP_TIME_24 = `^(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$`
	// 中国大陆手机号格式校验（严格，2019年最新手机号段，如：19876809091）
	REGEXP_PHONE = `^(?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[189]))\d{8}$`
	// 中国大陆身份证号（1代，15位，如：123419901220321）
	REGEXP_ID_CARD_NUMBER_FIRST = `^[1-9]\d{7}(?:0\d|10|11|12)(?:0[1-9]|[1-2][\d]|30|31)\d{3}$`
	// 中国大陆身份证号（2代，18位，如：429004199801012921）
	REGEXP_ID_CARD_NUMBER_SECOND = `^[1-9]\d{5}(?:18|19|20)\d{2}(?:0[1-9]|10|11|12)(?:0[1-9]|[1-2]\d|30|31)\d{3}[\dXx]$`
	// 香港身份证（如：A123456(1)）
	REGEXP_HONGKONG_ID_CARD = `^[a-zA-Z]\d{6}\([\dA]\)$`
	// 澳门身份证（如：1234567(1)）
	REGEXP_MACAU_ID_CARD = `^[1|5|7]\d{6}\(\d\)$`
	// 台湾身份证（如：A123456789）
	REGEXP_TAIWAN_ID_CARD = `^[a-zA-Z][0-9]{9}$`
	// 网址URL（带端口号，如：https://www.baidu.com:8080/）
	REGEXP_URL_WITH_PORT = `^((ht|f)tps?:\/\/)?[\w-]+(\.[\w-]+)+:\d{1,5}\/?$`
	// 网址URL（不带端口号，如：https://www.baidu.com/）
	REGEXP_URL = `^(((ht|f)tps?):\/\/)?([^!@#$%^&*?.\s-]([^!@#$%^&*?.\s]{0,63}[^!@#$%^&*?.\s])?\.)+[a-z]{2,6}\/?`
)

var (
	IdVerify               = Rules{"ID": []string{NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"UserName": {NotEmpty()}, "Phone": {NotEmpty(), RegexpMatch(REGEXP_PHONE)}, "Password": {NotEmpty(), RegexpMatch(REGEXP_PASSWORD)}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AutoPackageVerify      = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
)

func RegexpVerify(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}
