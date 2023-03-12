package weibo

import "testing"

func TestContents(t *testing.T) {
	content := NewContents()
	content.ToContent("https://weibo.cn/comment/MwBjBrZcX?uid=1667517061&rl=0#cmtfrm")
}
