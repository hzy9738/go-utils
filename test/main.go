package main

import (
	_ "github.com/go-sql-driver/mysql"
	common_utils "github.com/hzy9738/common-utils"
	"github.com/jinzhu/gorm"
)

func main() {
	mysqlString := common_utils.GetMysqlString(&common_utils.MysqlConfig{
		Host:     "192.168.1.69",
		User:     "bimeng",
		Pwd:      "yjkj2018",
		Database: "pmmppp",
	})
	db, err := gorm.Open("mysql", mysqlString)
	if err != nil {
		panic(err)
	}
	var roles []Role
	db.Table("pmmppp_auth_rule").Find(&roles)

	tree := common_utils.NewTree(
		common_utils.SetTreeOriginData(roles),
	)
	tree.GetChild()
}

type Role struct {
	Id   int64
	Name string
	Pid  int64
}
