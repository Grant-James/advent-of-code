package main

import "fmt"

func main() {
	err := run()
	if err != nil {
	    fmt.Printf("err: %v", err)
	    return
	}
}

func run() error {
	return nil
}