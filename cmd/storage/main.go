package main

import (
	"fmt"

	"github.com/KosmanG/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println(st)
}
