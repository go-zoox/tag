package attribute

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	a := New("AppName", "string", "", "")
	if a.GetKey() != "AppName" {
		t.Errorf("GetKey() should be AppName, but got %s", a.GetKey())
	}

	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be empty, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "" {
		t.Errorf("GetValue() should be empty, but got %s", a.GetValue())
	}

	if err := a.SetValue("gozoox"); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "gozoox" {
		t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	}
}

func TestAlias(t *testing.T) {
	a := New("AppName", "string", "", "app_name")
	if a.GetKey() != "app_name" {
		t.Errorf("GetKey() should be app_name, but got %s", a.GetKey())
	}

	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be empty, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "" {
		t.Errorf("GetValue() should be empty, but got %s", a.GetValue())
	}

	if err := a.SetValue("gozoox"); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "gozoox" {
		t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	}
}

func TestOmitEmpty(t *testing.T) {
	a := New("AppName", "string", "", "app_name,omitempty")
	if a.GetKey() != "app_name" {
		t.Errorf("GetKey() should be app_name, but got %s", a.GetKey())
	}

	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be empty, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "" {
		t.Errorf("GetValue() should be empty, but got %s", a.GetValue())
	}

	if err := a.SetValue("gozoox"); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "gozoox" {
		t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	}
}

func TestRequire(t *testing.T) {
	a := New("AppName", "string", "", "app_name,required")
	if a.GetKey() != "app_name" {
		t.Errorf("GetKey() should be app_name, but got %s", a.GetKey())
	}

	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be empty, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err == nil {
		t.Errorf("SetValue() should return error, but got nil")
	} else {
		if err.Error() != "app_name is required" {
			t.Errorf("expected error is app_name is required, but got %s", err.Error())
		}
	}

	if a.GetValue() != nil {
		t.Errorf("GetValue() should be empty, but got %s", a.GetValue())
	}

	if err := a.SetValue("gozoox"); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "gozoox" {
		t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	}
}

func TestDefaultValue(t *testing.T) {
	a := New("AppName", "string", "", "app_name,default=gozoox")
	if a.GetKey() != "app_name" {
		t.Errorf("GetKey() should be app_name, but got %s", a.GetKey())
	}

	// // should set before call value
	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "gozoox" {
		t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	}

	if err := a.SetValue("gozoox2"); err != nil {
		t.Errorf("SetValue() should not return error, but got %s", err)
	}

	if a.GetValue() != "gozoox2" {
		t.Errorf("GetValue() should be gozoox2, but got %s", a.GetValue())
	}
}

func TestEnum(t *testing.T) {
	a := New("AppName", "string", "", "app_name,enum=gozoox|gozoox2")
	if a.GetKey() != "app_name" {
		t.Errorf("GetKey() should be app_name, but got %s", a.GetKey())
	}

	// // should set before call value
	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err == nil {
		t.Errorf("SetValue() should return error(%s), but got nil", "app_name must be in enum(gozoox|gozoox2), but empty")
	}

	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	// }

	// if err := a.SetValue("gozoox"); err != nil {
	// 	t.Errorf("SetValue() should not return error, but got %s", err)
	// }

	// if a.GetValue() != "gozoox" {
	// 	t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	// }

	// if err := a.SetValue("gozoox2"); err != nil {
	// 	t.Errorf("SetValue() should not return error, but got %s", err)
	// }

	// if a.GetValue() != "gozoox2" {
	// 	t.Errorf("GetValue() should be gozoox2, but got %s", a.GetValue())
	// }

	// if err := a.SetValue("gozoox3"); err == nil {
	// 	t.Errorf("SetValue() should return error(%s), but got nil", "app_name must be in enum(gozoox|gozoox2), but empty")
	// } else {
	// 	if err.Error() != "app_name(value: gozoox3)) is not in enum(gozoox|gozoox2)" {
	// 		t.Errorf("expected error is app_name(value: gozoox3)) is not in enum(gozoox|gozoox2), but gozoox3, but got %s", err.Error())
	// 	}
	// }
}

