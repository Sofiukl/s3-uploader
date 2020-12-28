package utils

import (
	"fmt"
	"os"

	guuid "github.com/google/uuid"
)

//PrintErrorf - prints the error
func PrintErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

// GenUUID - generate uuid4
func GenUUID() string {
	id := guuid.New().String()
	return id
}
