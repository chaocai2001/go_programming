package client

import (
	"go_programming/ch13/mathext/series"
	"testing"
)

func TestPackageClient(t *testing.T) {
	series.GetFibonacciSerie(5)
}
