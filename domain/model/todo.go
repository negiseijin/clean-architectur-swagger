package model

/*
TODOモデル
*/
type Todo struct {
	Base
	Name string `json:"name,omitempty"` // TODO名
	Done bool   `json:"done,omitempty"` //Done
}
