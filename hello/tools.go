package main

import "fmt"


func Max(x,y int) int{
	fmt.Println("this is Max function in tools.go"); // import 的内容必须使用
	if x>y {
		return x;
	}else {
		return y;
	}
}

func loopTest(cnt int) { // 无返回值的函数直接不写返回值类型
	for x :=0; x<cnt; x++ { // 循环条件不需要加括号
		fmt.Printf("这是第%d次循环",x)
		fmt.Println();
	}
}