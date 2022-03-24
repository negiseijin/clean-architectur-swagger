package repository

type HelpersRepository interface {
	JsonToStruct(i interface{}, v interface{}) error
}
