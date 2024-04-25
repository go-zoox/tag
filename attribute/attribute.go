package attribute

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Attribute return a Attribute created from the given key + type + detail.
type Attribute struct {
	Key       string
	Type      string
	Alias     string
	Required  bool
	Default   string
	Min       float64
	Max       float64
	Enum      []string
	RegExp    string
	Seperator string
	//
	Value interface{}
	//
	isValueSetted bool
	//
	KeyPathParent string
}

// GetKey returns the key of the attribute.
func (a *Attribute) GetKey() string {
	if a.Alias != "" {
		if a.KeyPathParent != "" {
			return a.KeyPathParent + "." + a.Alias
		}

		return a.Alias
	}

	if a.KeyPathParent != "" {
		return a.KeyPathParent + "." + a.Key
	}
	return a.Key
}

// GetValue returns the value of the attribute.
func (a *Attribute) GetValue() interface{} {
	if !a.isValueSetted {
		panic("value is not setted")
	}

	return a.Value
}

// SetValue sets the value of the attribute.
func (a *Attribute) SetValue(value interface{}) (err error) {
	if value == nil {
		if a.Default != "" {
			value = a.Default
		} else {
			if strings.Contains(a.Type, "struct") {
				//
			} else {
				value = ""
			}
		}
	}

	if !a.isValueSetted {
		a.isValueSetted = true
	}

	switch v := value.(type) {
	case string:
		err = a.setValueString(v)
	case bool:
		err = a.setValueBool(v)
	case int64:
		err = a.setValueInt(v)
	case int:
		err = a.setValueInt(int64(v))
	case float64:
		err = a.setValueFloat(v)
	case float32:
		err = a.setValueFloat(float64(v))
	default:
		a.Value = v
	}

	return
}

// SetValue sets the value of the attribute.
func (a *Attribute) setValueString(value string) (err error) {
	// fmt.Println("setValueString:", value, a.GetKey())
	if strings.Index(a.Type, "struct") != -1 {
		// return fmt.Errorf("type(key: %s) is struct, can't set with string value(%s)", a.GetKey(), value)

		// ignore struct
		return nil
	}

	// value is empty
	if value == "" {
		if a.Default != "" {
			a.Value = a.Default
		}

		if a.Required {
			return fmt.Errorf("%s is required", a.GetKey())
		}

		if a.Enum != nil {
			return fmt.Errorf("%s must be in enum(%s), but empty", a.GetKey(), strings.Join(a.Enum, "|"))
		}

		if a.Min != 0 || a.Max != 0 {
			if a.Type == "string" {
				return fmt.Errorf("%s must be in range(%d, %d), but empty", a.GetKey(), int(a.Min), int(a.Max))
			}

			switch a.Type {
			case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
				return fmt.Errorf("%s must be in range(%d, %d), but empty", a.GetKey(), int(a.Min), int(a.Max))
			case "float", "float32", "float64":
				return fmt.Errorf("%s must be in range(%f, %f), but empty", a.GetKey(), a.Min, a.Max)
			}
		}

		if a.RegExp != "" {
			return fmt.Errorf("%s must be matched with regexp(%s), but empty", a.GetKey(), a.RegExp)
		}

		if a.Value == nil {
			a.Value = value // empty string
		}
	} else {
		// value is not empty
		// 1. check enum
		if a.Enum != nil {
			isInEnum := false
			for _, v := range a.Enum {
				if v == value {
					isInEnum = true
					break
				}
			}

			if !isInEnum {
				return fmt.Errorf("%s(value: %s)) is not in enum(%s)", a.GetKey(), value, strings.Join(a.Enum, "|"))
			}
		}

		// 2. check range
		//	1. string => length range
		//  2. int => value range
		if a.Min != 0 || a.Max != 0 {
			switch a.Type {
			case "string":
				if valueLen := len(value); valueLen < int(a.Min) || valueLen > int(a.Max) {
					err = fmt.Errorf("%s must be in range(%d, %d), but %d(value: %s)", a.GetKey(), int(a.Min), int(a.Max), valueLen, value)
				}
			case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float", "float32", "float64":
				valueX, errx := strconv.ParseFloat(value, 64)
				if errx != nil {
					err = fmt.Errorf("%s is invalid with min(%f) and max(%f)", a.GetKey(), a.Min, a.Max)
				} else if valueX < a.Min || valueX > a.Max {
					switch a.Type {
					case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
						return fmt.Errorf("%s must be in range(%d, %d), but %d(value: %s)", a.GetKey(), int(a.Min), int(a.Max), int(valueX), value)
					case "float", "float32", "float64":
						return fmt.Errorf("%s must be in range(%f, %f), but %f(value: %s)", a.GetKey(), a.Min, a.Max, valueX, value)
					}
				}
			}

			if err != nil {
				return err
			}
		}

		// 3. check regexp (string)
		if a.RegExp != "" {
			if ok, err := regexp.MatchString(a.RegExp, value); err != nil {
				return err
			} else if !ok {
				return fmt.Errorf("%s is invalid with regexp(%s)", a.GetKey(), a.RegExp)
			}
		}

		// if a.Value == "" {
		// 	a.Value = value
		// }
		a.Value = value
	}

	// Correct the value by type
	switch a.Type {
	case "string":
		// do nothing
	case "float64":
		if a.Value == "" {
			a.Value = float64(0)
		} else {
			a.Value, err = strconv.ParseFloat(a.Value.(string), 64)
			if err != nil {
				return fmt.Errorf("%s is not float", a.Key)
			}
		}
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		if a.Value == "" {
			a.Value = int64(0)
		} else {
			a.Value, err = strconv.ParseInt(a.Value.(string), 10, 64)
			if err != nil {
				return fmt.Errorf("%s is not int", a.Key)
			}
		}

	case "bool":
		if a.Value == "" {
			a.Value = false
		} else {
			a.Value, err = strconv.ParseBool(a.Value.(string))
			if err != nil {
				return fmt.Errorf("%s is not bool", a.Key)
			}
		}
	// slice
	case "[]string":
		if value == "" {
			a.Value = nil
		} else {
			if a.Seperator == "" {
				a.Seperator = ","
			}

			a.Value = strings.Split(value, a.Seperator)
		}
	case "[]int":
		if value == "" {
			a.Value = nil
		} else {
			if a.Seperator == "" {
				a.Seperator = ","
			}

			strs := strings.Split(value, a.Seperator)
			ints := make([]int, len(strs))
			for i, v := range strs {
				ints[i], err = strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s is not int", a.Key)
				}
			}

			a.Value = ints
		}
	case "[]int64":
		if value == "" {
			a.Value = nil
		} else {
			if a.Seperator == "" {
				a.Seperator = ","
			}

			strs := strings.Split(value, a.Seperator)
			ints := make([]int64, len(strs))
			for i, v := range strs {
				ints[i], err = strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s is not int64", a.Key)
				}
			}

			a.Value = ints
		}
	case "[]float64":
		if value == "" {
			a.Value = nil
		} else {
			if a.Seperator == "" {
				a.Seperator = ","
			}

			strs := strings.Split(value, a.Seperator)
			floats := make([]float64, len(strs))
			for i, v := range strs {
				floats[i], err = strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s is not float64", a.Key)
				}
			}

			a.Value = floats
		}
	default:
		fmt.Println("type:", a.Type)
		a.Value = nil
	}

	return nil
}

