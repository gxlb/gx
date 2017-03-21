package rdb

type Field interface {
	Name() string
}

type Table interface {
	Name() string
	Create() error
	Insert() error
	Update() error
	Delete() error
	Replace() error
	Truncate() error
	Drop() error
	Commit() error
}

type Database interface {
	Create() error
	Remove() error
	Use() error
	Name() string
}

type DatabaseServer interface {
	Name() string
	Ip() string
	Port() int
	Connect() error
}
