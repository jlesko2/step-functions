package step

type Config struct{}

type Function interface {
	Execute(input interface{}, output interface{}) error
	ProceedOnError() bool
}
