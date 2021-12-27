package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err == nil {
		fmt.Println("create file sucess!")
	} else {
		fmt.Printf("create file failure, err = %s\n", err.Error())
		return nil
	}
	return file
}

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err == nil {
		fmt.Println("open file sucess!")
	} else {
		fmt.Printf("open file failure, err = %s\n", err.Error())
		return nil
	}
	return file
}

var fileName = "score"

func getFile() * os.File {
	_, err := os.Lstat(fileName)
	if os.IsExist(err) {
		file, _ := os.Open(fileName)
		return file
	} else {
		file, _ := os.Create(fileName)
		return file
	}
}

func SaveData(data string)  {

	file := getFile()
	if file == nil {
		fmt.Println("文件指针不能为空！")
		return
	}
	_, err := file.Write([]byte(data))
	if err == nil {
		fmt.Println("文件写入内容成功！")
	} else {
		fmt.Println("文件写入内容失败！")
	}
}

func ReadData() string {
	bytes, err := ioutil.ReadFile(fileName)
	if err == nil {
		fmt.Println("文件的内容:",string(bytes))
		return string(bytes)
	} else {
		fmt.Println(err)
		return ""
	}
}