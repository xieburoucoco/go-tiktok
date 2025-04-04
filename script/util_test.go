package script

import (
	"context"
	"testing"
)

func TestGetXbValueByNodeCmd(t *testing.T) {
	valueByNodeCmd, err := New().GetXbValueByNodeCmd(context.Background(), "https://www.example.com/test?param1=value1&param2=value2", "nil")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(valueByNodeCmd)
}

//func TestViper(t *testing.T) {
//	url := "https://default.com"
//	userAgent := "DefaultAgent"
//	script, err := RunScript(url, userAgent)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	t.Log(script)
//}
