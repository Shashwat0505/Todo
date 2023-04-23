package models

import(
	"gorm.io/gorm"
)


type Activity struct{
	gorm.Model
	ActivityName string

}
