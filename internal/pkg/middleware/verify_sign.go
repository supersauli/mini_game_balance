package middleware

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifySign(use bool, token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !use {
			return
		}

		var body []byte
		var err error
		if c.Request.Method == "POST" {
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.Abort()
				return
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		} else {
			body = []byte(c.Request.URL.RawQuery)
		}
		reqSignString := c.GetHeader("signValue")
		if len(reqSignString) == 0 {
			c.Abort()
			return
		}
		reqTime := c.GetHeader("requestTime")
		if len(reqTime) == 0 {
			c.Abort()
			return
		}
		reqTimeInt, err := strconv.ParseInt(reqTime, 10, 64)
		if err != nil {
			c.Abort()
			return
		}

		//if time.Now().Unix()-reqTimeInt/1000 > 3600 {
		//	logging.Info("VerifySign request is timeOut timeNow ", time.Now().Unix(), "reqTime", reqTimeInt)
		//	c.Abort()
		//	return
		//}

		tempByte := []byte(fmt.Sprintf("%d", reqTimeInt))
		tempByte = append(tempByte, []byte(token)...)
		tempByte = append(tempByte, body...)

		md5Byte := md5.Sum(tempByte)
		tempSignString := strings.ToUpper(fmt.Sprintf("%x", string(md5Byte[:])))
		if reqSignString != tempSignString {
			c.Abort()
		}
	}
}
