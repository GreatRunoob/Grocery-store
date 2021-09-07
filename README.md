# Grocery-store
欢迎光临***杂货铺***，恰如其名，所有收录至本仓库的代码均为个人平时闲暇或学习所作！
其中绝大部分是实现简单功能的脚本工具、代码片段或教学模型，并在原项目的基础上尽可能优化代码逻辑，修复了一些bug。
一方面供学习所用，另一方面留存作为记录。有时间的话也会不定期维护仓库内容，但也仅限于优化代码逻辑与性能。(づ￣ 3￣)づ


## file-search-go
一个用于在指定文件路径下，进行模糊查找的小工具。
在不借助预先构建索引表的情况下，能够较为高效地按照给定关键词，模糊匹配目录与文件存储位置信息。

"烹饪指南"：
- 编译 `go build ./file-search-go/search.go`
- 运行 `./search`

参考教学视频链接：
- **[用多线程给代码提速800% —— Golang高并发教程+实战](https://www.bilibili.com/video/BV1qT4y1c77u)**


## filename-prefix-processor-go
一个用于批量重命名文件的小工具，**Golang**重构版本。
在指定文件路径下，查找具有特定前缀名称的直属文件，通过剔除共同前缀的方式，达到批量重命名文件的作用。

"烹饪指南"：
- 编译 `go build ./filename-prefix-processor-go/processor.go`
- 运行 `./processor`

原项目链接：
- **[Github](https://github.com/chemicalfiber/file-name-prefix-processor)**
- **[Gitee](https://gitee.com/chemicalfiber/file-name-prefix-processor)**

参考教学视频链接：
- **[一个5KB的小工具 不止能省5分钟! 带你手写一个文件名处理工具 批量文...](https://www.bilibili.com/video/BV1Bq4y1K7Uz)**


## filename-prefix-processor-py
一个用于批量重命名文件的小工具，**Python**重构版本。
在指定文件路径下，查找具有特定前缀名称的直属文件，通过剔除共同前缀的方式，达到批量重命名文件的作用。

"烹饪指南"：
- 环境要求 `python >= 3.5`
- 运行 `python3 ./filename-prefix-processor-go/processor.go`

原项目链接：
- **[Github](https://github.com/chemicalfiber/file-name-prefix-processor)**
- **[Gitee](https://gitee.com/chemicalfiber/file-name-prefix-processor)**

参考教学视频链接：
- **[一个5KB的小工具 不止能省5分钟! 带你手写一个文件名处理工具 批量文...](https://www.bilibili.com/video/BV1Bq4y1K7Uz)**

## message-go
一个简单的即时通讯系统教学模型，实现在线聊天室基础功能，包括在线广播、私聊、修改用户在线名称等。

"烹饪指南"：
- `server/`
  - 即时通讯系统服务端
  - 编译 `go build -o server main.go server.go user.go`
  - 运行 `./server -ip xxx.xxx.xxx.xxx -port xxx` 默认监听本地8888端口
- `client/`
  - 即时通讯系统客户端
  - 编译 `go build -o client main.go client.go`
  - 运行 `./client -ip xxx.xxx.xxx.xxx -port xxx` 默认监听本地8888端口

参考教学视频链接：
- **[8小时转职Golang工程师(如果你想低成本学习Go语言) P37~P51](https://www.bilibili.com/video/BV1gf4y1r79E)**


## raspi-tools
包含若干实用的树莓派4B **Python**脚本工具
- `date.py` 提供日期格式化方法
- `monitor.py` 提供若干有关树莓派4B硬件性能指标读取方法
- `startMonitor.sh` 用于启动监控脚本
- `smartFan.py` 树莓派温控风扇脚本
- `temperature.py` 树莓派4B SoC温度监控脚本


## sort-go
使用Golang实现经典算法模型，不定期更新中...
- `selection-sort.go` 选择排序算法
- `binary-search.go` 二分查找算法


## tcp-scanner
用Golang实现一个简单的TCP端口扫描器教学模型，麻雀虽小但五脏俱全，借助并发的强大特性，使得扫描器的效率异常高效

"烹饪指南"：
- 编译 `go build -o scanner ./tcp-scanner/main.go`
- 运行 `./scanner` 默认扫描localhost:1~1024
- 接受命令行参数 `./scanner [-ip|-from|-to]`
- e.g. `./scanner -ip 192.168.1.120 -from 22 -to 8080`

参考教学视频链接：
- **[Go语言（Golang）编写 TCP 端口扫描器](https://www.bilibili.com/video/BV13K4y1a7dt)**


## limiter-go
用Golang实现服务请求限流器模型

参考教学视频链接：
- **[手撸限流器，扛不住我还限不住嘛？](https://www.bilibili.com/video/BV1iq4y1p7VC)**
