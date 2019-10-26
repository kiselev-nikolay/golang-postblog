package main

import (
	"fmt"
	sass "github.com/wellington/go-libsass"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hello World")
	stop := time.Now()
	fmt.Println(stop.Sub(start))
	fmt.Println(sass)
}
