package main

import "fmt"
import "runtime"
import "reflect"

func main() {
  runtime.GOMAXPROCS(2)

  var intejer int = 13
  var setring string = "tigabelas"

  bacatipe1(reflect.ValueOf(intejer))
  go bacatipe2(reflect.ValueOf(setring))

  var input string
  fmt.Scanln(&input)
}

func bacatipe1(data reflect.Value) {
  fmt.Println("Tipe Variable tanpa go:", data.Kind())
}
func bacatipe2(data reflect.Value) {
  fmt.Println("Tipe Variable:", data.Kind())
}
