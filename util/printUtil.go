package util

import (
	"fmt"
	"go-sort-algorithm/model"
)

func PrintMenu() model.ChoiceType {
	fmt.Println("************学生管理系统************")
	fmt.Println("1、增加")
	fmt.Println("2、查询")
	fmt.Println("3、插入")
	fmt.Println("4、排序")
	fmt.Println("0、退出")
	fmt.Print("Please input your choice：")

	var choiceType model.ChoiceType
	fmt.Scanln(&choiceType)
	return choiceType
}

func PrintAddMenu () model.AddChoiceType {
	fmt.Println("************学生管理系统 增加学生信息************")
	fmt.Println("1、单个学生信息录入")
	fmt.Println("2、批量学生信息导入")
	fmt.Println("0、返回上一层")
	fmt.Print("Please input your choice：")

	var choiceType model.AddChoiceType
	fmt.Scanln(&choiceType)
	return choiceType
}


func PrintQueryMenu () model.QueryChoiceType {
	fmt.Println("************学生管理系统 查询学生信息************")
	fmt.Println("1、查询所有学生信息")
	fmt.Println("2、根据学号查询")
	fmt.Println("3、根据姓名查询")
	fmt.Println("0、返回上一层")
	fmt.Print("Please input your choice：")

	var choiceType model.QueryChoiceType
	fmt.Scanln(&choiceType)
	return choiceType
}

func PrintSortMenu () model.SortChoiceType {
	fmt.Println("************学生管理系统 排序菜单************")
	fmt.Println("1、冒泡排序")
	fmt.Println("2、希尔排序")
	fmt.Println("3、快速排序")
	fmt.Println("4、堆排序")
	fmt.Println("0、返回上一层")
	fmt.Print("Please input your choice：")

	var choiceType model.SortChoiceType
	fmt.Scanln(&choiceType)
	return choiceType
}