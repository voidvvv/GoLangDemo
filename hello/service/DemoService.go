package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	导包的时候，需要从当前的模块名称开始，使用/一步一步往下找.
	比如，我现在模块名称为 hello/test
	那么我要找到service子目录下的包，就需要这么写：
	import "hello/test/service"

	然后代码中就可以使用service包下的内容了

*/

// 这里测试别的package。需要注意，这里的方法首字母需要大写，这样才能被别的package可见
func Search(id int) string {
	fmt.Printf("search by id: %d", id)
	return "Target"
}

type Book struct {
	Name  string `json:"name"` // 使用tag来表示json序列化后的字段名称
	Price int    `json:"price"`
}

func (p *Book) getName() string {
	return p.Name
}

func (p *Book) getPrice() int {
	return p.Price
}

/*
*
测试rest API
*/
var bookList = []Book{
	{"JAVA入门", 20},
	{"GOLANG入门", 30},
}

// 写一个接口，用来查找所有的book
func GetBookList(c *gin.Context) {
	// 这里将对象写入响应体中
	c.IndentedJSON(http.StatusOK, bookList)
}

// 写一个方法，用来新增一本书
func AddNewBook(c *gin.Context) {
	var book Book

	// 使用BindJson方法，将请求体中的json数据绑定到指定对象中，这里需要传入指针
	if err := c.BindJSON(&book); err != nil {
		return
	}

	// 这里使用append方法向数组中插入数据。但是要注意，这里的数组在定义时不能指定大小，否则会报错
	// 还需要注意，这里的append方法，其实时返回了一个新的数组对象，与原来的不同
	bookList = append(bookList, book)

	c.IndentedJSON(http.StatusOK, bookList) // 写入相应数据

}

func NewServer() {
	router := gin.Default()
	router.GET("/bookList", GetBookList) // 绑定接口对应的方法
	router.POST("/addBook", AddNewBook)  // 绑定接口对应的方法

	router.Run("localhost:8080")
}
