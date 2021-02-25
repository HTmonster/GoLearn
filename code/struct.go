/*
 * @Author: Theo_hui
 * @Date: 2021-02-25 09:09:44
 * @Descripttion:
 */

package main

import (
	"fmt"
	"reflect"
	// "person"
)

// 类型定义
type T struct {
	i   int
	f   float32
	str string
}

// 带tag的类型定义
type Ttag struct {
	id   int    "ID"
	name string "名字"
}

//使用reflect来访问标签
func refTag(t Ttag, ix int) {
	ttype := reflect.TypeOf(t)      //获得类型
	ixfield := ttype.Field(ix)      //获得类型中的字段
	fmt.Printf("%v\n", ixfield.Tag) //获得字段中的标签
}

// 主函数
func main() {

	// 使用new函数来分配内存
	t := new(T)

	// 字段分配值
	t.i = 12
	t.f = 0.234
	t.str = "测试字符串"

	// 字段值访问
	fmt.Printf("t's i is %d,f is j%f,str is %s", t.i, t.f, t.str)
	fmt.Println(t)

	// 快速的分配赋值方式
	t2 := &T{1, 2.0, "hello"}
	fmt.Println(t2)

	// 使用工厂方法来创建
	// json := person.NewPerson(1, "json")
	// fmt.Println(json)

	// 带标签的结构体
	tt := Ttag{1, "hh"}
	refTag(tt, 0) //访问第一个字段的标签
	refTag(tt, 1)

}
