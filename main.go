package main

import (
	"encoding/json"
	"fmt"
	"go-sort-algorithm/model"
	"go-sort-algorithm/util"
	"os"
	"strings"
)

var stuList *model.StuNode
var stuLastNode *model.StuNode

/**
	初始化信息
 */
func initBaseInfo()  {

	stuList = nil
	stuLastNode = nil

	initData := util.ReadData()
	if len(initData) > 0 {
		json.Unmarshal([]byte(initData), &stuList)
	}

	if stuList != nil {
		var p *model.StuNode = stuList.Next
		if p == nil {
			stuLastNode = p
		} else {
			for p.Next != nil {
				var tmp *model.StuNode = p.Next
				p = tmp
				tmp = nil
			}
			stuLastNode = p
		}
	}
}

func importXlsxData(fileName string)  {
	studentList := util.ReadXlsTemplate(fileName)
	for index, student := range studentList {
		fmt.Printf("### index = %d, data = %v \n", index, student)

		node := model.NewStuNode()
		node.StuInfo = student
		node.Next = nil
		doInsert(node)
	}
}

func main() {

	initBaseInfo()

	for {
		choiceType := util.PrintMenu()
		switch choiceType {
		case model.ChoiceAdd:
			add()
		case model.ChoiceQuery:
			query()
		case model.ChoiceInsert:
			insert()
		case model.ChoiceSort:
			sort()
		case model.ChoiceExit:
			os.Exit(1)
		default:
			println("Not invalid, please input valid choice\n")
		}
	}
}

func add() {
	for {
		choice := util.PrintAddMenu()
		switch choice {
		case model.AddChoiceTypeAdd:
			inNode := util.InputStudent()
			doInsert(inNode)
		case model.AddChoiceTypeImport:
			fmt.Print("Please input your excel template fileName : ")
			var fileName string
			fmt.Scanln(&fileName)

			// 开始导入
			importXlsxData(fileName)

		case model.AddChoiceTypeExit:
			return
		default:
			println("Not invalid, please input valid choice\n")
		}
	}
}

func query() {
	for {
		choice := util.PrintQueryMenu()
		switch choice {
		case model.QueryChoiceTypeAll:
			doQueryAll()
		case model.QueryChoiceTypeById:
			fmt.Print("Please input query student id : ")
			var id string
			fmt.Scanln(&id)

			doQueryById(id)

		case model.QueryChoiceTypeByName:
			fmt.Print("Please input query student name : ")
			var name string
			fmt.Scanln(&name)

			doQueryByName(name)
		case model.QueryChoiceTypeExit:
			return
		default:
			println("Not invalid, please input valid choice\n")
		}
	}
}

func insert() {
	inNode := util.InputStudent()
	doInsert(inNode)
}

func sort() {

	fmt.Printf("Please input data file name:")
	var fileName string
	fmt.Scanln(&fileName)

	fmt.Printf("Import data starting....\n")
	studentList := util.ReadXlsTemplate(fileName)
	fmt.Printf("Import data finished....\n")

	for {
		choice := util.PrintSortMenu()
		switch choice {
		case model.SortChoiceTypeBubble:
			util.BubbleSort(studentList)
		case model.SortChoiceTypeShell:
			util.ShellSort(studentList)
		case model.SortChoiceTypeQuick:
			util.QuickSort(studentList)
		case model.SortChoiceTypeHeap:
			util.HeapSort(studentList)
		case model.SortChoiceTypeExit:
			return
		default:
			println("Not invalid, please input valid choice\n")
		}
	}
}

func isExist(id string) bool {
	stu := doQueryById(string(id))
	return len(stu.Id) > 0
}

func doInsert(inNode *model.StuNode) {

	// 判断是否存在
	if isExist(inNode.StuInfo.Id) {
		fmt.Printf("学号：%s, 学生已存在，请检查后输入信息正确性\n", inNode.StuInfo.Id)
		return
	}

	// 第一次添加
	if stuList == nil {
		stuList = inNode
		stuLastNode = inNode
	} else {
		// 非第一次添加
		var curNode *model.StuNode = stuList

		// 首个节点比输入的小，直接插入
		if curNode.StuInfo.Ave < inNode.StuInfo.Ave {
			inNode.Next = curNode
			stuList = inNode
		} else {

			if curNode.Next == nil {
				curNode.Next = inNode
				stuLastNode = inNode
			} else {

				var preNode *model.StuNode = stuList
				for curNode.Next != nil {
					// next
					var tmp *model.StuNode = curNode.Next
					curNode = tmp
					tmp = nil

					if curNode.StuInfo.Ave < inNode.StuInfo.Ave {
						preNode.Next = inNode
						inNode.Next = curNode
						break
					}
					// 先节点下滑一个
					preNode = curNode
				}

				if curNode.Next == nil && curNode.StuInfo.Ave >= inNode.StuInfo.Ave {
					curNode.Next = inNode
					stuLastNode = inNode
				}

				// 释放指针，防止野指针
				preNode = nil
			}
		}
		curNode = nil
	}
	inNode = nil

	// 保存数据
	jsonInfo, err := json.Marshal(stuList)
	if err != nil {
		fmt.Println("error:", err)
	}
	util.SaveData(string(jsonInfo))
	os.Stdout.Write(jsonInfo)

	fmt.Println("添加成功")
}

func doQueryAll() {
	var curNode *model.StuNode = stuList

	util.PrintStudentHeader()
	for curNode != nil {
		util.PrintStudentInfo(curNode.StuInfo)

		// next
		var tmp = curNode.Next
		curNode = tmp
		tmp = nil
	}
}

func doQueryById(id string) model.Student {
	var curNode *model.StuNode = stuList
	for curNode != nil {

		if strings.Compare(curNode.StuInfo.Id, id) == 0 {
			util.PrintStudent(curNode.StuInfo)
			return curNode.StuInfo
		}

		// next
		var tmp = curNode.Next
		curNode = tmp
		tmp = nil
	}
	return model.Student{}
}

func doQueryByName(name string) {
	var curNode *model.StuNode = stuList

	util.PrintStudentHeader()
	for curNode != nil {

		if strings.Compare(curNode.StuInfo.Name, name) == 0 {
			util.PrintStudentInfo(curNode.StuInfo)
		}

		// next
		var tmp = curNode.Next
		curNode = tmp
		tmp = nil
	}
}
