# stopwatch

A simply package for timing your code. The intention is to provide a simple,
light-weight library for benchmarking specific bits of your code when need be.

## Example

```go
package main

import (
  "fmt"

  "github.com/bradhe/stopwatch"
)

func main() {
  watch := stopwatch.Start()

  // Do some work.

  watch.Stop()
  fmt.Printf("Milliseconds elapsed: %v\n", watch.Milliseconds())
}
```

## Contributing

1. Fork and fix/implement in a branch.
1. Make sure tests pass.
1. Make sure you've added new coverage.
1. Submit a PR.
