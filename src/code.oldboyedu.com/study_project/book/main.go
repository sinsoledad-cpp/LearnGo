package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/book/list", bookListHandler)
	r.Run(":8000")
}

func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()

	if err != nil {
		fmt.Println("query default")
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	for _, book := range bookList {
		fmt.Printf("%#v\n", book)
	}
	bookLists := []*Book{
		{Id: 1, Title: "Go语言入门", Price: 399},
		{Id: 2, Title: "Go语言进阶", Price: 499},
	}
	for _, book := range bookLists {
		fmt.Printf("%#v\n", book)
	}
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": bookLists,
	})
}
