package runner

import (
	"fmt"

	"github.com/jlesko2/step-functions/step"
)

type Runner interface {
	Run(step.Function, step.Input, step.Output) error
}

type RunnerImpl struct {
	Funcs map[string]step.Function
}

func (r *RunnerImpl) Run(funcName string, i step.Input, o step.Output) error {
	// extract step function to run
	f := r.Funcs[funcName]
	// execute step
	err := f.Execute(i, o)
	// check if execution should halt on error
	if err != nil && !f.ProceedOnErr() {
		return fmt.Errorf("error 1")
	}
	// exit condition
	if f.GetNext() != "" {
		r.Run(f.GetNext(), i, o)
	}
	return nil
}
