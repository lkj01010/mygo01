package main

//import(
//	"fmt"
//	"math"
//)
//
//func Squrt(i int) int{
//	v := math.Sqrt(float64(i))
//	return int(v)
//}
//
//func main() {
//	fmt.Println("result is ", Squrt(23))
//}


import (
    "fmt"
    "reflect"
)

func main() {
    type S struct {
         F string `species:"gopher" color:"blue"`
     }

     s := S{}
     st := reflect.TypeOf(s)
     field := st.Field(0)
     fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

 }