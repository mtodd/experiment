# Experiment

Experiment is a Go package to test and evaluate new code paths without
interfering with the users end result.

This is inspired by the [GitHub Scientist gem](https://github.com/github/scientist).

## Usage

```go
func main() {
    exp, err := experiment.New(
        experiment.Name("my-test"),
        experiment.Enabled(shouldRunTest()),
        experiment.Percentage(10),
        experiment.Compare(comparisonMethod),
    )
    if err != nil {
        fmt.Println(err)
        return
    }

	exp.Control(func(ctx context.Context) (interface{}, error) {
		return "my-text", nil
	})
	exp.Test("buffer", func(ctx context.Context) (interface{}, error) {
		buf := bytes.NewBufferString("")
		buf.WriteString("new")
		buf.Write([]byte(`-`))
		buf.WriteString("text")

		return string(buf.Bytes()), nil
	})

	res, err := exp.Run()
    if err != nil {
        fmt.Println(err)
        return
    }
    str = res.(string)
    fmt.Println(str)
}

func shouldRunTest() bool {
	return os.Getenv("ENV") == "prod"
}

func comparisonMethod(control interface{}, test interface{}) bool {
	c := control.(string)
	t := test.(string)

	return c == t
}
```
