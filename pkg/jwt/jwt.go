package jwtUtil

import (
	"PlanetMsg/pkg/signalInfo"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var (
	ATokenExpiredDuration time.Duration
	RTokenExpiredDuration time.Duration
	signKey               []byte
	TokenIssuer           string
)

type UserClaim struct {
	UserID  int32  `json:"user_id"`
	Account string `json:"account"`
	jwt.RegisteredClaims
}

func init() {
	viper.SetConfigName("jwt")
	viper.AddConfigPath("../../../conf")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}
	// 读取jwt配置文件
	AHour, err := strconv.Atoi(viper.GetString("jwt.ATokenExpiredDuration"))
	if err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("%s", err))
	}
	RHour, err := strconv.Atoi(viper.GetString("jwt.RTokenExpiredDuration"))
	if err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("%s", err))
	}
	// 定义访问令牌（Access Token）和刷新令牌（Refresh Token）的有效时间
	ATokenExpiredDuration = time.Duration(AHour) * time.Hour
	RTokenExpiredDuration = time.Duration(RHour) * time.Hour
	// Token的签名
	signKey = []byte(viper.GetString("jwt.signKey"))
	TokenIssuer = ""
}

func getJWTTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return signKey, nil
}

// 颁发访问令牌（Access Token）和刷新令牌（Refresh Token）
func GenToken(userId int32, account string) (atoken string, rtoken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: getJWTTime(ATokenExpiredDuration),
		Issuer:    TokenIssuer,
	}
	uc := UserClaim{
		UserID:           userId,
		Account:          account,
		RegisteredClaims: rc,
	}
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, uc).SignedString(signKey)
	// refresh token 不需要保存保存任何用户信息
	rt := rc
	rt.ExpiresAt = getJWTTime(RTokenExpiredDuration)
	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rt).SignedString(signKey)
	return
}

// VerifyToken 验证Token，并获取
func VerifyToken(tokenStr string) (*UserClaim, error) {
	var uc = new(UserClaim)
	token, err := jwt.ParseWithClaims(tokenStr, uc, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		err = errors.New("verify Token Failed")
		return nil, err
	}
	return uc, nil
}

// 通过刷新令牌（Refresh Token）重新获取访问令牌（Access Token）
func RefreshToken(atoken, rtoken string) (newAtoken, newRtoken string, err error) {
	// 如果rtoken无效则返回 "RToken已超时"
	if _, err = jwt.Parse(rtoken, keyFunc); err != nil {
		return "", "", errors.New(signalInfo.GetMsg(signalInfo.ERROR_RAUTH_TIMEOUT))
	}
	// 从旧访问令牌中解析出claims数据
	var claim UserClaim
	_, err = jwt.ParseWithClaims(atoken, &claim, keyFunc)
	// 如果错误是有访问令牌过期导致，则重新生成令牌
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claim.UserID, claim.Account)
	}
	return
}
