package main

import "fmt"

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

func (s studentMgr) showStudents() {
	for _, stu := range s.allStudent {
		fmt.Printf("id: %d,name: %s\n", stu.id, stu.name)
	}
}

func (s studentMgr) addStudent() {
	var (
		stuId   int64
		stuName string
	)
	fmt.Print("Input id: ")
	fmt.Scanln(&stuId)
	fmt.Print("Input name: ")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuId,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
}

func (s studentMgr) editStudent() {
	var stuId int64
	fmt.Print("Input id: ")
	fmt.Scanln(&stuId)
	stu, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("no this boy")
		return
	}
	fmt.Printf("this boy,id: %d name:%s\n", stu.id, stu.name)
	fmt.Print("new name:")
	var newName string
	fmt.Scanln(&newName)
	stu.name = newName
}

func (s studentMgr) deleteStudent() {
	var stuId int64
	fmt.Print("Input id: ")
	fmt.Scanln(&stuId)
	_, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("no this boy")
		return
	}
	delete(s.allStudent, stuId)
	fmt.Println("delete success!")
}
