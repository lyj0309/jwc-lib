package esLib

import (
	"fmt"
	"testing"
)

func TestEs(t *testing.T) {
	es := NewElastic()
	fmt.Println(es)
}
