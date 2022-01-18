# step-functions

## What is this for?
An easy way to orchestrate a workflow by separating each step into its own package. Each `step.Function{}` implements its own `Execute()` function and indicates which step should be executed next depending on the result (success, failure).

## How to use this library?
1. Define your `Input` and `Output` types
2. Create your own package, `foo`, that implements the `step.Function{}` interface
3. Initialize a struct of type `foo` and add it to the map of functions in `runner.RunnerImpl{}` in `main.go`
4. Call `runner.Run()` providing it the first step's key


## How to run this package?
From the root of the repository, `go run main.go`

```
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
```
