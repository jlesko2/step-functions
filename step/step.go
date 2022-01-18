package step

type Config struct {
	Name string
}

type Function interface {
	GetConfig() *Config
	ProceedOnErr() bool
	GetNext() string
	Execute(input interface{}, output interface{}) error
}
