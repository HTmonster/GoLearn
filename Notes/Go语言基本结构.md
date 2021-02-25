# Go语言基本结构

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

## 常量

### 关键词

- const

### 格式

- const identifier [type] = value
- const Pi = 3.14159

### 要求

- 编译时刻值要确定
- 除了内置函数 例如len()

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

*XMind - Trial Version*