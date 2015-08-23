package genbase

import (
	"testing"
)

func TestPathJoinAll(t *testing.T) {
	ps := pathJoinAll("misc/fixture", "a", "b")

	if len(ps) != 2 {
		t.Fatalf("unexpected", len(ps))
	}

	if ps[0] != "misc/fixture/a" || ps[1] != "misc/fixture/b" {
		t.Fatalf("unexpected", ps)
	}
}

func TestGetKeys(t *testing.T) {

	result := GetKeys("a:\"foo\" b:\"bar\"")
	if len(result) != 2 {
		t.Log("keys length is not 2, actual", len(result))
		t.Fail()
	}
	if result[0] != "a" {
		t.Log("result[0] is not \"a\", actual", result[0])
		t.Fail()
	}
	if result[1] != "b" {
		t.Log("result[1] is not \"b\", actual", result[1])
		t.Fail()
	}
}
