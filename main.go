package main

import (
	"github.com/jlesko2/step-functions/foo"
	"github.com/jlesko2/step-functions/runner"
	"github.com/jlesko2/step-functions/step"
)

func main() {
	localFoo := &foo.Foo{}
	runner := runner.RunnerImpl{
		Funcs: map[string]step.Function{
			"First": localFoo,
		},
	}
	type Input struct {
		Field1 string
		Field2 int64
	}
	type Output struct {
	}
	runner.Run("First", &Input{}, &Output{})

}
