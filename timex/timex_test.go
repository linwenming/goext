package timex

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	v := Parse("2019-05-11 23:00:59")
	fmt.Println(v.Unix())

	s := String(v)
	fmt.Println(s)
}

func TestRandSecond(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandSecond(10))
	}
}
