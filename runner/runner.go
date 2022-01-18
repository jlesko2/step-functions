package runner

import (
	"github.com/jlesko2/step-functions/step"
)

type Runner interface {
	Run(f step.Function, input interface{}, output interface{}) error
}

type RunnerImpl struct {
	Funcs map[string]step.Function
}

func (r *RunnerImpl) Run(funcName string, i interface{}, o interface{}) error {
	// extract step function to run
	f := r.Funcs[funcName]
	// validate input
	err := f.Validate(i, o)
	if err != nil {
		return err
	}
	// execute step
	err = f.Execute(i, o)
	// check if execution should halt on error
	if err != nil {
		if f.OnFail() != "" {
			// check which function to execute next on failure
			r.Run(f.OnSuccess(), i, o)
		} else {
			// if no funciton next, return
			return err
		}
	}
	// exit condition
	if f.OnSuccess() != "" {
		r.Run(f.OnSuccess(), i, o)
	}
	return nil

}
