package main

import (
	"fmt"
	"reflect"
)

func main() {

	//t := reflect.TypeOf(3)
	//fmt.Println(t.String())
	//fmt.Println(t)
	//
	//var w io.Writer = os.Stdout
	//fmt.Println(reflect.TypeOf(w))
	//fmt.Printf("%T\n",w)
	//fmt.Printf("%T\n",3)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	vs := reflect.ValueOf("abc")
	fmt.Println(vs)
	fmt.Printf("%v\n", vs)
	fmt.Println(vs.String())

	fmt.Println("=======")
	t := v.Type()
	fmt.Println(t.String())
	ts := vs.Type()
	fmt.Println(ts.String())

}
