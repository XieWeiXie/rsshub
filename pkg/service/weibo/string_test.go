package weibo

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	toUrlValue("/mblog/operation/MwV0ZfawW?uid=7431810763&rl=1&gid=10001")
}

func TestUrl(t *testing.T) {
	fmt.Println(toUrlUid("https://weibo.cn/comment/M5qx4gHyI?uid=1667517061&rl=0#cmtfrm"))
}
