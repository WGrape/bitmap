package bitmap

import (
	"encoding/json"
	"testing"
)

func TestSS(t *testing.T) {
	m1 := Bitmap{}
	m1.Add(1)
	m1.Add(10)
	m1.Add(999)

	m2 := Bitmap{}
	m2.Add(1)
	m2.Add(10)
	m2.Add(998)

	m1.IntersectWith(&m2)
	b, _ := json.Marshal(m1.Elems())
	if string(b) != "[1,10]" {
		t.Fail()
		return
	}
	t.Logf("%s", b)

	m1.Clear()
	m1.Add(1)
	m1.Add(20)
	m1.Add(999)
	m1.UnionWith(&m2)
	b, _ = json.Marshal(m1.Elems())
	if string(b) != "[1,10,20,998,999]" {
		t.Fail()
		return
	}
	t.Logf("%s", b)
}
