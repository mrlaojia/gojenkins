# 这是一个 fork 项目

使用时，请使用 replace 替代源仓库（只要指定版本）

```shell
module jjenkins

go 1.19

require github.com/bndr/gojenkins v1.1.0

replace (
	github.com/bndr/gojenkins v1.1.0 => github.com/mrlaojia/gojenkins v1.5.0
)
```
