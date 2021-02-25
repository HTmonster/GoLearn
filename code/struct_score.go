/*
 * @Author: Theo_hui
 * @Date: 2021-02-25 10:21:05
 * @Descripttion:
 */

package main

import "fmt"

//定义成绩类型
type score struct {
	math    int "数学成绩"
	English int "英语成绩"
	Biology int "生物成绩"
	History int "历史成绩"
}

// 方法
func (s score) sumAll() int {
	// 计算所有的成绩
	return s.math + s.Biology + s.English + s.History
}

func (s score) sumScience() int {
	// 计算理科成绩
	return s.math + s.English + s.Biology
}

func (s score) sumMathPlus(plus int) int {
	// 计算加分后的数学成绩
	return s.math + plus
}

func main() {
	// 创建一个成绩
	s := score{99, 78, 67, 56}
	fmt.Println("score sum:%d", s.sumAll())
	fmt.Println("score Science sum:%d", s.sumScience())
	fmt.Println("score Math sum plus:%d", s.sumMathPlus(23))
}
