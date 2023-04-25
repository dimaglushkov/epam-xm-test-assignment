package ports

type Handler interface {
	Run() error
}
