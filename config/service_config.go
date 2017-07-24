package config

import (
	"reflect"
	"strings"
)

type ServiceConfig map[string]interface{}

func (sc ServiceConfig) Assign(to interface{}) error {
	t := reflect.TypeOf(to)
	v := reflect.ValueOf(to)
	if v.Kind() != reflect.Struct {
		return ErrNotStruct
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
		if _, ok := tf.Tag.Lookup("servconf"); !ok {
			continue
		}
		vf := v.Field(i)
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
		return &InvalidTagError{tf.Name, ""}
	}

	value, ok := sc[fName]

	required := false

	if tvc >= 2 {
		required = tagVals[1] == "required"
	}

	if !ok {
		if required {
			return &InvalidTagError{tf.Name, fName}
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
