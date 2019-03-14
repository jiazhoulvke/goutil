package goutil

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	s := RandomString(8)
	if len(s) != 8 {
		t.Fail()
	}
}

func TestInStringSlice(t *testing.T) {
	s := []string{"1", "2", "3"}
	if !InStringSlice("1", s) {
		t.Fail()
	}
	if InStringSlice("4", s) {
		t.Fail()
	}
}

func TestInIntSlice(t *testing.T) {
	s := []int{1, 2, 3}
	if !InIntSlice(1, s) {
		t.Fail()
	}
	if InIntSlice(4, s) {
		t.Fail()
	}
}

func TestInInt64Slice(t *testing.T) {
	s := []int64{1, 2, 3}
	if !InInt64Slice(1, s) {
		t.Fail()
	}
	if InInt64Slice(4, s) {
		t.Fail()
	}
}

func TestJoinIntSlice(t *testing.T) {
	s := []int{1, 2, 3}
	if JoinIntSlice(s, ",") != "1,2,3" {
		t.Fail()
	}
}

func TestJoinInt64Slice(t *testing.T) {
	s := []int64{1, 2, 3}
	if JoinInt64Slice(s, ",") != "1,2,3" {
		t.Fail()
	}
}

func TestString2IntSlice(t *testing.T) {
	s, err := String2IntSlice("1,2,3", ",")
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 3 {
		t.Fail()
	}
	if s[0] != 1 || s[1] != 2 || s[2] != 3 {
		t.Fail()
	}
}

func TestString2Int64Slice(t *testing.T) {
	s, err := String2Int64Slice("1,2,3", ",")
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 3 {
		t.Fail()
	}
	if s[0] != 1 || s[1] != 2 || s[2] != 3 {
		t.Fail()
	}
}

func TestFindInSet(t *testing.T) {
	s := "1,foo,1.2"
	if !FindInSet(1, s) {
		t.Fail()
	}
	if !FindInSet("foo", s) {
		t.Fail()
	}
	if !FindInSet(1.2, s) {
		t.Fail()
	}
	if !FindInSet("1", s) {
		t.Fail()
	}
	if !FindInSet("1.2", s) {
		t.Fail()
	}
	if FindInSet(2, s) {
		t.Fail()
	}
	if FindInSet("bar", s) {
		t.Fail()
	}
}
