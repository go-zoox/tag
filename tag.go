package tag

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-zoox/tag/attribute"
	"github.com/go-zoox/tag/datasource"
)

// Tag is a struct tag parser and decoder
type Tag struct {
	Name       string
	DataSource datasource.DataSource
}

// New creates a new Tag
func New(name string, dataSource datasource.DataSource) *Tag {
	return &Tag{
		Name:       name,
		DataSource: dataSource,
	}
}

// Decode decodes the given struct pointer from data source.
func (t *Tag) Decode(ptr interface{}) error {
	return t.decodeR(ptr, "")
}

func (t *Tag) decodeR(ptr interface{}, keyPathParent string) error {
	tagName, dataSource := t.Name, t.DataSource

	rt := reflect.TypeOf(ptr).Elem()
	rv := reflect.ValueOf(ptr).Elem()

	// example:
	// redis.host
	// config.redis.host
	for i := 0; i < rt.NumField(); i++ {
		rtt := rt.Field(i)
		rvv := rv.Field(i)

		// xkey := rtt.Name
		// xalias := rtt.Tag.Get(tagName)
		// xtype := rtt.Type

		// fmt.Println("typ:", xkey, xalias, xtype.String())

		attribute := attribute.New(rtt.Name, rtt.Type.String(), keyPathParent, rtt.Tag.Get(tagName))
		// fmt.Println("keyPathParent:", keyPathParent, rtt.Name, attribute.GetKey())
		// fmt.Println("key:", attribute.GetKey())
		if err := attribute.SetValue(dataSource.Get(attribute.GetKey())); err != nil {
			return err
		}

		if err := t.setValue(rtt.Type, rvv, attribute); err != nil {
			return err
		}
	}

	return nil
}

func (t *Tag) setValue(rt reflect.Type, rv reflect.Value, attribute *attribute.Attribute) error {
	value := attribute.GetValue()
	if value == nil {
		return nil
	}

	switch rv.Kind() {
	case reflect.String:
		rv.SetString(value.(string))

	case reflect.Bool:
		rv.SetBool(value.(bool))

	case reflect.Int64:
		if err := t.setValueInt(rv, value); err != nil {
			return err
		}

	case reflect.Float64:
		if err := t.setValueFloat(rv, value); err != nil {
			return err
		}

	case reflect.Struct:
		if err := t.decodeR(rv.Addr().Interface(), attribute.GetKey()); err != nil {
			// return fmt.Errorf("struct decode error at %s: %s", Key, err.Error())
			// return fmt.Errorf("%s decode error: %s", attribute.GetKey(), err.Error())
			return err
		}

	case reflect.Slice:
		if err := t.setValueSlice(rt, rv, value, attribute); err != nil {
			return err
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Errorf("type(%s) is not supported at %s, please int64", rv.Kind(), attribute.GetKey())

	case reflect.Float32:
		return fmt.Errorf("type(%s) is not supported at %s, please float64", rv.Kind(), attribute.GetKey())

	default:
		return fmt.Errorf("type(%s) is not supported at %s, fatal err", rv.Kind(), attribute.GetKey())
	}

	return nil
}

func (t *Tag) setValueInt(rv reflect.Value, value any) error {
	switch v := value.(type) {
	case int64:
		rv.SetInt(v)
	case int:
		rv.SetInt(int64(v))
	case float32:
		rv.SetFloat(float64(v))
	case int8:
		rv.SetInt(int64(v))
	case int16:
		rv.SetInt(int64(v))
	case int32:
		rv.SetInt(int64(v))
	case uint:
		rv.SetInt(int64(v))
	case uint8:
		rv.SetInt(int64(v))
	case uint16:
		rv.SetInt(int64(v))
	case uint32:
		rv.SetInt(int64(v))
	case uint64:
		rv.SetInt(int64(v))
	default:
		return fmt.Errorf("setValueInt unknown value type: %s", reflect.TypeOf(value).Kind())
	}

	return nil
}

func (t *Tag) setValueFloat(rv reflect.Value, value any) error {
	switch v := value.(type) {
	case float64:
		rv.SetFloat(v)
	case float32:
		rv.SetFloat(float64(v))
	default:
		return fmt.Errorf("setValueFloat unknown value type: %s", reflect.TypeOf(value).Kind())
	}

	return nil
}

func (t *Tag) setValueSlice(rt reflect.Type, rv reflect.Value, value any, attribute *attribute.Attribute) error {
	s := reflect.ValueOf(value)
	for index := 0; index < s.Len(); index++ {
		switch v := s.Index(index).Interface().(type) {
		case string, bool, int64, float64:
			rv.Set(reflect.Append(rv, reflect.ValueOf(v)))
		case float32:
			rv.Set(reflect.Append(rv, reflect.ValueOf(float64(v))))
		case int:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case int32:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case int16:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case int8:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case uint:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case uint64:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case uint32:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case uint16:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case uint8:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		case uintptr:
			rv.Set(reflect.Append(rv, reflect.ValueOf(int64(v))))
		default:
			value := reflect.New(rt.Elem())
			if err := t.decodeR(reflect.Value(value).Interface(), attribute.GetKey()+"."+strconv.Itoa(index)); err != nil {
				return fmt.Errorf("%s is not slice(%s)", attribute.Key, err.Error())
			}

			// j, _ := json.MarshalIndent(value.Elem().Interface(), "", "  ")
			// fmt.Println("value:", string(j))

			rv.Set(reflect.Append(rv, value.Elem()))
		}
	}

	return nil
}
