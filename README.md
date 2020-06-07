# goutil 

一些自己经常用到的函数集

### 列表

- CreateParentDir 递归创建上级目录
- FindInSet 类似mysql的FIND_IN_SET
- InInt64Slice 判断数字是否是[]int64的一个元素
- InIntSlice 判断数字是否是[]int的一个元素
- InStringSlice 判断字符串是否是[]string的一个元素
- IsExist 判断文件或目录是否存在
- JSON 获取json字符串
- JoinInt64Slice 类似strings.Join，不过这里是用于[]int64
- JoinIntSlice 类似strings.Join，不过这里是用于[]int
- Mkdir 创建目录
- RandomString 生成随机字符串
- String2Int64Slice 将"1,2,3"这种形式的字符串变成[]int64{1,2,3}
- String2IntSlice 将"1,2,3"这种形式的字符串变成[]int{1,2,3}
- RunesFilter 过滤字符集
