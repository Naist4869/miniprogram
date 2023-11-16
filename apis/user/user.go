// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package user 用户
package user

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiCode2Session       = "/sns/jscode2session"
	apiGetPaidUnionId     = "/wxa/getpaidunionid"
	apiGetUserPhoneNumber = "/wxa/business/getuserphonenumber"
)

type PhoneNumber struct {
	PhoneInfo struct {
		PhoneNumber     string `json:"phoneNumber"`
		PurePhoneNumber string `json:"purePhoneNumber"`
		CountryCode     string `json:"countryCode"`
		Watermark       struct {
			Timestamp int    `json:"timestamp"`
			Appid     string `json:"appid"`
		} `json:"watermark"`
	} `json:"phone_info"`
}

func GetUserPhoneNumber(ctx *miniprogram.Miniprogram, code string) (*PhoneNumber, error) {
	marshal, err := json.Marshal(struct {
		Code string `json:"code"`
	}{
		Code: code,
	})
	if err != nil {
		return nil, err
	}

	resp, err := ctx.Client.HTTPPost(apiGetUserPhoneNumber, bytes.NewReader(marshal), "application/json;charset=utf-8")
	if err != nil {
		return nil, err
	}

	v := &PhoneNumber{}
	if err = json.Unmarshal(resp, v); err != nil {
		return nil, err
	}
	return v, nil
}

/*
登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。更多使用方法详见 小程序登录。

See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html

GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
*/
func Code2Session(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiCode2Session + "?" + params.Encode())
}

/*
用户支付完成后，获取该用户的 UnionId，无需用户授权。本接口支持第三方平台代理查询。

See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html

GET https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&openid=OPENID
*/
func GetPaidUnionId(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetPaidUnionId + "?" + params.Encode())
}
