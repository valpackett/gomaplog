# gomaplog

A simple logging library for the Go programming language.

## Features

- log functions accept a `map[string]interface{}` of arbitrary parameters (hence `map` in the name)
- `io.Writer` is the output abstraction
- simple `Formatter` interface
- JSON ([GELF 1.1](http://graylog2.org/gelf#specs)) Formatter included
- colorful plain text Formatter included
