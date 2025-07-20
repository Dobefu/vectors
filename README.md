# Vectors

[![Go Reference](https://pkg.go.dev/badge/github.com/Dobefu/vectors.svg)](https://pkg.go.dev/github.com/Dobefu/vectors)

A Go package providing 2D and 3D vector types with mathematical operations.

## Installation

```bash
go get github.com/Dobefu/vectors
```

## Usage

```go
package main

import "github.com/Dobefu/vectors"

func main() {
    vec := vectors.Vector2{X: 3, Y: 4}
    magnitude := vec.Magnitude() // Returns 5.0
}
```
