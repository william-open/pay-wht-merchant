package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pquerna/otp/totp"
	"io/ioutil"
	"mime/multipart"
	"mwhtpay/core/response"
	"time"
)

var VerifyUtil = verifyUtil{}

// verifyUtil 参数验证工具类
type verifyUtil struct{}

func (vu verifyUtil) VerifyJSON(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyJSONArray(c *gin.Context, obj any) (e error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyBody(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBind(obj); err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyHeader(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindHeader(obj); err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyQuery(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyFile(c *gin.Context, name string) (file *multipart.FileHeader, e error) {
	file, err := c.FormFile(name)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

// ValidateTOTPWithRetry 安全的TOTP验证函数（带重试）
func (vu verifyUtil) ValidateTOTPWithRetry(code, secret string) (bool, error) {
	var valid bool
	var err error

	// 重试机制，防止时间同步问题
	for i := 0; i < 3; i++ {
		valid = totp.Validate(code, secret)
		if valid {
			return true, nil
		}

		// 如果不是最后一次尝试，等待一下再重试
		if i < 2 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	return false, err
}
