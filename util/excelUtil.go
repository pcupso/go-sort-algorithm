package util

import (
	"fmt"
	"go-sort-algorithm/model"
	"strconv"
)
import "github.com/360EntSecGroup-Skylar/excelize"


func ReadXlsTemplate(fileName string) []model.Student {
	retList := make([]model.Student, 0)

	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var rows [][]string  = xlsx.GetRows("Sheet1")
	for index, row := range rows {
		if index == 0 {
			continue
		}

		for i, item := range row {
			fmt.Printf("## cell %d : col-%d, value = %s\n", index, i, item)

		}

		student := model.Student{}
		student.Id = row[0]
		student.Name = row[1]
		student.Score[0], _ = strconv.Atoi(row[2])
		student.Score[1], _ = strconv.Atoi(row[3])
		student.Score[2], _ = strconv.Atoi(row[4])
		student.Ave, _ = strconv.ParseFloat(row[5], 64)
		retList = append(retList, student)
	}

	return retList
}