package errors

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	var code Code = 20000
	fmt.Println("error_print:", code) // 因为 Code 实现了 error 接口，所以打印的值是 code.Error()
}
