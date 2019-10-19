package uuid

import (
	"fmt"
	"testing"
)

func TestUUID_String(t *testing.T) {
	fmt.Println(New().String())

	fmt.Println(NewUUID())

}
