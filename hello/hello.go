package main

import (
	"fmt"
)

// 主函数
func main(){
	// 调用内置函数 打印
	fmt.Println(" hello world! ");

	fmt.Println(add(3,4));

	// 调用多返回值函数
	a,b,c := duplicate(); // 相当于定义了三个变量同时赋值
	fmt.Println(a,b,c);
	var d,e,f int=duplicate(); // 定义 d e f 三个变量
	fmt.Println(d,e,f);
	var x,y,z =duplicate(); // 定义 x,y,z 三个变量
	fmt.Println(x,y,z);
	fmt.Println( duplicate()); // 实测打印多个返回值会空格分开
}

// 自己定义函数
func add(x int, y int) int{
	// 两数相加
	return x+y;
}

// 多个返回值的函数
func duplicate() (int,int,int){
	return 1,22,3;
}