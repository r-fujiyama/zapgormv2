# zapgormv2

[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/r-fujiyama/zapgormv2/blob/master/LICENSE)
[![CI](https://github.com/r-fujiyama/zapgormv2/workflows/CI/badge.svg)](https://github.com/r-fujiyama/zapgormv2/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/r-fujiyama/zapgormv2)](https://goreportcard.com/report/github.com/r-fujiyama/zapgormv2)

## Guide

## Installation

```sh
$ go get -u github.com/r-fujiyama/zapgormv2
```

### Example

```go
import (
	"github.com/r-fujiyama/zapgormv2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Example() {
	zapLogger, _ := zap.NewProduction()
	logger := zapgormv2.New(zapLogger)
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	db, _ := gorm.Open(nil, &gorm.Config{Logger: logger})

	// do stuff normally
	var _ = db // avoid "unused variable" warn
}
```

## License

[MIT](https://github.com/r-fujiyama/zapgormv2/blob/master/LICENSE)
