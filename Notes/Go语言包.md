# Go语言包

## 包

### 标准包 150多个

- unsafe

	- 打破类型安全的命令
	- 用在 C/C++程序调用中

- os

	- 平台无关的系统操作

- os/exec

	- 运行外部操作系统命令和程序

- syscall

	- 操作系统底层调用的接口

- archive/tar

	- 解压缩

- /zip-compress

	- 解压缩

- fmt

	- 格式化输入输出

- io

	- 基本输入输出的封装  基本是底层功能的封装

- bufio

	- 缓冲输入输出

- path/filepath

	- 系统中文件名路径分析

- flag 

	- 命令行参数操作

- strings

	- 字符串的操作

- striconv

	- 字符串转为基本类型

- unicode

	- unicode字符串操作

- regexp

	- 正则表达式操作

- bytes

	- 字符型分配操作

- index/suffixarray

	- 子字符串快速查询

- math

	- 基本数学函数

- math/cmath

	- 数学复数操作

- math/rand

	- 伪随机数操作

- sort

	- 数组排序

- math/big

	- 大数计算和实现

- list

	- 双向列表

- ring

	- 环形列表

- time

	- 时间

- log

	- 日志

- encoding/json

	- json数据操作

- encoding/xml

	- xml数据操作

- text/template

	- 数据与文本混合的数据模板

- net

	- 网络数据基本操作

- http

	- 可拓展的http服务器和客户端

- html

	- html解析

- runtime

	- go程序运行交互操作

- reflect

	- 反射

## 时间和日期

### time 包

- 数据类型 time.Time
- 获得当前时间  time.Now()
- 部分时间

	- t.Day()
	- t.Minute()
	- t.Year()
	- ...

## 正则表达式

### 正则表达式匹配

- 字符串匹配

	- ok,_ := regexp.Match(pat,[]bytes(searchIn))
	- ok,_  := regexp.MatchString(pat,searchIn)

*XMind - Trial Version*