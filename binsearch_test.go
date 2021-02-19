package binsearchgo

import (
	"fmt"
	"testing"
)

func TestOneElement(t *testing.T) {

	array := []int{1}
	service := NewBinSearchService(array)
	ok, pos := service.Search(1)

	if !ok || pos != 0 {
		t.Fail()
	}

	ok, pos = service.Search(10)
	if ok {
		t.Fail()
	}
}

func TestTwoElements(t *testing.T) {

	array := []int{0, 1}
	service := NewBinSearchService(array)
	ok, pos := service.Search(1)
	fmt.Println(ok, pos)
	if !ok || pos != 1 {
		t.Fail()
	}

	ok, pos = service.Search(10)
	fmt.Println(ok, pos)
	if ok {
		t.Fail()
	}
}

func TestSixElements(t *testing.T) {

	array := []int{0, 1, 2, 3, 4, 5}
	service := NewBinSearchService(array)
	ok, pos := service.Search(3)
	fmt.Println(ok, pos)
	if !ok || pos != 3 {
		t.Fail()
	}

	ok, pos = service.Search(10)
	fmt.Println(ok, pos)
	if ok {
		t.Fail()
	}
}
