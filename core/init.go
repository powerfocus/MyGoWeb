package core

var (
	Serv     Server
	HHandler Filter
)

func init() {
	Serv = NewServer()
	HHandler = NewFilter()
}
