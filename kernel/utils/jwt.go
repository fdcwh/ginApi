package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"goGIn/kernel"
	"time"
)

// JwtCustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
type JwtCustomClaims struct {
	// 可根据需要自行添加字段
	JwtSysClaims         `json:"sysData"`
	jwt.RegisteredClaims // 内嵌标准的声明
}

type JwtTokenOut struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtSysClaims struct {
	UserID   int64 `json:"uid"`
	UserName int64 `json:"username"`
}

const (
	ATokenExpiredDuration = 2 * time.Hour
	RTokenExpiredDuration = 30 * 24 * time.Hour
)

// JwtGenToken 生成JWT
func JwtGenToken(JwtSysClaims JwtSysClaims) (JwtTokenOut, error) {
	var err error
	var out JwtTokenOut
	claims := JwtCustomClaims{
		JwtSysClaims,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ATokenExpiredDuration)), // 定义过期时间
			Issuer:    kernel.FdConfig.JWT.Issuer,                                // 签发人
		},
	}
	out.TokenType = "Bearer"
	// HS256、HS384、HS512
	out.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(kernel.FdConfig.JWT.SigningKey))
	if err != nil {
		return out, errors.New(err.Error())
	}

	// RefreshToken
	out.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(RTokenExpiredDuration)), // 定义过期时间
		Issuer:    kernel.FdConfig.JWT.Issuer,                                // 签发人
	}).SignedString([]byte(kernel.FdConfig.JWT.SigningKey))
	if err != nil {
		return out, errors.New(err.Error())
	}
	// 存储 jwtToken 用于验证

	return out, nil
}

// JwtParseToken 解析 token
func JwtParseToken(accessToken string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(kernel.FdConfig.JWT.SigningKey), nil
	})

	if err != nil {
		return nil, errors.New(err.Error())
	}

	// 对token对象中的Claim进行类型断言s
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JwtRefreshToken 刷新
func JwtRefreshToken(refreshToken string) (*JwtCustomClaims, error) {
	var jwtCustomClaims = new(JwtCustomClaims)

	var err error

	return jwtCustomClaims, err
}

/*======================= 存储验证  ================================*/
