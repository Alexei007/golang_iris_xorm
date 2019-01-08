package config

import "time"

// 是否关闭启动日志
const DisableStartupLog bool = false

// 时间格式化字符串
const TimeFormat string = "2019-01-01 00:00:00"

// 编码
const Charset string = "UTF-8"

// http限流 - 每秒最多请求5次
const Limiter = 5

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")