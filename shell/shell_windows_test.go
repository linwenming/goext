package utils

import (
	"context"
	"fmt"
	"testing"
)

func TestExecShell(t *testing.T) {

	// windows
	r,_:= ExecShell(context.Background(),"dir")
	fmt.Println(r)
}
