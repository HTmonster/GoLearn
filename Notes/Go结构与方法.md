# Go 结构与方法

## 类型定制

### 类型别名

### 结构体

- 字段构成

	- 名字（唯一）
	- 类型

## 结构体

### 定义

- type identifier struct {
    field1 type1
    field2 type2
    ...
}
- 支持匿名字段

	- type identifier struct {
    field1 type1
    type2
    ...
}

- 支持内嵌结构体

	- type A struct {
    a  int
} 
type B struct {
    s  str
    A
}

### 初始化

- 混合字面量语法

	- ms := &struct1{10, 15.5, "Chris"}
	- intr := Interval{0, 3}            (A)
intr := Interval{end:5, start:1}  (B)
intr := Interval{end:5}           (C)

### 字段赋值

- var s T
s.a = 5
s.b = 8
- 无论是值或者指针，都可以使用选择器（.)来进行访问

### new函数

- 为结构体分配内存并返回指针

	- var t *T = new(T)

### 使用工厂方法创建实例

### 标签（可选）

- 附属于字段的字符串
- 使用reflect包进行访问

	- 在运行时自省类型、属性和方法

- 子主题 3

## 方法

### 简介

- 作用在接收者（receiver) 上
- 特殊函数

### 定义

- func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
- 接收者

	- 除了接口类型之外都可以

- 不使用recv 的值，可以用 _ 替换它

### 方法与函数的区别

- 变量

	- 方法被变量调用
	- 函数将变量作为参数

### 可以用来定制String() 方法

*XMind - Trial Version*