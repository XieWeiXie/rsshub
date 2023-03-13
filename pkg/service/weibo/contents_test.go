package weibo

import (
	"fmt"
	"os"
	"testing"
)

func TestContents(t *testing.T) {
	content := NewContents()
	content.ToContent("https://weibo.cn/comment/MwBjBrZcX?uid=1667517061&rl=0#cmtfrm")
}

func TestEnv(t *testing.T) {
	for _, i := range os.Environ() {
		fmt.Println(i)
	}
	fmt.Println(os.Getenv("Cookie"))
}
