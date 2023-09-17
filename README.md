# gooseloggers

[![Default](https://github.com/mfinelli/go-gooseloggers/actions/workflows/default.yml/badge.svg)](https://github.com/mfinelli/go-gooseloggers/actions/workflows/default.yml)
[![Go Reference](https://pkg.go.dev/badge/go.finelli.dev/gooseloggers.svg)](https://pkg.go.dev/go.finelli.dev/gooseloggers)

A few custom [goose](https://github.com/pressly/goose) loggers that I use in
some of my other projects.

## usage

```go
package main

import (
	"github.com/rs/zerolog/log"
	glog "go.finelli.dev/gooseloggers/zerolog"
)

func main() {
	goose.SetLogger(glog.GooseZerologLogger(&log.Logger))
	// now do some goose-stuff
}
```

### loggers

I use a [zerolog](https://github.com/rs/zerolog)-based logger for my
applications and a [pterm](https://github.com/pterm/pterm)-based logger for
my CLI applications, but it's not too challenging to add others.

## license

```
Copyright 2023 Mario Finelli

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
