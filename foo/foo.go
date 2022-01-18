package foo

import (
	"fmt"

	"github.com/jlesko2/step-functions/step"
)

type Foo struct{}

func (f *Foo) GetConfig() *step.Config {
	return &step.Config{
		Name: "Foo",
	}
}

func (f *Foo) ProceedOnErr() bool {
	return false
}

func (f *Foo) Execute(input interface{}, output interface{}) error {
	fmt.Println("foo.Execute()...")
	input := input.(main.Input).Field2 +
	return fmt.Errorf("hello")
}

func (f *Foo) GetNext() string {
	return ""
}
