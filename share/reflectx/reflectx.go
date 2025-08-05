package reflectx

import "reflect"

func TypeOf(i any) reflect.Type {
	return Indirect(ValueOf(i)).Type()
}

func ValueOf(i any) reflect.Value {
	return reflect.ValueOf(i)
}

func Indirect(v reflect.Value) reflect.Value {
	return reflect.Indirect(v)
}

func KindOf(v any) reflect.Kind {
	return TypeOf(v).Kind()
}
