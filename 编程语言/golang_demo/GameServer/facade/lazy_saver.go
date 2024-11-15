package facade

type LazySaver interface {
	Save()
	Key() string
}
