package arraylist_test

import (
	"testing"

	arraylist "github.com/andersonlemos/coding-interview-go/src/arrayList"
)

func TestInitObject(t *testing.T) {
	arrLst := arraylist.NewArrayList[string]()
	if arrLst == nil {
		t.Log("object was not loaded")
		t.Fail()
	}

}

func TestAddAndsize(t *testing.T) {
	arrLst := arraylist.NewArrayList[string]()
	arrLst.Add("Object")
	if arrLst.Size() <= 0 {
		t.Logf("object size is wrong Went: %v Wrong: %v", arrLst.Size(), 0)
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	arrLst := arraylist.NewArrayList[string]()
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

func TestGetOutOfBounds(t *testing.T) {
	arrLst := arraylist.NewArrayList[string]()

	_, err := arrLst.Get(1)
	if err != nil {
		t.Logf("index out of range")
		t.Skip("error was expected in this case")
	}

}

func TestRemoveOutOfBounds(t *testing.T) {
	arrLst := arraylist.NewArrayList[string]()

	arrLst.Add("a")
	arrLst.Add("b")
	arrLst.Add("c")

	if arrLst.Size() != 3 {
		t.Logf("error adding indexes: expecting 3 has %v", arrLst.Size())
		t.Fail()
	}

	err := arrLst.Remove(3)

	if err != nil {
		t.Log("error on removing index out of bound")
		t.Skip("error was expected in this case")
	}

}
