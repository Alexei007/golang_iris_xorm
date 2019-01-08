package extend

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"golang_iris_xorm/application/common"
	"strings"
	"time"
)

const (
	JWT_SECRET string = "GyJtMRkd98afYy7OTpL5qvc0xolPaoRS"
	JWT_EXPIRE_DATE int64 = 20 // 过期天数
)

type Token struct {
	Token string `json:"token"`
}

type Jwt struct {

}

// 获取用户ID
func (j *Jwt) GetUserId(ctx iris.Context) string {
	// 获取token字段
	authStr := ctx.GetHeader("Authorization")
	if authStr == "" {
		return ""
	}

	// token校验
	token := j.ParseToken(authStr)
	if token == nil {
		return ""
	}

	user_id := token.Claims.(jwt.MapClaims)["sub"].(float64)

	rs := string(int64(user_id))

	return rs
}

// 获取token
func (j *Jwt) GetToken(sub int) (response interface{}) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	// 到期时间
	//claims["exp"] = time.Now().Add(time.Hour * time.Duration(JWT_TTL)).Unix()
	claims["exp"] = time.Now().Unix() + 86400*JWT_EXPIRE_DATE

	// 发布时间
	claims["iat"] = time.Now().Unix()

	// 发行人
	claims["iss"] = "http://api.iris.io"

	// JWT ID 用于标识该JWT
	claims["jti"] = ""

	// 在此之前不可用
	claims["nbf"] = time.Now().Unix()

	// 私有字段(如user_id)
	claims["sub"] = sub
	//claims["user_info"] = model.GetUserInfo()

	token.Claims = claims

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		response = nil
	}

	response = tokenString

	return
}

// ParseToken parse JWT token in http header.
func (j *Jwt) ParseToken(authString string) (t *jwt.Token) {
	//beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		//beego.Error("AuthString invalid:", authString)
		return nil
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		//beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil
			} else {
				// Couldn't handle this token
				return nil
			}
		} else {
			// Couldn't handle this token
			return nil
		}
	}
	if !token.Valid {
		//beego.Error("Token invalid:", tokenString)
		return nil
	}
	//beego.Debug("Token:", token)

	return token
}

// 登录检测
func (j *Jwt) CheckLogin(ctx iris.Context) {
	// 获取token字段
	authStr := ctx.GetHeader("Authorization")
	if authStr == "" {
		ctx.JSON(common.JsonReturn(10004, "用户未登录", nil))
		return
	}

	// token校验
	token := j.ParseToken(authStr)
	if token == nil {
		ctx.JSON(common.JsonReturn(10004, "Token已失效,请重新登录。", nil))
		return
	}

	// 过期校验
	exp := token.Claims.(jwt.MapClaims)["exp"].(float64)
	exp_int := int64(exp)

	if exp_int<time.Now().Unix() {
		ctx.JSON(common.JsonReturn(10004, "Token已失效,请重新登录。", nil))
		return
	}

	/*usersModel := new(models.Users)
	user_info := usersModel.GetUserInfoById(GetUserId(ctx).(int), ctx).(models.Users)
	if user_info.Id == 0 {
		ctx.JSON(JsonReturn(10004, "Token已失效,请重新登录。", nil))
		return
	}

	// 锁定校验
	if user_info.Status == 1 {
		ctx.JSON(ErrReturn(10118,"用户已被禁用"))
	}*/

	ctx.Next()
}