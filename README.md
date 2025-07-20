# Vectors

[![Go Reference](https://pkg.go.dev/badge/github.com/Dobefu/vectors.svg)](https://pkg.go.dev/github.com/Dobefu/vectors)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_vectors&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Dobefu_vectors)
[![Go Report Card](https://goreportcard.com/badge/github.com/Dobefu/vectors)](https://goreportcard.com/report/github.com/Dobefu/vectors)

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
