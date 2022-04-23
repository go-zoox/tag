package tag

import (
	"strings"
	"testing"
)

type TestStruct struct {
	AppName  string `custom_struct_tag:"app_name,omitempty"`
	LogLevel string `custom_struct_tag:"log_level"`
	// Age      int
	Redis struct {
		IP   string `custom_struct_tag:"ip"`
		Port int    `custom_struct_tag:"port"`
		// Port int64  `custom_struct_tag:"port"`
	} `custom_struct_tag:"redis"`
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
}

func (t *TestStructDataSource) Get(key string) interface{} {
	if strings.Contains(key, ".") {
		keys := strings.Split(key, ".")
		keyLength := len(keys)
		tmp := TestStructDataSourceData
		for index, k := range keys {
			if v, ok := tmp[k]; ok {
				if index == keyLength-1 {
					return v
				}

				tmp = v.(map[string]interface{})
			} else {
				return nil
			}
		}
	} else {
		if v, ok := TestStructDataSourceData[key]; ok {
			return v
		}
	}

	return nil
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
}
