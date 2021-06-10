package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	common_utils "github.com/hzy9738/common-utils"
	"github.com/hzy9738/common-utils/convert"
	"github.com/hzy9738/common-utils/tree"
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
	db.Table("pmmppp_user_rule").Find(&roles)
	treeData, err := tree.NewInit(
		tree.SetTreeOriginData(roles),
	)
	if err != nil {
		fmt.Println(err)
	}

	data := treeData.GetTreeArray(0)
	var treeData1 []RoleTree
	_ = convert.SwapTo(data,&treeData1)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": treeData1,
		})
	})
	r.Run(":5000") // listen and serve on 0.0.0.0:8080
}

type Role struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Pid   int64  `json:"pid"`
}

type RoleTree struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Title     string     `json:"title"`
	Pid       int64      `json:"pid"`
	Childlist []RoleTree `json:"childlist"`
}
