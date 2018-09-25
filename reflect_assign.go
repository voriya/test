package main

import (
	"reflect"
	"fmt"
)

type T struct {
	A *int64
	B string
}

func main() {
	var ti int64
	ti = 23
	t := T{&ti, "skidoo"}
	ChangeContent(&t)
	//s := reflect.ValueOf(&t).Elem()
	//typeOfT := s.Type()
	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i,
	//		typeOfT.Field(i).Name, f.Type(), f.Interface())
	//}
	//s.Field(0).SetInt(77)
	//s.Field(1).SetString("Sunset Strip")
	//fmt.Println("t is now", t)
}

func ChangeContent(t interface{}) {
	//reflect.ValueOf(t).Interface()

	tt := reflect.ValueOf(t)
	can := tt.CanSet()
	fmt.Printf("tt.canset?%v\n",can)

	can = tt.Elem().CanSet()
	fmt.Printf("tt.file0.canset?%v\n",can)

	var a int64
	a = 4
	tt.Elem().Field(0).Elem().SetInt(a)
	fmt.Println(tt.Interface())

	can = tt.Elem().Field(1).CanSet()
	fmt.Printf("tt.filed1.canset?%v\n",can)
}
