package arraylist_test

import (
	"testing"

	arraylist "github.com/andersonlemos/coding-interview-go/src/arrayList"
)

func TestInitObject(t *testing.T) {
	arrLst := arraylist.NewArrayList()
	if arrLst == nil {
		t.Log("object was not loaded")
		t.Fail()
	}

}

func TestAddAndsize(t *testing.T) {
	arrLst := arraylist.NewArrayList()
	arrLst.Add("Object")
	if arrLst.Size() <= 0 {
		t.Logf("object size is wrong Went: %v Wrong: %v", arrLst.Size(), 0)
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	arrLst := arraylist.NewArrayList()
	arrLst.Add("Object")
	item, err := arrLst.Get(0)
	if err != nil {
		t.Logf("error on get item: %v", err)
		t.Fail()
	}

	if item != "Object" {
		t.Logf("wrong Went: 'Object' Wrong: %v", item)
		t.Fail()
	}
}
