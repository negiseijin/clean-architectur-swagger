package repository

type JsonRepository interface {
	Marshal(i interface{}, v interface{}) error
	Unmarshal(item interface{}) error
}
