package types

type App struct{
	id int64
	Name string `validate:"required"`
	content string `validate:"required"`
}