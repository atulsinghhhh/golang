package storage

type Storage interface {
	CreateApp(name string, content string)(int64,error)
}