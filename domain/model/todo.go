package model

import "gorm.io/gorm"

/*
TODOモデル
*/
type Todo struct {
	gorm.Model
	Name string `json:"name,omitempty"` // TODO名
	Done bool   `json:"done,omitempty"` //Done
}
