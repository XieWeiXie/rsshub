package hub

import "testing"

func TestPage(t *testing.T) {
	nr := NewToResponse()
	nr.Response("https://weibo.cn/comment/MwFyL2GkK?ckAll=1")

}
