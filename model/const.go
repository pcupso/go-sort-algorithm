package model

// ChoiceType 主菜单/**
type ChoiceType int

const (
	ChoiceAdd     ChoiceType = 1
	ChoiceQuery  ChoiceType = 2
	ChoiceInsert ChoiceType = 3
	ChoiceSort ChoiceType = 4
	ChoiceExit ChoiceType = 0
)

// AddChoiceType 增加子菜单 /**
type AddChoiceType int

const (
	AddChoiceTypeAdd		AddChoiceType = 1
	AddChoiceTypeImport		AddChoiceType = 2
	AddChoiceTypeExit		AddChoiceType = 0
)

// QueryChoiceType 查询子菜单 /**
type QueryChoiceType int

const (
	QueryChoiceTypeAll		QueryChoiceType = 1
	QueryChoiceTypeById		QueryChoiceType = 2
	QueryChoiceTypeByName	QueryChoiceType = 3
	QueryChoiceTypeExit		QueryChoiceType = 0
)

// SortChoiceType 排序子菜单 /**
type SortChoiceType int

const (
	SortChoiceTypeBubble	SortChoiceType = 1
	SortChoiceTypeShell		SortChoiceType = 2
	SortChoiceTypeQuick		SortChoiceType = 3
	SortChoiceTypeHeap		SortChoiceType = 4
	SortChoiceTypeExit		SortChoiceType = 0
)

const ImportBaseFileName string = "template.xlsx"