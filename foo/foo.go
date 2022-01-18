package foo

import (
	"fmt"

	"github.com/jlesko2/step-functions/data"
	"github.com/jlesko2/step-functions/step"
)

type Foo struct {
	Name    string
	Success string
	Fail    string
}

func (f *Foo) GetConfig() *step.Config {
	return &step.Config{
		Name: "foo",
	}
}

func (f *Foo) Execute(input interface{}, output interface{}) error {
	fmt.Printf("executing foo.Execute(%s)\n", f.Name)
	out := output.(*data.Output)
	in := input.(*data.Input)
	out.ImportantOutput = in.ImportantInput + 1
	return nil
}

func (f *Foo) OnSuccess() string {
	return f.Success
}

func (f *Foo) OnFail() string {
	return f.Fail
}

func (f *Foo) Validate(input interface{}, output interface{}) error {
	_, ok := input.(*data.Input)
	if !ok {
		return fmt.Errorf("error: Input is not correct type")
	}
	_, ok = output.(*data.Output)
	if !ok {
		return fmt.Errorf("error: Output is not correct type")
	}

	return nil
}
