# Simple Mock Generator

This shows a really simple implementation of a mock generator. It uses AST and _go/packages_ to read in an interface and output a mock object based on it.

If you open _airplane.go_ you will see the go:generate statement there, run it and you'll get a mock object generated.

Have a look through the code in _main.go_ to see how it works.
