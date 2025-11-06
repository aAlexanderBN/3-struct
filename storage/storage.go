package storage

type DataStorage interface {
	Read() (string, error)
	Write(string)
}
