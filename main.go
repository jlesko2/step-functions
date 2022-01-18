package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/jlesko2/step-functions/data"
	"github.com/jlesko2/step-functions/foo"
	"github.com/jlesko2/step-functions/runner"
	"github.com/jlesko2/step-functions/step"
)

func main() {
	input := &data.Input{ImportantInput: 1}
	output := &data.Output{}
	spew.Dump(input)
	runner := runner.RunnerImpl{
		Funcs: map[string]step.Function{
			"first": &foo.Foo{
				Name:    "firstfunc",
				Fail:    "",
				Success: "second",
			},
			"second": &foo.Foo{
				Name:    "secondfunc",
				Fail:    "",
				Success: "",
			},
		},
	}

	err := runner.Run("first", input, output)
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(output)
}
