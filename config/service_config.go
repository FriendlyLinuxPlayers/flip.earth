package config

import (
	"fmt"
	"reflect"
	"strings"
)

type ServiceConfig map[string]interface{}

func (sc ServiceConfig) Assign(to interface{}) error {
	t := reflect.TypeOf(to)
	v := reflect.ValueOf(to)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("TODO implement error type for when interface is not a struct")
	}
	fc := t.NumField()
	for i := 0; i < fc; i++ {
		tf := t.Field(i)
		if tf.Anonymous {
			continue
		}

		if !v.CanSet() {
			continue
		}
		tag, ok := tf.Tag.Lookup("servconf")
		if !ok {
			continue
		}
		vf := v.Field(i)
		sc.handleField(tag, &vf)

	}

	return nil
}

func (sc ServiceConfig) handleField(tag string, vf *reflect.Value) error {
	tagVals := strings.Split(tag, ",")
	tvc := len(tagVals)

	fName := strings.TrimSpace(tagVals[0])
	if fName == "" {
		return fmt.Errorf("TODO implement error type for empty name field")
	}

	value, ok := sc[fName]

	required := false

	if tvc >= 2 {
		required = tagVals[1] == "required"
	}

	if !ok {
		if required {
			return fmt.Errorf("TODO implement error type for missing required field")
		}
	} else {
		if err := assignValueToField(value, vf, required); err != nil {
			return err
		}

	}

	return nil
}

func assignValueToField(value interface{}, vf *reflect.Value, required bool) error {
	vsc, ok := value.(ServiceConfig)
	if ok {
		//TODO figure out if this mess works
		iface := vf.Interface()
		err := vsc.Assign(&iface)
		if err != nil {
			return err
		}
		value = iface
	}

	valType := reflect.TypeOf(value)
	valVal := reflect.ValueOf(value)
	t := vf.Type()

	if valType.AssignableTo(t) {
		vf.Set(valVal)
	} else if valType.ConvertibleTo(t) {
		converted := valVal.Convert(t)
		vf.Set(converted)
	}
	return nil
}
