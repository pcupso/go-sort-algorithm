package util

import (
	"fmt"
	"go-sort-algorithm/model"
	"time"
)

func BubbleSort(stuList []model.Student) []model.Student {
	sortedList := stuList

	startTime := time.Now().UnixNano() / 1e6
	fmt.Printf("start time : %v\n", time.Now())

	for i := len(sortedList)-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if sortedList[j].Ave > sortedList[i].Ave {
				//exchange
				var tmp model.Student
				tmp = sortedList[j]
				sortedList[j] = sortedList[i]
				sortedList[i] = tmp
			}
		}
	}

	endTime := time.Now().UnixNano() / 1e6

	PrintStudentList(sortedList)
	totalMills := endTime - startTime
	fmt.Printf("end time : %v, 共耗时= %vs%vms\n", time.Now(), totalMills/1000, totalMills%1000 )

	return sortedList
}

func ShellSort(stuList []model.Student) []model.Student{
	sortedList := stuList

	startTime := time.Now().UnixNano() / 1e6
	fmt.Printf("start time : %v\n", time.Now())

	// 确定h大小
	h := 1
	for h < len(sortedList)/2 {
		h = 2*h + 1
	}

	for h>=1 {
		for i:=h; i < len(sortedList); i++ {
			for j := i; j>=h; j-=h {
				if sortedList[j-h].Ave > sortedList[j].Ave {
					//exchange
					var tmp model.Student
					tmp = sortedList[j]
					sortedList[j] = sortedList[i]
					sortedList[i] = tmp
				}
			}
		}
		h = h/2
	}

	endTime := time.Now().UnixNano() / 1e6
	PrintStudentList(sortedList)
	totalMills := endTime - startTime
	fmt.Printf("end time : %v, 共耗时= %vs%vms\n", time.Now(), totalMills/1000, totalMills%1000 )

	return sortedList
}

func QuickSort(stuList []model.Student) []model.Student{
	sortedList := stuList

	startTime := time.Now().UnixNano() / 1e6
	fmt.Printf("start time : %v\n", time.Now())

	minIndex := 0
	maxIndex := len(stuList)-1
	doQuickSort(sortedList, minIndex, maxIndex)

	endTime := time.Now().UnixNano() / 1e6
	PrintStudentList(sortedList)
	totalMills := endTime - startTime
	fmt.Printf("end time : %v, 共耗时= %vs%vms\n", time.Now(), totalMills/1000, totalMills%1000 )

	return sortedList
}

func doQuickSort(stuList []model.Student, minIndex int, maxIndex int)  {
	if maxIndex <= minIndex {
		return
	}

	seperatedValues := partion(stuList, minIndex, maxIndex)
	doQuickSort(stuList, minIndex, seperatedValues-1)
	doQuickSort(stuList, seperatedValues+1, maxIndex)
}

func partion(stuList []model.Student, minIndex int, maxIndex int) int {
	// 确定分组界限值
	compareItem := stuList[minIndex]
	left := minIndex
	right := maxIndex+1

	//切分
	for {
		for ;compareItem.Ave < stuList[right-1].Ave; {
			if right==minIndex {
				break
			}
			right-=1
		}
		right-=1

		for ;stuList[left].Ave < compareItem.Ave; {
			if left==maxIndex {
				break
			}
			left+=1
		}
		left+=1

		if left >= right {
			break
		} else {
			// 交换值
			tmp := stuList[left]
			stuList[left] = stuList[right]
			stuList[right] = tmp
		}
	}

	tmp := stuList[minIndex]
	stuList[minIndex] = stuList[right]
	stuList[right] = tmp

	return right
}


func HeapSort(stuList []model.Student) []model.Student {
	sortedList := stuList

	startTime := time.Now().UnixNano() / 1e6
	fmt.Printf("start time : %v\n", time.Now())

	buildHeap(sortedList)
	for i := len(sortedList); i > 1; i-- {
		sortedList[0], sortedList[i-1] = sortedList[i-1], sortedList[0]
		adjustHeap(sortedList[:i-1], 0)
		//fmt.Println("the heap is ", values)
	}

	endTime := time.Now().UnixNano() / 1e6
	PrintStudentList(sortedList)
	totalMills := endTime - startTime
	fmt.Printf("end time : %v, 共耗时= %vs%vms\n", time.Now(), totalMills/1000, totalMills%1000 )

	return sortedList
}

func buildHeap(values []model.Student) {
	for i := len(values); i>=0; i-- {
		adjustHeap(values, i)
	}
}

func adjustHeap(values []model.Student, pos int) {
	node := pos
	length := len(values)
	for node < length {
		var child int = 0
		if 2*node+2 < length {
			if values[2*node+1].Ave > values[2*node+2].Ave {
				child = 2*node + 1
			} else {
				child = 2*node + 2
			} ////////选出大子节点
		} else if 2*node+1 < length {
			child = 2*node + 1
		}
		if child > 0 && values[child].Ave > values[node].Ave {
			values[node], values[child] = values[child], values[node]
			node = child
		} else {
			break
		}
	}
}