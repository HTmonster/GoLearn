# go语言综合

## 基本结构

### 函数

- 格式 

	- func functionName(parameter_list) (return_value_list) {
   …
}

- main函数

	- 无参数
	- 无返回值

### 注释

- 单行注释 //
- 多行注释  /*  */
- 包注释

	- 开头  作为文档

- 全局作用域注释

	- 函数前面 函数名字开头

### 类型

- 基本类型

	- int 
	- float
	- bool
	- string

- 复合类型

	- struct
	- array
	- slice
	- map
	- channel

- 只描述类型的行为

	- interface

- 自定义类型

	- type

### 一般结构

- import 引入包
- 对常量，变量，类型的定义或声明
- init函数（如果有）
- main函数 （main 包）
- 其余函数

### 类型转换

- valueOfTypeB = typeB(valueOfTypeA)
- a := 5.0
b := int(a)

## 背景

### 起源

- 2007年雏形设计
- 2009年首次公布
- 2010年 Google投入使用
- google 20%兼职项目

### 特性

- 基础：C系编程语言
- 声明和包设计：Pascal,Modula,Oberon语言
- 并发原理：Tony Horare CSP
- 动态语言特性：Python 和 Ruby

### 设计理念

- 快速编译 （C++运行快但是编译速度不理想）
- 高效执行  （Java编译快 但是执行不理想）
- 易于开发 （动态语言开发容易 但是执行速度慢）

### 发展目标

- 类型安全和内容安全
- 对网络通信，并发和并行的支持
- 构建速度快
- 包模型管理

### 语言特性

- 没有类和继承，使用接口完成多态
- 函数是基本构件
- 静态类型，强类型语言
- 支持交叉编译

### 用途

- Web服务器 存储集群 
- 复杂时间处理  CEP
- 子主题 3

### 缺点

- 不支持函数重载和操作符重载
- 不支持隐式转换
- 不支持变体类型
- 不支持动态加载
- 不支持动态链接
- 不支持泛型
- 不支持断言
- 不支持静态变量

## 环境

### 编译器

- 原生编译器gc

	- 基于Plan 9操作系统使用的C工具链、
	- 不同操作系统和处理器架构区别对待
	- Go 1.0.3

		- go build
		- go install

- gccgo编译器

	- 更加传统的编译器  gcc作为后端
	- 编译速度较gc慢

### 文件拓展名和包

- 源文件拓展名.go
- 包拓展名 .a

### 环境变量

- $GOROOT  安装位置
- $GOARCH 目标机器的处理器架构

	- 386
	- amd64
	- arm

- $GOOS  目标机器操作系统

	- darwin
	- freebsd
	- linux
	- windows

- $GOBIN 编译器和连接器的安装位置
- $GOPATH 
- $GPARM 针对arm架构的处理器
- $GOMAXPROCS 设置应用程序可以使用的处理器核数和个数

## 常量

### 关键词

- const

### 格式

- const identifier [type] = value
- const Pi = 3.14159

### 要求

- 编译时刻值要确定
- 除了内置函数 例如len()

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

## Map

### 概念

- 引用类型
- 初始化

	- 初始化方法

		- map literals

			- maplit = map[string]int{"one":1,"two":2}

		- 引用类型

			- mapcreat:=make(map[string]int)

		- 不要使用new

	- 未初始化

		- nil

- 声明与赋值

	- 要求

		- key

			- 可以：int float  string   指针  接口类型等
			- 可以：结构体

				- key()
				- hash()

			- 不可以：数组  切片 结构体

		- value

			- 任意值

				- 切片（一对多）

					- map1  := make(map[int][]int)
					- map1 := make(map[int]*[]int)

			- 函数

	- 值赋值

		- map1[key] =value
		- val :=map[key]

	- var map map[keytype] valuetype

- 容量

	- 容量不定
	- 动态伸缩
	- 可以在make时候进行指定

		- mapcreated := make(map[string]int,100)

	- 到达容量上限会自动增加
	- 出于性能考虑

		- 尽量表明容量

- 存在判断

	- value,IsPresent = map[key]

- 删除

	- delete(map1,key1)

- 循环

	- for key value := range map{ }

- map 切片

	- 1. 分配切片

		- maplist  := make([] map[string]int,100)

	- 2. 对切片的每个元素进行分配值

		- for i := range mapList{
         mapList[i]=map[string]int{"one":1,"two":2}
	}

- map 排序

	- 无序的
	- 解决方案

		- 复制到切片中
		- 在切片中进行排序

## 变量

### 关键词

- var

### 格式

- var identifier type

### 特点

- 自动赋值为0值

### 命名规则

- 骆驼命名法

### 作用域

- 全局变量
- 局部变量

## 基本类型

### 布尔类型

- var b bool = true
- 逻辑运算符

	- 与 &&
	- 或 ||
	- 非 ！

### 数字类型

- 基于架构的类型

	- int uint

		- 64位  64bit
		- 32位 32bit

	- uintptr

		- 足够存放一个指针即可

- 与操作系统架构无关类型

	- 有符号整数

		- int8
		- int16
		- int32
		- int64

	- 无符号整数

		- uint8
		- uint16
		- uint32
		- uint 64

	- 浮点数

		- float32
		- float64

- 进制

	- 077 八进制
	- 0xFF 十六进制

- 复数

	- complex64
	- complex128

- 位运算

	- 要求

		- 整数
		- 等位长

	- 二元运算符

		- 按位与 &
		- 按位或 ！
		- 按位异或 ^

	- 一元运算符

		- 按位补足 ^
		- 位左移 <<
		- 位右移 >>

