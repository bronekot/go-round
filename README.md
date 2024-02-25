# Round Package

The `round` package provides a simple and efficient way to round floating-point numbers to a specified number of decimal places in Go. It leverages the Go standard library's `math` package to perform its calculations.

## Installation

To use the `round` package in your Go project, first, ensure you have Go installed on your system. Then, you can install the package by adding it to your project's dependencies:

```go
go get -u github.com/bronekot/go-round
```

## Usage

Here's how you can use the `round` package to round a floating-point number:

```go
package main

import (
    "fmt"
    "github.com/yourusername/round" // Replace with the actual import path
)

func main() {
    roundedValue := round.Round(3.14159, 2)
    fmt.Println(roundedValue) // Output: 3.14
}
```

### Function: Round

The `Round` function is the core of this package. Here's its signature:

```go
func Round(x float64, prec int) float64
```

Parameters:
- `x`: The floating-point number you want to round.
- `prec`: The number of decimal places to round `x` to.

Returns:
- The rounded value of `x` to `prec` decimal places.

## How It Works

The `Round` function rounds a floating-point number `x` to `prec` decimal places. It does this by:
1. Multiplying `x` by 10 raised to the power of `prec`.
2. Separating the result into its integer and fractional parts.
3. If the fractional part is 0.5 or greater, or if it's between -0.5 and 0 (inclusive) for negative numbers, it rounds up by using the `Ceil` function. Otherwise, it rounds down using the `Floor` function.
4. Dividing the result by 10 raised to the power of `prec` to adjust the value back to the original scale.

## Contributing

Contributions to improve the `round` package are welcome. Please feel free to submit issues and pull requests through the GitHub repository.

## License

This package is released under the MIT License. See the LICENSE file for details.
