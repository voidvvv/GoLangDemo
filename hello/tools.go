package main

import "fmt"

func Max(x, y int) int {
	fmt.Println("this is Max function in tools.go") // import 的内容必须使用
	if x > y {
		return x
	} else {
		return y
	}
}

func loopTest(cnt int) { // 无返回值的函数直接不写返回值类型
	for x := 0; x < cnt; x++ { // 循环条件不需要加括号
		fmt.Printf("这是第%d次循环", x)
		fmt.Println()
	}
}

func arrTest() {
	// 1. 声明数组
	var arr01 [10]int
	arr01[0] = 0

	// 2. 声明数组顺带初始化
	var arr02 = [10]int{1, 2, 3} // 剩下的会用0填充
	fmt.Println(arr02)

	// 如果不确定长度
	var arr03 = [...]int{1, 2, 3} // 自动判断数组长度

	lenStr := fmt.Sprint(len(arr03))
	fmt.Println("数组长度：" + lenStr)
}
