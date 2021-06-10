# common-utils
golang common util func

## use
```
 go get github.com/hzy9738/common-utils
```
### Tree
+ new
```
treeData, _ := tree.NewInit(
	tree.SetTreeOriginData(roles),  //roles is slice
)
```
+ getTreeArray
```
pid := 0
treeData.GetTreeArray(pid)
```

### mysql-config
``` 
mysqlString := mysql.GetMysqlString(&mysql.DBConfig{
	Host:     "xxx",
	User:     "xxx",
	Pwd:      "xxx",
	Database: "xxx",
})
db, err := gorm.Open("mysql", mysqlString)  // gorm use
```
### zap-init