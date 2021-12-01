package utils

import (
	"fmt"
	"reflect"
)

func CopyFields(dest interface{}, src interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(dest)
	av := reflect.ValueOf(dest)
	bt := reflect.TypeOf(src)
	bv := reflect.ValueOf(src)
	if at.Kind() != reflect.Ptr {
		err = fmt.Errorf("dest must be dest struct pointer")
		return
	}
	av = reflect.ValueOf(av.Interface())
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}
	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return
	}
	//tips: copy
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.FieldByName(name)
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			fmt.Printf("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}
