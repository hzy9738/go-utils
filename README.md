# go-utils

golang常用的基础方法

## 如何使用

```
 go get github.com/hzy9738/go-utils
```

### Tree型结构获取

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

### mysql字符串配置

``` 
mysqlString := mysql.GetMysqlString(&mysql.DBConfig{
	Host:     "xxx",
	User:     "xxx",
	Pwd:      "xxx",
	Database: "xxx",
})
db, err := gorm.Open("mysql", mysqlString)  // gorm use
```

### channel聚合查询

```
res, _ := chans.GetResByParamSlice([]string{"1", "2", "3"}, 10, func(param interface{}, resChan chan<- interface{}) {
    resChan <- param })
fmt.Println(res)

```

### 翻转字符串
```
strings.ReverseString("abc")
```