### 字符类型

- ASCII字符

	- byte

- Unicode字符

	- int \u十六进制

### 字符串类型

- 特性

	- 值类型
	- 值不可变

- 分类

	- 解释字符串

		- 双引号

	- 非解释字符串

		- 反引号

## 时间和日期

### time 包

- 数据类型 time.Time
- 获得当前时间  time.Now()
- 部分时间

	- t.Day()
	- t.Minute()
	- t.Year()
	- ...

## 指针

### 控制指定集合

- 数据结构
- 分配的数量
- 内存访问模式

### 取地址

- &

### 空指针

- nil

### 不能发得到一个文字 或者常量的地址

## 控制结构

### if-else结构

- if initialization; condition {
    // do something 
}
- if condition {
    // do something 
} else {
    // do something 
}
- if condition1 {
    // do something 
} else if condition2 {
    // do something else    
} else {
    // catch-all or default
}
- 格式要求

	- 左括号和条件同行
	- 右括号和下一级同行

### switch 结构

- switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
- 特性

	- var 可以是任何类型
	- 可以为表达式
	- 各个直接的类型相同
	- 不需要单独使用break

		- fallthrough 关键字接着执行

### for 结构

- for 初始化语句; 条件语句; 修饰语句 {}
- 无限循环  for{}
- for-range

	- for pos, char := range str {
...
}

### break continue

### 标签与goto

## 函数

### 类型

- 普通带名字函数
- 匿名函数 lambda函数
- 方法 Methods

### 值传递

- 按值传递  （默认）
- 按引用传递

	- 通过指针

### 变长参数

- ...type

### defer和追踪

- defer关键字

	- 允许我们推迟函数返回前执行某个语句或函数

- 多个defer 

	- 逆序执行 

- 常见情境

	- 关闭文件
	- 解锁加锁资源
	- 打印最终报告
	- 关闭数据库链接

- 代码追踪

	- func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

- 回调函数

### 闭包

- 匿名函数

	- 赋值给变量

		- fplus := func(a,b int) int {return x+y}

	- 直接调用

		- func(x, y int) int { return x + y }

- 函数作为返回值

	- func main() {
    // make an Add2 function, give it a name p2, and call it:
    p2 := Add2()
    fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
    // make a special Adder function, a gets value 2:
    TwoAdder := Adder(2)
    fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
    return func(b int) int {
        return b + 2
    }
}

func Adder(a int) func(b int) int {
    return func(b int) int {
        return a + b
    }
}

- 闭包调试

	- 使用runtime.Caller()函数

		- 	where :=func(){
		_,file,line,_:=runtime.Caller(1)
		log.Printf("%s:%d",file,line)
	}

	a:=1
	where()
	print(a)
	where()

	- 使用log.Print函数

		- 	where :=log.Print

	a:=1
	where()
	print(a)
	where()

## 数组和切片

### 概念

- 相同唯一类型

	- 原始类型

- 编号
- 长度固定

### 定义

- var identifier [len]type

### new创建

- var arr=new([5]int)

	- arr类型为 *[5]int

### 切片

- 长度可变的数组
- 计算容量 cap()
- 定义

	- var identifier []type

## 可见性规则

### 标识符大写开头   外部包可使用

### 标识符小写开头  外部包不可使用

## 内置函数

### close()

- 关闭管道

### len()

- 获取长度

### cap()

- 获取容量

### new()

- 分配类型的零值 返回地址

### make()

- 返回类型的初始化之后的值 （对于 切片  map 管道）

### copy()

- 复制切片

### append()

- 连接切片‘

### panic()

- 错误处理

### recover()

- 错误处理

### print()  println()

- 底层打印函数

### complex()  real() imag()

- 创建和操作复数

## strconv

### int位数

- strconv.InteSize

### 数字转字符串

-  strconv.Itoa(i int) string
- strconv.FormatFloat( f float64,fmt byte, prec int,bitSize int)

### 字符串转数字

- strconv.Atoi(s string) (i int,err error)
- strconv.ParseFloat(s string, bitSize int)(f float64, err error)

## strings

### 前缀和后缀

- strings.HasPrefix(s, prefix string) bool
- strings.HasSuffix(s, suffix string) bool

### 包含

- strings.Contains(s,substr string) bool

### 索引

- strings.Index(s,str string) int
- strings.LastIndex(s,str string) int
- 非ASCII 编码字符串 strings.IndexRune(s string,r  rune) int

### 替换

- strings.Replace(str .old,new,n) string

### 统计

- strings.CountA(s,str string) int

### 重复

- strings.Repeat(s,count int) string

### 大小写修改

- strings.ToLower(s) string
- strings.ToUpper(s) string

### 修剪

- 首尾空格  strings.TrimSpace(s) 
- 首尾 指定字符串 strings.Trim(s,str)
- 首指定字符串 strings.TrimLefts,str)
- 尾指定字符串 strings.TrimRight(s,str)

### 分割

- 空白字符分割  strings.Fields(s)
- 自定义字符分割 strings.Split(s, sep)

### 拼接

- strings.join(sl []string,sep string) string

### 内容读取

- strings.NewReader(str)

	- Read() 读取内容从 [] vyte
	- ReadByte()
	- ReadRune()

## 可以隐式声明

## 任意类型：空接口

## := 赋值操作符  （初始化声明）

*XMind - Trial Version*