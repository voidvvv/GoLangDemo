package main

import "fmt"

// 测试面向对象
type Player struct{
	name string; 
	// var name string; 类（结构体）中的属性前不加 var
	age int;
}

// 给结构体（类） 定义方法，目前看起来应该是在外面定义
func (t *Player) sayHi(){
	fmt.Println("this is player say hi")
}

// 给结构体（类） 定义方法，尝试使用对象属性
func (t *Player) sayName(){
	
	fmt.Printf("this is Player: %s",t.name);
}

func objTest(){
	var p Player;
	p2 := Player{"Player02",20};

	p.name = "Player01";
	p.age = 20;

	fmt.Println(p); // 测试直接打印对象
	fmt.Println(p2);
}

// 尝试结构体继承
