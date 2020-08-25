# Go-json
json是一种轻量级的数据交换格式。在go语言中有以下几种方式进行解析。

结构体定义：[Book](./model/book.go)

# 自带json解析
go自带的json是通过反射进行的解析操作，所以性能有些影响。

[encodingjson](./tools/book_encodingJson.go)

# easyjson
easyjson是一个go的开源库，是一种更快速的json解析方式。

## 使用
终端输入

```bash
easyjson -all <file>.go
```

[easyjson](./model/book_easyjson.go)

# 根据json生成结构体
可以将json文件生成对应的go结构体


[https://mholt.github.io/json-to-go/](https://mholt.github.io/json-to-go/)

# License
[MIT](./LICENSE)

Copyright (c) 2020 golang collection