func (a *Attribute) setValueBool(value bool) (err error) {
	if a.Type != "bool" {
		return fmt.Errorf("type of %s is not bool", a.GetKey())
	}

	a.Value = value

	return nil
}

func (a *Attribute) setValueInt(value int64) (err error) {
	// if a.Type != "int64" {
	// 	return fmt.Errorf("type of %s is not int64", a.GetKey())
	// }

	// switch a.Type {
	// case "int":
	// 	a.Value = int(value)
	// case "int64":
	// 	a.Value = value
	// default:
	// 	return fmt.Errorf("type of %s is not int or int64", a.GetKey())
	// }

	if a.Min != 0 || a.Max != 0 {
		if float64(value) < a.Min || float64(value) > a.Max {
			return fmt.Errorf("%s must be in range(%d, %d), but %d", a.GetKey(), int(a.Min), int(a.Max), value)
		}
	}

	a.Value = value

	return nil
}

func (a *Attribute) setValueFloat(value float64) (err error) {
	// if a.Type != "float64" {
	// 	return fmt.Errorf("type of %s is not float64", a.GetKey())
	// }

	if a.Min != 0 || a.Max != 0 {
		if value < a.Min || value > a.Max {
			return fmt.Errorf("%s must be in range(%f, %f), but %f", a.GetKey(), a.Min, a.Max, value)
		}
	}

	a.Value = value

	return nil
}

// New creates a new Attribute
//
//	type struct {
//		AppName  string `ini:"app_name,omitempty"`
//	 LogLevel string `ini:"log_level,default=DEBUG"`
//	 Secret  string 	`ini:"secret,min=8,max=16"`
//	 Type 		string 	`ini:"type,enum=male|female"`
//		RegExp  string  `ini:"regexp,regexp=/xxx/"`
//	}
//
// key: AppName
// typ: string
// detail: "app_name,omitempty"
//
// key: LogLevel
// typ: string
// detail: "log_level,default=DEBUG"
func New(key string, typ string, keyPathParent string, detail string) *Attribute {
	parts := strings.Split(detail, ",")
	var alias string
	var required bool
	var defaultValue string
	var min float64
	var max float64
	var enum []string
	var regexp string
	var seperator string

	var err error

	if len(parts) > 0 {
		for _, part := range parts {
			if part == "omitempty" {
				required = false
			} else if part == "required" {
				required = true
			} else if strings.Contains(part, "=") {
				kv := strings.Split(part, "=")
				if kv[0] == "default" {
					defaultValue = kv[1]
				} else if kv[0] == "min" {
					min, err = strconv.ParseFloat(kv[1], 64)
					if err != nil {
						panic(err)
					}
				} else if kv[0] == "max" {
					max, err = strconv.ParseFloat(kv[1], 64)
					if err != nil {
						panic(err)
					}
				} else if kv[0] == "enum" {
					enum = strings.Split(kv[1], "|")
				} else if kv[0] == "regexp" {
					reparts := strings.Split(kv[1], "/")
					if len(reparts) == 3 {
						regexp = reparts[1]
					}
				} else if kv[0] == "seperator" {
					if len(kv) != 2 {
						panic("seperator must have a value")
					}

					seperator = kv[1]
					if seperator == "" {
						panic("seperator must have a value")
					}
				}
			} else {
				if alias == "" {
					alias = part
				}
			}
		}
	}

	return &Attribute{
		Key:           key,
		Type:          typ,
		Alias:         alias,
		Required:      required,
		Default:       defaultValue,
		Min:           min,
		Max:           max,
		Enum:          enum,
		RegExp:        regexp,
		Seperator:     seperator,
		KeyPathParent: keyPathParent,
	}
}
