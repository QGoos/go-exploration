package main

import (
	"os"
	"time"

	"exploration/maths"
)

func main() {
	t := time.Now()
	maths.SVGWriter(os.Stdout, t)
}
