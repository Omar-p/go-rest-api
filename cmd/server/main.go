package main

import "fmt"

type Server struct {
}

func Run() error {
	fmt.Println("RUN")
	return nil
}

func main() {
	fmt.Println("GO REST API")
	if err := Run(); err != nil {

	}
}
