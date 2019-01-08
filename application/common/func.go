package common

import (
	"encoding/binary"
	"github.com/kataras/iris"
	"golang_iris_xorm/config"
	"net"
	"regexp"
	"strings"
)

type ApiJson struct {
	Code int        `json:"code"`
	Message    interface{} `json:"message"`
	Data   interface{} `json:"data"`
}

func ApiResource(code int, data interface{}, message string) (apijson *ApiJson) {
	apijson = &ApiJson{Code: code, Data: data, Message: message}
	return
}

// 404
func NotFound(ctx iris.Context) {
	ctx.JSON(ApiResource(0, nil, "404 Not Found"))
}

// 成功返回
func SuccReturn(data interface{}) (apijson *ApiJson) {
	if data==nil {
		data = new([0]string)
	}

	apijson = &ApiJson{Code: 1, Data: data, Message: "success"}
	return
}

// 失败返回
func ErrReturn(code int, message string) (apijson *ApiJson) {
	var data [0]string
	apijson = &ApiJson{Code: code, Message: message, Data: data}
	return
}

// json格式化
func JsonReturn(code int, message string, data interface{}) (apijson *ApiJson) {
	if data==nil {
		data = new([0]string)
	}

	apijson = &ApiJson{Code: code, Message: message, Data: data}
	return
}

// 判断是否是邮箱
func IsEmail(email string) bool {
	var reg *regexp.Regexp
	var pattern string

	pattern = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg = regexp.MustCompile(pattern)
	res := reg.FindAllString(email, -1)
	if res != nil {
		return true
	}

	return false
}

// 判断是不是手机
func IsPhone(phone string) bool {
	var reg *regexp.Regexp
	var pattern string

	pattern = `1[\d]{10}`
	reg = regexp.MustCompile(pattern)
	res := reg.FindAllString(phone, -1)
	if res != nil {
		return true
	}

	return false
}

// 图片处理
func ImageHandle(img_url string) string {
	if img_url!="" && strings.Contains(img_url, config.CDN_URL)==false {
		img_url = config.CDN_URL + img_url
	}

	return img_url
}

// 获取ip地址
func GetClientIp(ctx iris.Context) string {
	ip := ctx.RemoteAddr()

	si := strings.Split(ip, ":")

	return si[0]
}

// ip to long
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

// long to ip
func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

func InArray(need interface{}, needArr []interface{}) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}