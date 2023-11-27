package main // 必须有一个main package，GO 会在main package中寻找main函数来启动，没有的话会被报错

import ( // 多重import
	"example/hello/service"
	"fmt" // import的内容需要加引号
)

// 主函数
func main() {
	service.NewServer()
}

func test() {
	// 调用内置函数 打印
	fmt.Println(" hello world! ")

	fmt.Println(add(3, 4))

	// 调用多返回值函数
	a, b, c := duplicate() // 相当于定义了三个变量同时赋值
	fmt.Println(a, b, c)
	var d, e, f int = duplicate() // 定义 d e f 三个变量
	fmt.Println(d, e, f)
	var x, y, z = duplicate() // 定义 x,y,z 三个变量
	fmt.Println(x, y, z)
	fmt.Println(duplicate()) // 实测打印多个返回值会空格分开

	// 尝试调用别的文件的方法
	fmt.Println(Max(1, 2))

	// 循环
	loopTest(50)

	// 面向对象测试
	objTest()

	// 调用本包的类构造对象
	// p := Player{"Player test",20}; 全参构造不需要指定参数名称

	p := Player{name: "Test Obj"} // 部分构造需要指定字段名称
	fmt.Println(p)

	p.sayHi()
	p.sayName()

	service.Search(10086)
	arrTest()
}

// 自己定义函数
func add(x int, y int) int {
	// 两数相加
	return x + y
}

// 多个返回值的函数
func duplicate() (int, int, int) {
	return 1, 22, 3
}
