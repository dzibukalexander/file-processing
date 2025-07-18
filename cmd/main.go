package main

import (
	"fmt"
	"os"

	"github.com/dzibukalexander/file-processing/internal/app"
)

func main() {
	application := app.NewApp()
	if err := application.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("File processed successfully")
}
