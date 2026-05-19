package attribute

import "testing"

func TestScalarPointerHelpers(t *testing.T) {
	cases := []struct {
		typ  string
		want bool
	}{
		{"*bool", true},
		{"*int", true},
		{"*int64", true},
		{"*uint32", true},
		{"*float64", true},
		{"*string", true},
		{"bool", false},
		{"*[]int", false},
		{"*struct { X int }", false},
		{"[]*int", false},
	}
	for _, c := range cases {
		if got := isScalarPointer(c.typ); got != c.want {
			t.Errorf("isScalarPointer(%q) = %v, want %v", c.typ, got, c.want)
		}
	}

	if baseType("*int64") != "int64" {
		t.Fatalf("baseType(*int64) = %q", baseType("*int64"))
	}
	if !matchesScalarType("*bool", "bool") || matchesScalarType("*bool", "int") {
		t.Fatal("matchesScalarType bool")
	}
	if !isIntBaseType("uint16") || isFloatBaseType("int") {
		t.Fatal("isIntBaseType / isFloatBaseType")
	}
}

func TestSetValueScalarPointerOmitted(t *testing.T) {
	types := []string{"*bool", "*int", "*int64", "*uint", "*float64", "*string"}
	for _, typ := range types {
		t.Run(typ, func(t *testing.T) {
			a := New("Field", typ, "", "field")
			if err := a.SetValue(nil); err != nil {
				t.Fatal(err)
			}
			if a.GetValue() != nil {
				t.Fatalf("expected nil, got %v", a.GetValue())
			}
		})
	}
}

func TestSetValueScalarPointerExplicit(t *testing.T) {
	t.Run("*bool", func(t *testing.T) {
		a := New("Flag", "*bool", "", "flag")
		for _, v := range []bool{true, false} {
			if err := a.SetValue(v); err != nil {
				t.Fatal(err)
			}
			if a.GetValue() != v {
				t.Fatalf("got %v want %v", a.GetValue(), v)
			}
		}
	})

	t.Run("*int zero", func(t *testing.T) {
		a := New("Port", "*int", "", "port")
		if err := a.SetValue(int64(0)); err != nil {
			t.Fatal(err)
		}
		if a.GetValue() != int64(0) {
			t.Fatalf("got %v", a.GetValue())
		}
	})

	t.Run("*float64", func(t *testing.T) {
		a := New("Rate", "*float64", "", "rate")
		if err := a.SetValue(float64(1.5)); err != nil {
			t.Fatal(err)
		}
		if a.GetValue() != float64(1.5) {
			t.Fatalf("got %v", a.GetValue())
		}
	})

	t.Run("*string empty", func(t *testing.T) {
		a := New("Note", "*string", "", "note")
		if err := a.SetValue(""); err != nil {
			t.Fatal(err)
		}
		if a.GetValue() != "" {
			t.Fatalf("got %q", a.GetValue())
		}
	})
}

func TestSetValueScalarPointerStringCoercion(t *testing.T) {
	t.Run("*int from string", func(t *testing.T) {
		a := New("Port", "*int", "", "port")
		if err := a.SetValue("9090"); err != nil {
			t.Fatal(err)
		}
		if a.GetValue() != int64(9090) {
			t.Fatalf("got %v", a.GetValue())
		}
	})

	t.Run("*bool from string", func(t *testing.T) {
		a := New("Flag", "*bool", "", "flag")
		if err := a.SetValue("true"); err != nil {
			t.Fatal(err)
		}
		if a.GetValue() != true {
			t.Fatalf("got %v", a.GetValue())
		}
	})

	t.Run("*float64 from string", func(t *testing.T) {
		a := New("Rate", "*float64", "", "rate")
		if err := a.SetValue("3.14"); err != nil {
			t.Fatal(err)
		}
		if a.GetValue() != float64(3.14) {
			t.Fatalf("got %v", a.GetValue())
		}
	})
}

func TestSetValueScalarPointerRequired(t *testing.T) {
	a := New("Flag", "*bool", "", "flag,required")
	err := a.SetValue(nil)
	if err == nil {
		t.Fatal("expected required error")
	}
	if err.Error() != "flag is required" {
		t.Fatalf("got %q", err.Error())
	}
}

func TestSetValueBoolRejectsNonBoolPointer(t *testing.T) {
	a := New("Count", "*int", "", "count")
	if err := a.SetValue(true); err == nil {
		t.Fatal("expected error when setting bool on *int field")
	}
}

func TestSetValueNonPointerBoolStillUsesDefaultOnFalse(t *testing.T) {
	a := New("Flag", "bool", "", "flag,default=true")
	if err := a.SetValue(false); err != nil {
		t.Fatal(err)
	}
	if a.GetValue() != true {
		t.Fatalf("non-pointer false with default=true should use default, got %v", a.GetValue())
	}
}
