package util

import "testing"

/*
 测试单个模块
*/
func TestGenShortId(t *testing.T)  {
	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		t.Error("genshortId failed!")
	}
	t.Log("genshortId test pass")
}