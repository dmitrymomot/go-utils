package utils_test

import (
	"testing"

	"github.com/dmitrymomot/go-utils"
)

type Person struct {
	Name string
	Age  int
}

func TestFullyQualifiedStructName(t *testing.T) {
	// Test non-pointer struct type
	p := Person{}
	expected1 := "utils_test.Person"
	actual1 := utils.FullyQualifiedStructName(p)
	if actual1 != expected1 {
		t.Errorf("Expected %q but got %q", expected1, actual1)
	}

	// Test pointer struct type
	ptr := &Person{}
	expected2 := "utils_test.Person"
	actual2 := utils.FullyQualifiedStructName(ptr)
	if actual2 != expected2 {
		t.Errorf("Expected %q but got %q", expected2, actual2)
	}
}

func TestStructName(t *testing.T) {
	// Test non-pointer struct type
	p := Person{}
	expected1 := "Person"
	actual1 := utils.StructName(p)
	if actual1 != expected1 {
		t.Errorf("Expected %q but got %q", expected1, actual1)
	}

	// Test pointer struct type
	ptr := &Person{}
	expected2 := "Person"
	actual2 := utils.StructName(ptr)
	if actual2 != expected2 {
		t.Errorf("Expected %q but got %q", expected2, actual2)
	}
}

type NamedPerson struct {
	FirstName string
	Age       int
}

func (p NamedPerson) Name() string {
	return "NamedPerson"
}

func TestNamedStruct(t *testing.T) {
	// Define a fallback function
	fallback := func(v interface{}) string {
		return "fallback"
	}

	// Test non-pointer namedStruct type
	p := NamedPerson{}
	expected1 := "NamedPerson"
	actual1 := utils.NamedStruct(fallback)(p)
	if actual1 != expected1 {
		t.Errorf("Expected %q but got %q", expected1, actual1)
	}

	// Test pointer namedStruct type
	ptr := &NamedPerson{}
	expected2 := "NamedPerson"
	actual2 := utils.NamedStruct(fallback)(ptr)
	if actual2 != expected2 {
		t.Errorf("Expected %q but got %q", expected2, actual2)
	}

	// Test non-namedStruct type
	i := 42
	expected3 := "fallback"
	actual3 := utils.NamedStruct(fallback)(i)
	if actual3 != expected3 {
		t.Errorf("Expected %q but got %q", expected3, actual3)
	}
}
