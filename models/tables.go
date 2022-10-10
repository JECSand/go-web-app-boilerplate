package models

// Table ...
type Table struct {
	Class     string
	Id        string
	TableHead *TableRow
	TableBody *TableBody
	Script    *Script
	Category  string
}

// NewTable ...
func NewTable(class string, id string, thead *TableRow, tbody *TableBody) *Table {
	return &Table{
		Class:     class,
		Id:        id,
		TableHead: thead,
		TableBody: tbody,
		Category:  "default",
	}
}
