package main

import (
	"context"
	"fmt"
	"github.com/Omar-p/go-rest-api/internal/db"
	"os"
)

type Server struct {
}

func Run() error {
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to db")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API")

	if err := Run(); err != nil {
		os.Exit(1)
	}
}
