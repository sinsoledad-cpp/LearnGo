package main

import (
	"fmt"
	"os"
)

var smr studentMgr

func showMenu() {
	fmt.Println("welcome sms!")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Println("your choice:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("your choice: ", choice)
		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("please input again~~~")
		}
	}
}