func TestStringMinMax(t *testing.T) {
	a := New("Password", "string", "", "password,min=6,max=10")
	if a.GetKey() != "password" {
		t.Errorf("GetKey() should be password, but got %s", a.GetKey())
	}

	// // should set before call value
	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err == nil {
		t.Error("expect error, but got nil")
	} else {
		if err.Error() != "password must be in range(6, 10), but empty" {
			t.Errorf("expected error is password must be in range(6, 10), but empty, but got %s", err.Error())
		}
	}

	if a.GetValue() != nil {
		t.Errorf("expect empty, but got %s", a.GetValue())
	}

	if err := a.SetValue("a"); err == nil {
		t.Error("expect error, but got nil")
	} else {
		if err.Error() != "password must be in range(6, 10), but 1(value: a)" {
			t.Errorf("expected error is password must be in range(6, 10), but 1(value: a), but got %s", err.Error())
		}
	}

	if err := a.SetValue("1234567890a"); err == nil {
		t.Error("expect error, but got nil")
	} else {
		if err.Error() != "password must be in range(6, 10), but 11(value: 1234567890a)" {
			t.Errorf("expected error is password must be in range(6, 10), but 11(value: 1234567890a), but got %s", err.Error())
		}
	}

	if err := a.SetValue("1234567890"); err != nil {
		t.Errorf("expect nil, but got %s", err)
	}

	if a.GetValue() != "1234567890" {
		t.Errorf("expect 1234567890, but got %s", a.GetValue())
	}

	if err := a.SetValue("123456"); err != nil {
		t.Errorf("expect nil, but got %s", err)
	}

	if a.GetValue() != "123456" {
		t.Errorf("expect 123456, but got %s", a.GetValue())
	}
}

func TestNumberMinMax(t *testing.T) {
	a := New("Age", "int", "", "age,min=3,max=18")
	if a.GetKey() != "age" {
		t.Errorf("GetKey() should be age, but got %s", a.GetKey())
	}

	// // should set before call value
	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err == nil {
		t.Error("expect error, but got nil")
	} else {
		if err.Error() != "age must be in range(3, 18), but empty" {
			t.Errorf("expected error is age must be in range(3, 18), but empty, but got %s", err.Error())
		}
	}

	if a.GetValue() != nil {
		t.Errorf("expect empty, but got %s", a.GetValue())
	}

	if err := a.SetValue("1"); err == nil {
		t.Error("expect error, but got nil")
	} else {
		if err.Error() != "age must be in range(3, 18), but 1(value: 1)" {
			t.Errorf("expected error is age must be in range(3, 18), but 1(value: 1), but got %s", err.Error())
		}
	}

	if err := a.SetValue("19"); err == nil {
		t.Error("expect error, but got nil")
	} else {
		if err.Error() != "age must be in range(3, 18), but 19(value: 19)" {
			t.Errorf("expected error is age must be in range(3, 18), but 19(value: 19), but got %s", err.Error())
		}
	}

	if err := a.SetValue("18"); err != nil {
		t.Errorf("expect nil, but got %s", err)
	}

	if a.GetValue() != int64(18) {
		t.Errorf("expect 18, but got %s", a.GetValue())
	}

	if err := a.SetValue("3"); err != nil {
		t.Errorf("expect nil, but got %s", err)
	}

	if a.GetValue() != int64(3) {
		t.Errorf("expect 3, but got %s", a.GetValue())
	}

	if err := a.SetValue("12"); err != nil {
		t.Errorf("expect nil, but got %s", err)
	}

	if a.GetValue() != int64(12) {
		t.Errorf("expect 12, but got %s", a.GetValue())
	}
}

func TestRegExp(t *testing.T) {
	a := New("Email", "string", "", "email,regexp=/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$/")
	if a.GetKey() != "email" {
		t.Errorf("GetKey() should be email, but got %s", a.GetKey())
	}

	// // should set before call value
	// if a.GetValue() != nil {
	// 	t.Errorf("GetValue() should be gozoox, but got %s", a.GetValue())
	// }

	if err := a.SetValue(""); err == nil {
		t.Error("expect error, but got nil")
	} else {
		if err.Error() != "email must be matched with regexp(^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$), but empty" {
			t.Errorf("expected error is email must be matched with regexp(^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$), but empty, but got %s", err.Error())
		}
	}

	if a.GetValue() != nil {
		t.Errorf("expect empty, but got %s", a.GetValue())
	}

	if err := a.SetValue("tobewhatwewant@gmail.com"); err != nil {
		t.Errorf("expect nil, but got %s", err)
	}

	if a.GetValue() != "tobewhatwewant@gmail.com" {
		t.Errorf("expect tobewhatwewant@gmail.com, but got %s", a.GetValue())
	}
}
