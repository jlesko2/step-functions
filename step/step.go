package step

type Config struct {
	Name string
}

type Function interface {
	GetConfig() *Config
	OnSuccess() string
	OnFail() string
	Validate(input interface{}, output interface{}) error
	Execute(input interface{}, output interface{}) error
}
