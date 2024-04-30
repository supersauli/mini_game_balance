package middleware

import (
	"mini_game_balance/configs"
	"mini_game_balance/internal/pkg/response"
	"mini_game_balance/internal/pkg/utils"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := utils.NewJWT(&configs.ServerConfig.JWT)
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func NewJwt(accountID int64, accType int) string {
	j := utils.NewJWT(&configs.ServerConfig.JWT)
	account := utils.JwtClaim{
		AccID:      accountID,
		AccountStr: strconv.FormatInt(accountID, 10),
		AccType:    accType,
		RandStr:    utils.RandStr(10),
	}
	account.ExpiresAt = time.Now().Add(time.Duration(configs.ServerConfig.JWT.ExpiresTime) * time.Second).Unix()
	token, err := j.CreateToken(account)
	if err != nil {
		return ""
	}
	return token
}
func ParseJwt(token string) (*utils.JwtClaim, error) {
	j := utils.NewJWT(&configs.ServerConfig.JWT)
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func GetAccountIDByJwt(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); exists {
		c := claims.(*utils.JwtClaim)
		return c.AccID
	}
	return 0
}

func GetAccountTypeByJwt(c *gin.Context) int {
	if claims, exists := c.Get("claims"); exists {
		c := claims.(*utils.JwtClaim)
		return c.AccType
	}
	return 0
}

//func GetUUID(c *gin.Context) string {
//	if claims, exists := c.Get("claims"); exists {
//		c := claims.(*utils.JwtClaim)
//		return c.Uid
//	}
//	return ""
//}
//

func CheckJwt(c *gin.Context) bool {
	rawClaim, exists := c.Get("claims")
	if !exists {
		tokenStr := c.Request.Header.Get("Token")
		if tokenStr == "" {
			return false
		}
		var err error
		rawClaim, err = ParseJwt(tokenStr)
		if err != nil {
			zap.L().Error("process token failed, err:", zap.Error(err))
			return false
		}
	}
	claims := rawClaim.(*utils.JwtClaim)
	// 在有效范围
	if claims.ExpiresAt > time.Now().Unix() && claims.Addr == c.ClientIP() {
		return true
	}
	zap.L().Error("process  token failed, expiresat:", zap.Int64("expire", claims.ExpiresAt), zap.String(",new_addr:", c.ClientIP()))
	return false
}
