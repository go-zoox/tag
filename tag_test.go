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
	Ports     []int64             `custom_struct_tag:"ports"`
	Maps      map[string]string   `custom_struct_tag:"maps"`
	Providers map[string]Provider `custom_struct_tag:"providers"`
	Users     []User              `custom_struct_tag:"users"`
}

type Provider struct {
	ClientID     string `custom_struct_tag:"client_id"`
	ClientSecret string `custom_struct_tag:"client_secret"`
}

type User struct {
	Name string `custom_struct_tag:"name"`
	Age  int    `custom_struct_tag:"age"`
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
	"maps": map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	},
	"providers": map[string]any{
		"google": map[string]any{
			"client_id":     "google_client_id",
			"client_secret": "google_client_secret",
		},
		"facebook": map[string]any{
			"client_id":     "facebook_client_id",
			"client_secret": "facebook_client_secret",
		},
		"github": map[string]any{
			"client_id":     "github_client_id",
			"client_secret": "github_client_secret",
		},
	},
	"users": []map[string]any{
		{
			"name": "user1",
			"age":  18,
		},
		{
			"name": "user2",
			"age":  20,
		},
	},
	"type_transform": "666",
}

func (t *TestStructDataSource) Get(key string) (interface{}, error) {
	return object.Get(TestStructDataSourceData, key), nil
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

	if len(test.Ports) != 2 {
		t.Fatalf("Ports length must be %d, bug got %d", 2, len(test.Ports))
	}

	if len(test.Maps) != 2 {
		t.Fatalf("Maps length must be %d, bug got %d", 2, len(test.Maps))
	}

	if test.Maps["key1"] != "value1" {
		t.Fatalf("Maps[key1] must be %s, bug got %s", "value1", test.Maps["key1"])
	}

	if test.Maps["key2"] != "value2" {
		t.Fatalf("Maps[key2] must be %s, bug got %s", "value2", test.Maps["key2"])
	}

	if test.Maps["key3"] != "" {
		t.Fatalf("Maps[key3] must be %s, bug got %s", "", test.Maps["key3"])
	}

	if len(test.Providers) != 3 {
		t.Fatalf("Providers length must be %d, bug got %d", 3, len(test.Providers))
	}

	if test.Providers["google"].ClientID != "google_client_id" {
		t.Fatalf("Providers[google].ClientID must be %s, bug got %s", "google_client_id", test.Providers["google"].ClientID)
	}

	if test.Providers["google"].ClientSecret != "google_client_secret" {
		t.Fatalf("Providers[google].ClientSecret must be %s, bug got %s", "google_client_secret", test.Providers["google"].ClientSecret)
	}

	if test.Providers["facebook"].ClientID != "facebook_client_id" {
		t.Fatalf("Providers[facebook].ClientID must be %s, bug got %s", "facebook_client_id", test.Providers["facebook"].ClientID)
	}

	if test.Providers["facebook"].ClientSecret != "facebook_client_secret" {
		t.Fatalf("Providers[facebook].ClientSecret must be %s, bug got %s", "facebook_client_secret", test.Providers["facebook"].ClientSecret)
	}

	if test.Providers["github"].ClientID != "github_client_id" {
		t.Fatalf("Providers[github].ClientID must be %s, bug got %s", "github_client_id", test.Providers["github"].ClientID)
	}

	if test.Providers["github"].ClientSecret != "github_client_secret" {
		t.Fatalf("Providers[github].ClientSecret must be %s, bug got %s", "github_client_secret", test.Providers["github"].ClientSecret)
	}

	if len(test.Users) != 2 {
		t.Fatalf("Users length must be %d, bug got %d", 2, len(test.Users))
	}

	if test.Users[0].Name != "user1" {
		t.Fatalf("Users[0].Name must be %s, bug got %s", "user1", test.Users[0].Name)
	}

	if test.Users[0].Age != 18 {
		t.Fatalf("Users[0].Age must be %d, bug got %d", 18, test.Users[0].Age)
	}

	if test.Users[1].Name != "user2" {
		t.Fatalf("Users[1].Name must be %s, bug got %s", "user2", test.Users[1].Name)
	}

	if test.Users[1].Age != 20 {
		t.Fatalf("Users[1].Age must be %d, bug got %d", 20, test.Users[1].Age)
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
