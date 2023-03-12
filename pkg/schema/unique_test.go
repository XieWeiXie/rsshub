package schema

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T) {
	var request = new(RequestUniqueOp)
	req1 := ToRequestUniqueOp(
		request.WithRequestUniqueOpType("weibo"),
		request.WithRequestUniqueOpUID("123"),
		request.WithRequestUniqueOpURL([]string{"12", "34", "56", "78"}),
	)
	fmt.Println(fmt.Sprintf("%#v", req1))
	req1.WithRequestUniqueOpURL([]string{"12", "56", "78"})(&req1)
	fmt.Println(fmt.Sprintf("%#v", req1))
}

func TestFields(t *testing.T) {
	f := Unique{}
	fmt.Println(f.Fields())

}
