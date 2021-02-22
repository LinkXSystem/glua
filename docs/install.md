# 环境

## 项目开发的前置环境

```shell
brew install lua

brew install go
```

## 设置 Go 的环境

<!-- TODO：Go Module 的方式待定 -->

```shell
# 关闭 go module
go env -w GO111MODULE=off

# root 表示项目所在的目录
go env -w GOPATH=${root}/glua
```
