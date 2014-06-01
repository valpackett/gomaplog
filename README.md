# gomaplog

A simple logging library for the Go programming language.

Because [spacelog](https://github.com/spacemonkeygo/spacelog) was not enough and too much at the same time.
gomaplog borrows some tricks from spacelog though (basically, the coloring in text/template.)

## Features

- log functions accept a `map[string]interface{}` of arbitrary parameters (hence `map` in the name)
- `io.Writer` is the output abstraction
- same level names and numbers as in syslog (and GELF)
- level filtering
- no magic, you create a `Logger` in your own namespace
- simple `Formatter` interface
- JSON ([GELF 1.1](http://graylog2.org/gelf#specs)) Formatter included
- `text/template` Formatter included, with a shiny colorful template

## Usage

```go
import "github.com/myfreeweb/gomaplog"

var logger = gomaplog.StdoutLogger(gomaplog.DefaultTemplateFormatter)
// var logger = gomaplog.StdoutLogger(gomaplog.DefaultJSONFormatter)
// var logger = &gomaplog.Logger{Formatter: MyOwnFormatter, Writer: os.Stderr, Host: "app", MaxLevel: gomaplog.Debug}
// etc.
// you can, like, replace the template logger with a json one based on command line flags

func main() {
  logger.Host = "myapp@localhost" // use os.Hostname or something
  logger.MaxLevel = gomaplog.Info // the check is <=, so, only Debug won't be printed

  logger.Info("Build started", gomaplog.Extras{"version": build.Version})
  logger.WarningL("Build failed", "Long long\ndescription\nof why it failed", gomaplog.Extras{"version": build.Version})
  logger.Error("The dungeon collapsed", gomaplog.Extras{"seed": 012345})
}
```
