# Simple Mock Generator

This shows a really simple implementation of a mock generator. It uses AST and _go/packages_ to read in an interface and output a mock object based on it.

## Install the generator

Ensure you are in the 04-simple-mock-gen dir on your command line.

If you are running on GitPod, then you can skip this part as it will be setup when you open the repo, if you are running this fresh on your machine, run the below command from the repo root to install the generator to your path:

```go
    cd 04-simple-mock-gen && go install && cd ../
```

## Running the generator

If you open [airplane.go](airplane.go) you will see the go:generate statement there, run it and you'll get a mock object generated.

Have a look through the code in [main.go](main.go) to see how it works.

When you make changes either run with `go run main.go -type=AirplaneReader` or re-run the go install command and execute the inline go:generate command.
