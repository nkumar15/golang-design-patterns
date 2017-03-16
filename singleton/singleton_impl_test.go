package singleton

import "testing"

func TestGetInstace(t *testing.T) {
	t1 := GetInstance()
	t2 := GetInstance()

	if t1 != t2 {
		t.Error("Instance different")
	}
}
