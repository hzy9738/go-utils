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
treeData.GetTreeArray("0")
```

### mysql-config
### zap-init