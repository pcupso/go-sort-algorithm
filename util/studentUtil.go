package util

import (
	"fmt"
	"go-sort-algorithm/model"
)

func InputStudent() *model.StuNode {
	stu := model.Student{}
	fmt.Print("请输入学生学号: ")
	fmt.Scanln(&stu.Id)
	fmt.Print("请输入学生姓名: ")
	fmt.Scanln(&stu.Name)
	fmt.Print("请输入学生第一科成绩: ")
	fmt.Scanln(&stu.Score[0])
	fmt.Print("请输入学生第二科成绩: ")
	fmt.Scanln(&stu.Score[1])
	fmt.Print("请输入学生第三科成绩: ")
	fmt.Scanln(&stu.Score[2])

	stu.Ave = (float64(stu.Score[0]) + float64(stu.Score[1]) + float64(stu.Score[2]) )/3

	stuNode := model.StuNode{
		StuInfo: stu,
	}
	return &stuNode
}

func PrintStudentHeader()  {
	fmt.Printf("\t学号    \t姓名    \t第一科成绩\t第二科成绩\t\t第三科成绩\t平均成绩\n")
}

func PrintStudentInfo(stu model.Student) {

	if len(stu.Id) == 0{
		return
	}

	fmt.Printf("\t%s    \t%s    \t%d\t\t%d\t\t\t%d\t\t%.2f\n", stu.Id, stu.Name,
		stu.Score[0],
		stu.Score[1],
		stu.Score[2],
		stu.Ave)
}

func PrintStudent(stu model.Student) {
	PrintStudentHeader()
	PrintStudentInfo(stu)
}


func PrintStudentList(stuList []model.Student) {
	//PrintStudentHeader()
	//for _, stu := range stuList {
	//	PrintStudentInfo(stu)
	//}
}