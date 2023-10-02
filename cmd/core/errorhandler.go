package core

import (
	"fmt"
	"os"
)

func HandlerError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
