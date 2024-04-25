package errors

import (
	"fmt"
	"testing"
)

var reviewMessageMap = map[Code]string{
	10000: "测试1",
	10001: "测试2",
}

var businessMessageMap = map[Code]string{
	10001: "哈哈哈哈",
	20000: "2测试1",
	20001: "2测试2",
}

func TestPrint(t *testing.T) {
	Init(reviewMessageMap, businessMessageMap)
	var code Code = 10001
	fmt.Println("error_print:", code.ToCodeError()) // 因为 Code 实现了 error 接口，所以打印的值是 code.Error()
}
