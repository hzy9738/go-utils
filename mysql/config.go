package mysql

import (
	"fmt"
)

type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlString(mysqlInfo *DBConfig) string {
	if mysqlInfo.Port == 0 {
		mysqlInfo.Port = 3306
	}
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysqlInfo.User, mysqlInfo.Pwd, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Database)
}
