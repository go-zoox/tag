package tag

import (
	"fmt"
	"reflect"

	"github.com/go-zoox/tag/attribute"
	typ "github.com/go-zoox/tag/type"
)

// Tag is a struct tag parser and decoder
type Tag struct {
	Name       string
	DataSource typ.DataSource
}

// New creates a new Tag
func New(name string, dataSource typ.DataSource) *Tag {
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
		fmt.Println("key:", attribute.GetKey())
		if err := attribute.SetValue(dataSource.Get(attribute.GetKey())); err != nil {
			return err
		}

		if err := t.setValue(keyPathParent, rtt.Type, rvv, attribute); err != nil {
			return err
		}
	}

	return nil
}

func (t *Tag) setValue(keyPathParent string, rt reflect.Type, rv reflect.Value, attribute *attribute.Attribute) error {
	Key, Value := attribute.Key, attribute.GetValue()
	if Value == nil {
		return nil
	}

	switch rv.Kind() {
	case reflect.String:
		rv.SetString(Value.(string))

	case reflect.Bool:
		rv.SetBool(Value.(bool))

	case reflect.Int, reflect.Int64:
		rv.SetInt(Value.(int64))

	case reflect.Float64:
		rv.SetFloat(Value.(float64))

	case reflect.Struct:
		var newKeyParentPath string
		if keyPathParent == "" {
			newKeyParentPath = attribute.GetKey()
		} else {
			newKeyParentPath = keyPathParent + "." + attribute.GetKey()
		}

		if err := t.decodeR(rv.Addr().Interface(), newKeyParentPath); err != nil {
			return fmt.Errorf("%s is not struct(%s)", Key, err.Error())
		}

	default:
		return fmt.Errorf("%s is not supported", Key)
	}

	return nil
}
