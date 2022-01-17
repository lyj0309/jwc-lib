package lib

import (
	"fmt"
	"os"
	"testing"
)

func TestLib(t *testing.T) {
	fmt.Println("env", os.Getenv(ConfEnvName))
}
