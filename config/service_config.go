package config

import (
	"reflect"
	"strings"
)

type ServiceConfig map[string]interface{}

func (sc ServiceConfig) Assign(to interface{}) error {
	vPtr := reflect.ValueOf(to)
	if vPtr.Kind() != reflect.Ptr {
		return &InvalidArgumentError{vPtr.Kind(), false}
	}
	v := vPtr.Elem()
	t := reflect.TypeOf(v.Interface())
	if v.Kind() != reflect.Struct {
		return &InvalidArgumentError{v.Kind(), true}
	}
	fc := t.NumField()
	for i := 0; i < fc; i++ {
		tf := t.Field(i)
		vf := v.Field(i)
		if tf.Anonymous {
			continue
		}

		if !vf.CanSet() {
			continue
		}
		if _, ok := tf.Tag.Lookup("servconf"); !ok {
			continue
		}
		if err := sc.handleField(&tf, &vf); err != nil {
			return err
		}

	}

	return nil
}

func (sc ServiceConfig) handleField(tf *reflect.StructField, vf *reflect.Value) error {
	tagVals := strings.Split(tf.Tag.Get("servconf"), ",")
	tvc := len(tagVals)

	fName := strings.TrimSpace(tagVals[0])
	if fName == "" {
		return &EmptyNameTagError{tf.Name}
	}

	value, ok := sc[fName]

	required := false

	if tvc >= 2 {
		required = tagVals[1] == "required"
	}

	if !ok {
		if required {
			return &MissingRequiredError{fName}
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
