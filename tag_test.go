package tag

import (
	"testing"

	"github.com/go-zoox/core-utils/object"
)

type TestStruct struct {
	AppName  string `custom_struct_tag:"app_name,omitempty"`
	LogLevel string `custom_struct_tag:"log_level"`
	// Age      int
	Redis struct {
		IP   string `custom_struct_tag:"ip"`
		Port int64  `custom_struct_tag:"port"`
		// Port int64  `custom_struct_tag:"port"`
		Database string `custom_struct_tag:"database,default=zoox"`
	} `custom_struct_tag:"redis"`
	Ports []int64 `custom_struct_tag:"ports"`
}

type TestStructDataSource struct {
}

var TestStructDataSourceData = map[string]interface{}{
	"app_name":  "gozoox",
	"log_level": "DEBUG",
	"redis": map[string]interface{}{
		"ip":   "127.0.0.1",
		"port": "6739",
		// "port": 6739,
	},
	"ports": []int64{
		6739,
		6740,
	},
	"type_transform": "666",
}

func (t *TestStructDataSource) Get(key string) interface{} {
	return object.Get(TestStructDataSourceData, key)
}

func TestTag(t *testing.T) {
	var test TestStruct
	tag := New("custom_struct_tag", &TestStructDataSource{})
	if err := tag.Decode(&test); err != nil {
		t.Error(err)
	}

	if test.AppName != "gozoox" {
		t.Errorf("AppName should be gozoox, but got %s", test.AppName)
	}

	if test.LogLevel != "DEBUG" {
		t.Errorf("LogLevel should be DEBUG, but got %s", test.LogLevel)
	}

	if test.Redis.IP != "127.0.0.1" {
		t.Errorf("Redis.IP should be 127.0.0.1, but got %s", test.Redis.IP)
	}

	if test.Redis.Port != 6739 {
		t.Errorf("Redis.Port should be 6739, but got %d", test.Redis.Port)
	}

	if test.Redis.Database != "zoox" {
		t.Errorf("Redis.Database should be zoox, but got %s", test.Redis.Database)
	}
}

func TestDefaultValue(t *testing.T) {
	type TestStruct struct {
		AppName2 string `custom_struct_tag:"app_name_2,default=gozoox2"`
		Age      int64  `custom_struct_tag:"age,default=18"`
		Bool     bool   `custom_struct_tag:"bool,default=true"`
	}

	var test TestStruct
	tag := New("custom_struct_tag", &TestStructDataSource{})
	if err := tag.Decode(&test); err != nil {
		t.Error(err)
	}

	if test.AppName2 != "gozoox2" {
		t.Errorf("AppName2 should be gozoox2, but got %s", test.AppName2)
	}

	if test.Age != 18 {
		t.Errorf("Age should be 18, but got %d", test.Age)
	}

	if !test.Bool {
		t.Errorf("Bool should be true, but got %v", test.Bool)
	}
}

func TestRequired(t *testing.T) {
	var test1 struct {
		AppName2 string `custom_struct_tag:"app_name_2,required"`
	}
	if err := New("custom_struct_tag", &TestStructDataSource{}).Decode(&test1); err == nil {
		t.Error("should be error, but got nil")
	}

	var test2 struct {
		Age int64 `custom_struct_tag:"age,required"`
	}
	if err := New("custom_struct_tag", &TestStructDataSource{}).Decode(&test2); err == nil {
		t.Error("should be error, but got nil")
	}

	var test3 struct {
		Bool bool `custom_struct_tag:"bool,required"`
	}
	if err := New("custom_struct_tag", &TestStructDataSource{}).Decode(&test3); err == nil {
		t.Error("should be error, but got nil")
	}
}

func TestStringLengthMinMax(t *testing.T) {
	var test struct {
		AppName string `custom_struct_tag:"app_name,min=10,max=30"`
	}
	if err := New("custom_struct_tag", &TestStructDataSource{}).Decode(&test); err == nil {
		t.Error("should be error, but got nil")
	}
}

func TestStringRegExp(t *testing.T) {
	var test struct {
		AppName string `custom_struct_tag:"app_name,regexp=/^gozoox2/"`
	}
	if err := New("custom_struct_tag", &TestStructDataSource{}).Decode(&test); err == nil {
		t.Error("should be error, but got nil")
	}
}

func TestStringEnum(t *testing.T) {
	var test struct {
		AppName string `custom_struct_tag:"app_name,enum=gozoox3|gozoox2"`
	}

	if err := New("custom_struct_tag", &TestStructDataSource{}).Decode(&test); err == nil {
		t.Error("should be error, but got nil")
	}
}

func TestTypeTransform(t *testing.T) {
	var test struct {
		TypeTransform       int64  `custom_struct_tag:"type_transform"`
		TypeTransformInt    int    `custom_struct_tag:"type_transform"`
		TypeTransformUInt   uint   `custom_struct_tag:"type_transform"`
		TypeTransformUInt32 uint32 `custom_struct_tag:"type_transform"`
		TypeTransformUInt64 uint64 `custom_struct_tag:"type_transform"`
	}
	if err := New("custom_struct_tag", &TestStructDataSource{}).Decode(&test); err != nil {
		t.Error(err)
	}

	if test.TypeTransform != 666 {
		t.Errorf("TypeTransform should be 666, but got %d", test.TypeTransform)
	}

	if test.TypeTransformInt != 666 {
		t.Errorf("TypeTransformInt should be 666, but got %d", test.TypeTransformInt)
	}

	if test.TypeTransformUInt != 666 {
		t.Errorf("TypeTransformInt should be 666, but got %d", test.TypeTransformUInt)
	}

	if test.TypeTransformUInt32 != 666 {
		t.Errorf("TypeTransformInt should be 666, but got %d", test.TypeTransformUInt32)
	}

	if test.TypeTransformUInt64 != 666 {
		t.Errorf("TypeTransformInt should be 666, but got %d", test.TypeTransformUInt64)
	}
}
