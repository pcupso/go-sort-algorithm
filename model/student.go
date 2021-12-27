package model

type Student struct {
	Id string
	Name string
	Score [3]int
	Ave float64
}

type StuNode struct {
	StuInfo Student
	Next	*StuNode
}

func NewStuNode() *StuNode {
	return &StuNode{

	}
}
