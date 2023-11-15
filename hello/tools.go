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