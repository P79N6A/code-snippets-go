package main

import (
	"fmt"
	"git.byted.org/ee/lark/facade/voip/utils/json"
	"github.com/binderclip/code-snippets-go/utils"
	"reflect"
)

type T3 struct {
	f1     string "f one"
	f2     string ``
	f3     string `f three`
	f4, f5 int64  `f four and five`
	f6     string
}

func tagSampleTest() {
	utils.PrintFuncName()
	t := reflect.TypeOf(T3{})
	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag) // f one
	f4, _ := t.FieldByName("f4")
	fmt.Println(f4.Tag)
	f5, _ := t.FieldByName("f5")
	fmt.Println(f5.Tag)

	f2, _ := t.FieldByName("f2")
	fmt.Printf("%q\n", f2.Tag)
	f6, _ := t.FieldByName("f6")
	fmt.Printf("%q\n", f6.Tag)

}

type T4 struct {
	f string `one:"1" two:"2" blank:""`
	f2 string "one:`10` two:\"20\"" // 必须是 `""` 的才能被正确 Lookup
}

type T5 struct {
	f string "one:\"1\" two:\"20\""
}

func tagSampleTest2() {
	utils.PrintFuncName()
	t := reflect.TypeOf(T4{})
	f, _ := t.FieldByName("f")
	v, ok := f.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
	v, ok = f.Tag.Lookup("blank")
	fmt.Printf("%s, %t\n", v, ok) // , true
	v, ok = f.Tag.Lookup("five")
	fmt.Printf("%s, %t\n", v, ok) // , false

	f2, _ := t.FieldByName("f2")
	v, ok = f2.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // , false
	v, ok = f2.Tag.Lookup("two")
	fmt.Printf("%s, %t\n", v, ok) // , false

	t = reflect.TypeOf(T5{})
	f, _ = t.FieldByName("f")
	fmt.Println(f.Tag) // one:"1"
	v, ok = f.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
	v, ok = f.Tag.Lookup("two")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
}

type T6 struct {
	F1 int `json:"f_1"`
	F2 int `json:"f_2,omitempty"`
	F3 int `json:"f_3,omitempty"`
	F4 int `json:"-"`
}

func tagForMarshaling() {
	utils.PrintFuncName()
	t := T6{1, 0, 2, 3}
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
	fmt.Println(b)
}

func main() {
	tagSampleTest()
	tagSampleTest2()
	tagForMarshaling()
}

// https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e
