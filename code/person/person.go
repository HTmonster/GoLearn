/*
 * @Author: Theo_hui
 * @Date: 2021-02-25 09:30:47
 * @Descripttion:
 */

package person

// 结构体类型定义
type person struct {
	id   int64
	name string
}

// 工厂方法

func NewPerson(id int64, name string) *person {
	return &person{id, name}
}
