package weibo

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	first := NewFirst("1667517061", 1)
	first.ToFirst(func(strings []string) {
		for _, one := range strings {
			fmt.Println(one)
		}
	})
}
