// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePerson = "people"

// Person mapped from table <people>
type Person struct {
	UUID    string `gorm:"column:UUID" json:"UUID"`
	Name    string `gorm:"column:name" json:"name"`
	Age     int32  `gorm:"column:age" json:"age"`
	Version int32  `gorm:"column:version" json:"version"`
}

// TableName Person's table name
func (*Person) TableName() string {
	return TableNamePerson
}
