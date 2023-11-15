package service

import "fmt"

// 这里测试别的package。需要注意，这里的方法首字母需要大写，这样才能被别的package可见
func Search(id int) (string){
	fmt.Printf("search by id: %d",id);
	return "Target";
}
/*
	// 需要设置 %GOPATH%, 并且需要把所有引用的代码放在 %GOPATH%/src 下
// 并且引入的时候，
	需要import 包在 %GOPATH%/src 下的文件路径，如果包名与文件夹名称不一致
	这样使用起来的时候就会及其的怪异,个人感觉对于开发来说及其不友好

*/