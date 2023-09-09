package utils

import (
	"testing"

	"github.com/ozeer/go-api/model/common/request"
)

type PageInfoTest struct {
	PageInfo request.PageInfo
	Name     string
}

type RegisterInfoTest struct {
	UserName string
	Phone    string
	Password string
	Email    string
	Sex      uint8
	Birthday string
}

func TestVerify(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 0
	testInfo.PageInfo.PageSize = 0
	err := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉0值")
	}
	testInfo.Name = ""
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能正常检测name为空")
	}
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通过检测")
	}
}

func TestRegisterVerify(t *testing.T) {
	var testInfo RegisterInfoTest
	testInfo.UserName = "Tom"
	testInfo.Password = "123456zy"
	testInfo.Phone = "18701588471"
	testInfo.Email = "demo2023@163.com"
	testInfo.Birthday = "2022-02-29"
	err := Verify(testInfo, RegisterVerify)
	if err != nil {
		t.Error(err)
	}
}
