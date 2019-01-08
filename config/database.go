package config

const (
	DriverName string = "mysql"
	Username string = "test_db"
	Password string = "123456"
	Host string = "127.0.0.1"
	Port string = "3306"
	DbName string = "test_db"
	DbCharset string = "utf8"
	DbPrefix string = "tt_"
	ShowSQL bool = true
)

var DataSourceName string = Username+":"+Password+"@tcp("+Host+":"+Port+")/"+DbName+"?charset="+DbCharset