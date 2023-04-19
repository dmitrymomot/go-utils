package utils_test

import (
	"reflect"
	"testing"

	"github.com/dmitrymomot/go-utils"
)

type TestStruct struct {
	StringField string `mytag:"foo"`
	IntField    int    `mytag:"bar"`
}

func TestMapToStruct(t *testing.T) {
	testCases := []struct {
		source map[string]interface{}
		target TestStruct
		want   TestStruct
	}{
		{
			source: map[string]interface{}{
				"foo": "hello",
				"bar": 42,
			},
			target: TestStruct{},
			want: TestStruct{
				StringField: "hello",
				IntField:    42,
			},
		},
	}

	for _, tc := range testCases {
		if err := utils.MapToStruct(tc.source, &tc.target, "mytag"); err != nil {
			t.Errorf("MapToStruct returned an error: %v", err)
		}

		if !reflect.DeepEqual(tc.want, tc.target) {
			t.Errorf("MapToStruct returned wrong result.\nWant: %+v\nGot: %+v", tc.want, tc.target)
		}
	}
}

func TestStructToMap(t *testing.T) {
	testCases := []struct {
		source TestStruct
		want   map[string]interface{}
	}{
		{
			source: TestStruct{
				StringField: "hello",
				IntField:    42,
			},
			want: map[string]interface{}{
				"foo": "hello",
				"bar": 42,
			},
		},
	}

	for _, tc := range testCases {
		got, err := utils.StructToMap(tc.source, "mytag")
		if err != nil {
			t.Errorf("StructToMap returned an error: %v", err)
		}

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("StructToMap returned wrong result.\nWant: %+v\nGot: %+v", tc.want, got)
		}
	}
}
