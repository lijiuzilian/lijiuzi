[![pipeline status](https://api.travis-ci.org/33cn/plugin.svg?branch=master)](https://travis-ci.org/33cn/plugin/)
[![Go Report Card](https://goreportcard.com/badge/github.com/33cn/plugin?branch=master)](https://goreportcard.com/report/github.com/33cn/plugin)


# 基于区块链开发的 lijiuzi公有链系统


### 编译

```
git clone https://github.com/lijiuzilian/lijiuzi $GOPATH/src/github.com/lijiuzilian/lijiuzi
cd $GOPATH/src/github.com/lijiuzilian/lijiuzi
go build -i -o lijiuzi
go build -i -o lijiuzi-cli github.com/lijiuzilian/lijiuzi/cli
```

### 运行
拷贝编译好的lijiuzi, lijiuzi-cli, lijiuzi.toml这三个文件置于同一个文件夹下，执行：
```
./lijiuzi -f ocean.toml
```
