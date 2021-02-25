# Go高层数据结构

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

*XMind - Trial Version*